/*
 *  Copyright (c) 2022, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
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

package controllers

import (
	"context"
	"fmt"

	xds "github.com/wso2/apk/common-controller/internal/xds"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/source"

	"github.com/wso2/apk/adapter/pkg/discovery/api/wso2/discovery/apkmgt"
	"github.com/wso2/apk/adapter/pkg/logging"
	"github.com/wso2/apk/common-controller/internal/config"
	loggers "github.com/wso2/apk/common-controller/internal/loggers"
	cpv1alpha1 "github.com/wso2/apk/common-controller/internal/operator/api/cp/v1alpha1"
	constants "github.com/wso2/apk/common-controller/internal/operator/constant"
)

// OrganizationReconciler reconciles a Organization object
type OrganizationReconciler struct {
	client client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=cp.wso2.com,resources=organizations,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cp.wso2.com,resources=organizations/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=cp.wso2.com,resources=organizations/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Organization object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.15.0/pkg/reconcile
func (r *OrganizationReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	loggers.LoggerAPKOperator.Infof("Reconciling organizations...")
	ratelimitKey := req.NamespacedName

	loggers.LoggerAPKOperator.Infof("Reconciling organization: %s", ratelimitKey.Name)

	orgKey := req.NamespacedName
	var orgList = new(cpv1alpha1.OrganizationList)
	if err := r.client.List(ctx, orgList); err != nil {
		return ctrl.Result{}, fmt.Errorf("failed to get organizations %s/%s",
			orgKey.Namespace, orgKey.Name)
	}

	marshalledOrgList := marshalOrganizationList(orgList.Items)
	xds.UpdateEnforcerOrganizations(marshalledOrgList)
	return ctrl.Result{}, nil
}

func marshalOrganizationList(organizationList []cpv1alpha1.Organization) *apkmgt.OrganizationList {
	organizations := []*apkmgt.Organization{}
	for _, orgInternal := range organizationList {
		org := &apkmgt.Organization{
			Name: orgInternal.Spec.Name,
			Id:   orgInternal.Spec.ID,
		}
		organizations = append(organizations, org)
	}
	return &apkmgt.OrganizationList{
		List: organizations,
	}
}

// NewOrganizationController creates a new organizationController instance.
func NewOrganizationController(mgr manager.Manager) error {

	organizationReconciler := &OrganizationReconciler{
		client: mgr.GetClient(),
	}

	c, err := controller.New(constants.OrganizationController, mgr, controller.Options{Reconciler: organizationReconciler})
	if err != nil {
		loggers.LoggerAPKOperator.ErrorC(logging.PrintError(logging.Error2663, logging.BLOCKER,
			"Error creating Organization controller: %v", err.Error()))
		return err
	}

	conf := config.ReadConfigs()
	predicates := []predicate.Predicate{predicate.NewPredicateFuncs(FilterByNamespaces(conf.CommonController.Operator.Namespaces))}

	if err := c.Watch(source.Kind(mgr.GetCache(), &cpv1alpha1.Organization{}), &handler.EnqueueRequestForObject{}, predicates...); err != nil {
		loggers.LoggerAPKOperator.ErrorC(logging.PrintError(logging.Error2639, logging.BLOCKER,
			"Error watching Organization resources: %v", err.Error()))
		return err
	}

	loggers.LoggerAPKOperator.Debug("Organization Controller successfully started. Watching Organization Objects...")
	return nil
}
