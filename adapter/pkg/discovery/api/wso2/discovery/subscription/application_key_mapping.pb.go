// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: wso2/discovery/subscription/application_key_mapping.proto

package subscription

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

// ApplicationKeyMapping data model
type ApplicationKeyMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerKey     string `protobuf:"bytes,1,opt,name=consumerKey,proto3" json:"consumerKey,omitempty"`
	KeyType         string `protobuf:"bytes,2,opt,name=keyType,proto3" json:"keyType,omitempty"`
	KeyManager      string `protobuf:"bytes,3,opt,name=keyManager,proto3" json:"keyManager,omitempty"`
	ApplicationId   int32  `protobuf:"varint,4,opt,name=applicationId,proto3" json:"applicationId,omitempty"`
	TenantId        int32  `protobuf:"varint,5,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	TenantDomain    string `protobuf:"bytes,6,opt,name=tenantDomain,proto3" json:"tenantDomain,omitempty"`
	Timestamp       int64  `protobuf:"varint,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ApplicationUUID string `protobuf:"bytes,8,opt,name=applicationUUID,proto3" json:"applicationUUID,omitempty"`
}

func (x *ApplicationKeyMapping) Reset() {
	*x = ApplicationKeyMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_subscription_application_key_mapping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationKeyMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationKeyMapping) ProtoMessage() {}

func (x *ApplicationKeyMapping) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_subscription_application_key_mapping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationKeyMapping.ProtoReflect.Descriptor instead.
func (*ApplicationKeyMapping) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_subscription_application_key_mapping_proto_rawDescGZIP(), []int{0}
}

func (x *ApplicationKeyMapping) GetConsumerKey() string {
	if x != nil {
		return x.ConsumerKey
	}
	return ""
}

func (x *ApplicationKeyMapping) GetKeyType() string {
	if x != nil {
		return x.KeyType
	}
	return ""
}

func (x *ApplicationKeyMapping) GetKeyManager() string {
	if x != nil {
		return x.KeyManager
	}
	return ""
}

func (x *ApplicationKeyMapping) GetApplicationId() int32 {
	if x != nil {
		return x.ApplicationId
	}
	return 0
}

func (x *ApplicationKeyMapping) GetTenantId() int32 {
	if x != nil {
		return x.TenantId
	}
	return 0
}

func (x *ApplicationKeyMapping) GetTenantDomain() string {
	if x != nil {
		return x.TenantDomain
	}
	return ""
}

func (x *ApplicationKeyMapping) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ApplicationKeyMapping) GetApplicationUUID() string {
	if x != nil {
		return x.ApplicationUUID
	}
	return ""
}

var File_wso2_discovery_subscription_application_key_mapping_proto protoreflect.FileDescriptor

var file_wso2_discovery_subscription_application_key_mapping_proto_rawDesc = []byte{
	0x0a, 0x39, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x6d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x77, 0x73, 0x6f,
	0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa1, 0x02, 0x0a, 0x15, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x4d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x6b, 0x65, 0x79, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6b, 0x65, 0x79, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x24,
	0x0a, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x22, 0x0a, 0x0c, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x55, 0x55, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x55, 0x49, 0x44, 0x42, 0x9f, 0x01, 0x0a,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x65, 0x6f,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x1a, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x4d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x3b, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wso2_discovery_subscription_application_key_mapping_proto_rawDescOnce sync.Once
	file_wso2_discovery_subscription_application_key_mapping_proto_rawDescData = file_wso2_discovery_subscription_application_key_mapping_proto_rawDesc
)

func file_wso2_discovery_subscription_application_key_mapping_proto_rawDescGZIP() []byte {
	file_wso2_discovery_subscription_application_key_mapping_proto_rawDescOnce.Do(func() {
		file_wso2_discovery_subscription_application_key_mapping_proto_rawDescData = protoimpl.X.CompressGZIP(file_wso2_discovery_subscription_application_key_mapping_proto_rawDescData)
	})
	return file_wso2_discovery_subscription_application_key_mapping_proto_rawDescData
}

var file_wso2_discovery_subscription_application_key_mapping_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wso2_discovery_subscription_application_key_mapping_proto_goTypes = []interface{}{
	(*ApplicationKeyMapping)(nil), // 0: wso2.discovery.subscription.ApplicationKeyMapping
}
var file_wso2_discovery_subscription_application_key_mapping_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wso2_discovery_subscription_application_key_mapping_proto_init() }
func file_wso2_discovery_subscription_application_key_mapping_proto_init() {
	if File_wso2_discovery_subscription_application_key_mapping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wso2_discovery_subscription_application_key_mapping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationKeyMapping); i {
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
			RawDescriptor: file_wso2_discovery_subscription_application_key_mapping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wso2_discovery_subscription_application_key_mapping_proto_goTypes,
		DependencyIndexes: file_wso2_discovery_subscription_application_key_mapping_proto_depIdxs,
		MessageInfos:      file_wso2_discovery_subscription_application_key_mapping_proto_msgTypes,
	}.Build()
	File_wso2_discovery_subscription_application_key_mapping_proto = out.File
	file_wso2_discovery_subscription_application_key_mapping_proto_rawDesc = nil
	file_wso2_discovery_subscription_application_key_mapping_proto_goTypes = nil
	file_wso2_discovery_subscription_application_key_mapping_proto_depIdxs = nil
}
