package org.wso2.apk.enforcer.discovery.service.apkmgt;

import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;

/**
 * <pre>
 * [#protodoc-title: OrganizationDS]
 * </pre>
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler",
    comments = "Source: wso2/discovery/service/apkmgt/organizationds.proto")
public final class APKMgtOrgDiscoveryServiceGrpc {

  private APKMgtOrgDiscoveryServiceGrpc() {}

  public static final String SERVICE_NAME = "discovery.service.apkmgt.APKMgtOrgDiscoveryService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<io.envoyproxy.envoy.service.discovery.v3.DiscoveryRequest,
      io.envoyproxy.envoy.service.discovery.v3.DiscoveryResponse> getStreamAPKMgtOrganizationsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "StreamAPKMgtOrganizations",
      requestType = io.envoyproxy.envoy.service.discovery.v3.DiscoveryRequest.class,
      responseType = io.envoyproxy.envoy.service.discovery.v3.DiscoveryResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.BIDI_STREAMING)
  public static io.grpc.MethodDescriptor<io.envoyproxy.envoy.service.discovery.v3.DiscoveryRequest,
      io.envoyproxy.envoy.service.discovery.v3.DiscoveryResponse> getStreamAPKMgtOrganizationsMethod() {
    io.grpc.MethodDescriptor<io.envoyproxy.envoy.service.discovery.v3.DiscoveryRequest, io.envoyproxy.envoy.service.discovery.v3.DiscoveryResponse> getStreamAPKMgtOrganizationsMethod;
    if ((getStreamAPKMgtOrganizationsMethod = APKMgtOrgDiscoveryServiceGrpc.getStreamAPKMgtOrganizationsMethod) == null) {
      synchronized (APKMgtOrgDiscoveryServiceGrpc.class) {
        if ((getStreamAPKMgtOrganizationsMethod = APKMgtOrgDiscoveryServiceGrpc.getStreamAPKMgtOrganizationsMethod) == null) {
          APKMgtOrgDiscoveryServiceGrpc.getStreamAPKMgtOrganizationsMethod = getStreamAPKMgtOrganizationsMethod =
              io.grpc.MethodDescriptor.<io.envoyproxy.envoy.service.discovery.v3.DiscoveryRequest, io.envoyproxy.envoy.service.discovery.v3.DiscoveryResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.BIDI_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "StreamAPKMgtOrganizations"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.envoyproxy.envoy.service.discovery.v3.DiscoveryRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.envoyproxy.envoy.service.discovery.v3.DiscoveryResponse.getDefaultInstance()))
              .setSchemaDescriptor(new APKMgtOrgDiscoveryServiceMethodDescriptorSupplier("StreamAPKMgtOrganizations"))
              .build();
        }
      }
    }
    return getStreamAPKMgtOrganizationsMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static APKMgtOrgDiscoveryServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<APKMgtOrgDiscoveryServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<APKMgtOrgDiscoveryServiceStub>() {
        @java.lang.Override
        public APKMgtOrgDiscoveryServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new APKMgtOrgDiscoveryServiceStub(channel, callOptions);
        }
      };
    return APKMgtOrgDiscoveryServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static APKMgtOrgDiscoveryServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<APKMgtOrgDiscoveryServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<APKMgtOrgDiscoveryServiceBlockingStub>() {
        @java.lang.Override
        public APKMgtOrgDiscoveryServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new APKMgtOrgDiscoveryServiceBlockingStub(channel, callOptions);
        }
      };
    return APKMgtOrgDiscoveryServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static APKMgtOrgDiscoveryServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<APKMgtOrgDiscoveryServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<APKMgtOrgDiscoveryServiceFutureStub>() {
        @java.lang.Override
        public APKMgtOrgDiscoveryServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new APKMgtOrgDiscoveryServiceFutureStub(channel, callOptions);
        }
      };
    return APKMgtOrgDiscoveryServiceFutureStub.newStub(factory, channel);
  }

  /**
   * <pre>
   * [#protodoc-title: OrganizationDS]
   * </pre>
   */
  public static abstract class APKMgtOrgDiscoveryServiceImplBase implements io.grpc.BindableService {

    /**
     */
    public io.grpc.stub.StreamObserver<io.envoyproxy.envoy.service.discovery.v3.DiscoveryRequest> streamAPKMgtOrganizations(
        io.grpc.stub.StreamObserver<io.envoyproxy.envoy.service.discovery.v3.DiscoveryResponse> responseObserver) {
      return asyncUnimplementedStreamingCall(getStreamAPKMgtOrganizationsMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getStreamAPKMgtOrganizationsMethod(),
            asyncBidiStreamingCall(
              new MethodHandlers<
                io.envoyproxy.envoy.service.discovery.v3.DiscoveryRequest,
                io.envoyproxy.envoy.service.discovery.v3.DiscoveryResponse>(
                  this, METHODID_STREAM_APKMGT_ORGANIZATIONS)))
          .build();
    }
  }

  /**
   * <pre>
   * [#protodoc-title: OrganizationDS]
   * </pre>
   */
  public static final class APKMgtOrgDiscoveryServiceStub extends io.grpc.stub.AbstractAsyncStub<APKMgtOrgDiscoveryServiceStub> {
    private APKMgtOrgDiscoveryServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected APKMgtOrgDiscoveryServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new APKMgtOrgDiscoveryServiceStub(channel, callOptions);
    }

    /**
     */
    public io.grpc.stub.StreamObserver<io.envoyproxy.envoy.service.discovery.v3.DiscoveryRequest> streamAPKMgtOrganizations(
        io.grpc.stub.StreamObserver<io.envoyproxy.envoy.service.discovery.v3.DiscoveryResponse> responseObserver) {
      return asyncBidiStreamingCall(
          getChannel().newCall(getStreamAPKMgtOrganizationsMethod(), getCallOptions()), responseObserver);
    }
  }

  /**
   * <pre>
   * [#protodoc-title: OrganizationDS]
   * </pre>
   */
  public static final class APKMgtOrgDiscoveryServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<APKMgtOrgDiscoveryServiceBlockingStub> {
    private APKMgtOrgDiscoveryServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected APKMgtOrgDiscoveryServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new APKMgtOrgDiscoveryServiceBlockingStub(channel, callOptions);
    }
  }

  /**
   * <pre>
   * [#protodoc-title: OrganizationDS]
   * </pre>
   */
  public static final class APKMgtOrgDiscoveryServiceFutureStub extends io.grpc.stub.AbstractFutureStub<APKMgtOrgDiscoveryServiceFutureStub> {
    private APKMgtOrgDiscoveryServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected APKMgtOrgDiscoveryServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new APKMgtOrgDiscoveryServiceFutureStub(channel, callOptions);
    }
  }

  private static final int METHODID_STREAM_APKMGT_ORGANIZATIONS = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final APKMgtOrgDiscoveryServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(APKMgtOrgDiscoveryServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_STREAM_APKMGT_ORGANIZATIONS:
          return (io.grpc.stub.StreamObserver<Req>) serviceImpl.streamAPKMgtOrganizations(
              (io.grpc.stub.StreamObserver<io.envoyproxy.envoy.service.discovery.v3.DiscoveryResponse>) responseObserver);
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class APKMgtOrgDiscoveryServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    APKMgtOrgDiscoveryServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return org.wso2.apk.enforcer.discovery.service.apkmgt.OrganizationDsProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("APKMgtOrgDiscoveryService");
    }
  }

  private static final class APKMgtOrgDiscoveryServiceFileDescriptorSupplier
      extends APKMgtOrgDiscoveryServiceBaseDescriptorSupplier {
    APKMgtOrgDiscoveryServiceFileDescriptorSupplier() {}
  }

  private static final class APKMgtOrgDiscoveryServiceMethodDescriptorSupplier
      extends APKMgtOrgDiscoveryServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    APKMgtOrgDiscoveryServiceMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (APKMgtOrgDiscoveryServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new APKMgtOrgDiscoveryServiceFileDescriptorSupplier())
              .addMethod(getStreamAPKMgtOrganizationsMethod())
              .build();
        }
      }
    }
    return result;
  }
}
