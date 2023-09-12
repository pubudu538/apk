/*
 *  Copyright (c) 2023, WSO2 LLC. (http://www.wso2.org).
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package xds

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	envoy_cachev3 "github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/wso2/apk/adapter/pkg/logging"

	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/wso2/apk/adapter/pkg/discovery/api/wso2/discovery/apkmgt"
	wso2_cache "github.com/wso2/apk/adapter/pkg/discovery/protocol/cache/v3"
	wso2_resource "github.com/wso2/apk/adapter/pkg/discovery/protocol/resource/v3"
	"github.com/wso2/apk/common-controller/internal/loggers"
	dpv1alpha1 "github.com/wso2/apk/common-controller/internal/operator/api/v1alpha1"
)

// EnforcerInternalAPI struct use to hold enforcer resources
type EnforcerInternalAPI struct {
	organizations []types.Resource
}

var (
	enforcerLabelMap          map[string]*EnforcerInternalAPI // Enforcer Label -> EnforcerInternalAPI struct map
	enforcerOrganizationCache wso2_cache.SnapshotCache
)

const (
	maxRandomInt             int    = 999999999
	grpcMaxConcurrentStreams        = 1000000
	apiKeyFieldSeparator     string = ":"
	commonEnforcerLabel      string = "commonEnforcerLabel"
)

func init() {
	enforcerOrganizationCache = wso2_cache.NewSnapshotCache(false, IDHash{}, nil)
	enforcerLabelMap = make(map[string]*EnforcerInternalAPI)
	enforcerLabelMap[commonEnforcerLabel] = &EnforcerInternalAPI{}
}

func maxRandomBigInt() *big.Int {
	return big.NewInt(int64(maxRandomInt))
}

// IDHash uses ID field as the node hash.
type IDHash struct{}

// ID uses the node ID field
func (IDHash) ID(node *corev3.Node) string {
	if node == nil {
		return "unknown"
	}
	return node.Id
}

var _ envoy_cachev3.NodeHash = IDHash{}

// GetRateLimiterCache returns xds server cache for rate limiter service.
func GetRateLimiterCache() envoy_cachev3.SnapshotCache {
	return rlsPolicyCache.xdsCache
}

// UpdateRateLimitXDSCache updates the xDS cache of the RateLimiter.
func UpdateRateLimitXDSCache(resolveRatelimit dpv1alpha1.ResolveRateLimitAPIPolicy) {
	// Add Rate Limit inline policies in API to the cache
	rlsPolicyCache.AddAPILevelRateLimitPolicies(resolveRatelimit)
}

// UpdateRateLimitXDSCacheForCustomPolicies updates the xDS cache of the RateLimiter for custom policies.
func UpdateRateLimitXDSCacheForCustomPolicies(customRateLimitPolicies dpv1alpha1.CustomRateLimitPolicyDef) {
	if customRateLimitPolicies.Key != "" {
		rlsPolicyCache.AddCustomRateLimitPolicies(customRateLimitPolicies)
	}
}

// DeleteAPILevelRateLimitPolicies delete the ratelimit xds cache
func DeleteAPILevelRateLimitPolicies(resolveRatelimit dpv1alpha1.ResolveRateLimitAPIPolicy) {
	var org = resolveRatelimit.Organization
	var environment = resolveRatelimit.Environment
	var context = resolveRatelimit.Context
	rlsPolicyCache.DeleteAPILevelRateLimitPolicies(org, environment, context)
}

// DeleteResourceLevelRateLimitPolicies delete the ratelimit xds cache
func DeleteResourceLevelRateLimitPolicies(resolveRatelimit dpv1alpha1.ResolveRateLimitAPIPolicy) {
	var org = resolveRatelimit.Organization
	var environment = resolveRatelimit.Environment
	var context = resolveRatelimit.Context
	var path = resolveRatelimit.Resources[0].Path
	var method = resolveRatelimit.Resources[0].Method
	rlsPolicyCache.DeleteResourceLevelRateLimitPolicies(org, environment, context, path, method)
}

// DeleteCustomRateLimitPolicies delete the ratelimit xds cache
func DeleteCustomRateLimitPolicies(customRateLimitPolicy dpv1alpha1.CustomRateLimitPolicyDef) {
	rlsPolicyCache.DeleteCustomRateLimitPolicies(customRateLimitPolicy)
}

// GenerateIdentifierForAPIWithUUID generates an identifier unique to the API
func GenerateIdentifierForAPIWithUUID(vhost, uuid string) string {
	return fmt.Sprint(vhost, apiKeyFieldSeparator, uuid)
}

// UpdateRateLimiterPolicies update the rate limiter xDS cache with latest rate limit policies
func UpdateRateLimiterPolicies(label string) {
	_ = rlsPolicyCache.updateXdsCache(label)
}

// SetEmptySnapshotupdate update empty snapshot
func SetEmptySnapshotupdate(lable string) bool {
	return rlsPolicyCache.SetEmptySnapshot(lable)
}

// GetEnforcerOrganizationCache returns xds server cache.
func GetEnforcerOrganizationCache() wso2_cache.SnapshotCache {
	return enforcerOrganizationCache
}

// UpdateEnforcerOrganizations updates xds server cache for organizations
func UpdateEnforcerOrganizations(organizations *apkmgt.OrganizationList) {

	loggers.LoggerAPKOperator.Infof("Updating Enforcer Organization Cache")
	label := commonEnforcerLabel
	organizationList := append(enforcerLabelMap[label].organizations, organizations)

	version := fmt.Sprint(rand.Int(rand.Reader, maxRandomBigInt()))
	snap, _ := wso2_cache.NewSnapshot(fmt.Sprint(version), map[wso2_resource.Type][]types.Resource{
		wso2_resource.OrganizationListType: organizationList,
	})
	snap.Consistent()

	errSetSnap := enforcerOrganizationCache.SetSnapshot(context.Background(), label, snap)
	if errSetSnap != nil {
		loggers.LoggerAPKOperator.ErrorC(logging.PrintError(logging.Error1716, logging.MAJOR,
			"Error while setting the snapshot : %v", errSetSnap.Error()))
	}

	enforcerLabelMap[label].organizations = organizationList
	loggers.LoggerAPKOperator.Infof("New Organization cache update for the label: " + label +
		" version: " + fmt.Sprint(version))
}
