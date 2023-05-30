// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: wso2/discovery/config/enforcer/security.proto

package enforcer

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Enforcer config model
type Security struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKey       *APIKeyEnforcer `protobuf:"bytes,1,opt,name=apiKey,proto3" json:"apiKey,omitempty"`
	RuntimeToken *APIKeyEnforcer `protobuf:"bytes,2,opt,name=runtimeToken,proto3" json:"runtimeToken,omitempty"`
	AuthHeader   *AuthHeader     `protobuf:"bytes,3,opt,name=authHeader,proto3" json:"authHeader,omitempty"`
	MutualSSL    *MutualSSL      `protobuf:"bytes,4,opt,name=mutualSSL,proto3" json:"mutualSSL,omitempty"`
}

func (x *Security) Reset() {
	*x = Security{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_config_enforcer_security_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Security) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Security) ProtoMessage() {}

func (x *Security) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_config_enforcer_security_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Security.ProtoReflect.Descriptor instead.
func (*Security) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_config_enforcer_security_proto_rawDescGZIP(), []int{0}
}

func (x *Security) GetApiKey() *APIKeyEnforcer {
	if x != nil {
		return x.ApiKey
	}
	return nil
}

func (x *Security) GetRuntimeToken() *APIKeyEnforcer {
	if x != nil {
		return x.RuntimeToken
	}
	return nil
}

func (x *Security) GetAuthHeader() *AuthHeader {
	if x != nil {
		return x.AuthHeader
	}
	return nil
}

func (x *Security) GetMutualSSL() *MutualSSL {
	if x != nil {
		return x.MutualSSL
	}
	return nil
}

var File_wso2_discovery_config_enforcer_security_proto protoreflect.FileDescriptor

var file_wso2_discovery_config_enforcer_security_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72,
	0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x1a,
	0x30, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2f, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65,
	0x72, 0x2f, 0x6d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x73, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x35, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x65, 0x6e, 0x66, 0x6f, 0x72,
	0x63, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x02, 0x0a, 0x08, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x46, 0x0a, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65,
	0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x45, 0x6e,
	0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x52,
	0x0a, 0x0c, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66,
	0x6f, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x45, 0x6e, 0x66, 0x6f,
	0x72, 0x63, 0x65, 0x72, 0x52, 0x0c, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x4a, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65,
	0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x47,
	0x0a, 0x09, 0x6d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x53, 0x53, 0x4c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63,
	0x65, 0x72, 0x2e, 0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x53, 0x53, 0x4c, 0x52, 0x09, 0x6d, 0x75,
	0x74, 0x75, 0x61, 0x6c, 0x53, 0x53, 0x4c, 0x42, 0x92, 0x01, 0x0a, 0x2f, 0x6f, 0x72, 0x67, 0x2e,
	0x77, 0x73, 0x6f, 0x32, 0x2e, 0x61, 0x70, 0x6b, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65,
	0x72, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x42, 0x0d, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x66, 0x6f, 0x72,
	0x63, 0x65, 0x72, 0x3b, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wso2_discovery_config_enforcer_security_proto_rawDescOnce sync.Once
	file_wso2_discovery_config_enforcer_security_proto_rawDescData = file_wso2_discovery_config_enforcer_security_proto_rawDesc
)

func file_wso2_discovery_config_enforcer_security_proto_rawDescGZIP() []byte {
	file_wso2_discovery_config_enforcer_security_proto_rawDescOnce.Do(func() {
		file_wso2_discovery_config_enforcer_security_proto_rawDescData = protoimpl.X.CompressGZIP(file_wso2_discovery_config_enforcer_security_proto_rawDescData)
	})
	return file_wso2_discovery_config_enforcer_security_proto_rawDescData
}

var file_wso2_discovery_config_enforcer_security_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wso2_discovery_config_enforcer_security_proto_goTypes = []interface{}{
	(*Security)(nil),       // 0: wso2.discovery.config.enforcer.Security
	(*APIKeyEnforcer)(nil), // 1: wso2.discovery.config.enforcer.APIKeyEnforcer
	(*AuthHeader)(nil),     // 2: wso2.discovery.config.enforcer.AuthHeader
	(*MutualSSL)(nil),      // 3: wso2.discovery.config.enforcer.MutualSSL
}
var file_wso2_discovery_config_enforcer_security_proto_depIdxs = []int32{
	1, // 0: wso2.discovery.config.enforcer.Security.apiKey:type_name -> wso2.discovery.config.enforcer.APIKeyEnforcer
	1, // 1: wso2.discovery.config.enforcer.Security.runtimeToken:type_name -> wso2.discovery.config.enforcer.APIKeyEnforcer
	2, // 2: wso2.discovery.config.enforcer.Security.authHeader:type_name -> wso2.discovery.config.enforcer.AuthHeader
	3, // 3: wso2.discovery.config.enforcer.Security.mutualSSL:type_name -> wso2.discovery.config.enforcer.MutualSSL
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_wso2_discovery_config_enforcer_security_proto_init() }
func file_wso2_discovery_config_enforcer_security_proto_init() {
	if File_wso2_discovery_config_enforcer_security_proto != nil {
		return
	}
	file_wso2_discovery_config_enforcer_auth_header_proto_init()
	file_wso2_discovery_config_enforcer_mutual_ssl_proto_init()
	file_wso2_discovery_config_enforcer_api_key_enforcer_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wso2_discovery_config_enforcer_security_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Security); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wso2_discovery_config_enforcer_security_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wso2_discovery_config_enforcer_security_proto_goTypes,
		DependencyIndexes: file_wso2_discovery_config_enforcer_security_proto_depIdxs,
		MessageInfos:      file_wso2_discovery_config_enforcer_security_proto_msgTypes,
	}.Build()
	File_wso2_discovery_config_enforcer_security_proto = out.File
	file_wso2_discovery_config_enforcer_security_proto_rawDesc = nil
	file_wso2_discovery_config_enforcer_security_proto_goTypes = nil
	file_wso2_discovery_config_enforcer_security_proto_depIdxs = nil
}
