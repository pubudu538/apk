// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: wso2/discovery/config/enforcer/api_key_enforcer.proto

package org.wso2.apk.enforcer.discovery.config.enforcer;

public final class APIKeyEnforcerProto {
  private APIKeyEnforcerProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_wso2_discovery_config_enforcer_APIKeyEnforcer_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_wso2_discovery_config_enforcer_APIKeyEnforcer_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n5wso2/discovery/config/enforcer/api_key" +
      "_enforcer.proto\022\036wso2.discovery.config.e" +
      "nforcer\"N\n\016APIKeyEnforcer\022\017\n\007enabled\030\001 \001" +
      "(\010\022\016\n\006issuer\030\002 \001(\t\022\033\n\023certificateFilePat" +
      "h\030\003 \001(\tB\230\001\n/org.wso2.apk.enforcer.discov" +
      "ery.config.enforcerB\023APIKeyEnforcerProto" +
      "P\001ZNgithub.com/envoyproxy/go-control-pla" +
      "ne/wso2/discovery/config/enforcer;enforc" +
      "erb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        });
    internal_static_wso2_discovery_config_enforcer_APIKeyEnforcer_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_wso2_discovery_config_enforcer_APIKeyEnforcer_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_wso2_discovery_config_enforcer_APIKeyEnforcer_descriptor,
        new java.lang.String[] { "Enabled", "Issuer", "CertificateFilePath", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}
