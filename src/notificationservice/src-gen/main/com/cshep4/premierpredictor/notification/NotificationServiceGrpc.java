package com.cshep4.premierpredictor.notification;

import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.5.0)",
    comments = "Source: notification.proto")
public final class NotificationServiceGrpc {

  private NotificationServiceGrpc() {}

  public static final String SERVICE_NAME = "model.NotificationService";

  // Static method descriptors that strictly reflect the proto.
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<com.cshep4.premierpredictor.notification.SaveRequest,
      com.google.protobuf.Empty> METHOD_SAVE_USER =
      io.grpc.MethodDescriptor.<com.cshep4.premierpredictor.notification.SaveRequest, com.google.protobuf.Empty>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "model.NotificationService", "SaveUser"))
          .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.cshep4.premierpredictor.notification.SaveRequest.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.google.protobuf.Empty.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<com.cshep4.premierpredictor.notification.SingleRequest,
      com.google.protobuf.Empty> METHOD_SEND =
      io.grpc.MethodDescriptor.<com.cshep4.premierpredictor.notification.SingleRequest, com.google.protobuf.Empty>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "model.NotificationService", "Send"))
          .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.cshep4.premierpredictor.notification.SingleRequest.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.google.protobuf.Empty.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<com.cshep4.premierpredictor.notification.GroupRequest,
      com.google.protobuf.Empty> METHOD_SEND_TO_GROUP =
      io.grpc.MethodDescriptor.<com.cshep4.premierpredictor.notification.GroupRequest, com.google.protobuf.Empty>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "model.NotificationService", "SendToGroup"))
          .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.cshep4.premierpredictor.notification.GroupRequest.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.google.protobuf.Empty.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<com.cshep4.premierpredictor.notification.Notification,
      com.google.protobuf.Empty> METHOD_SEND_TO_ALL =
      io.grpc.MethodDescriptor.<com.cshep4.premierpredictor.notification.Notification, com.google.protobuf.Empty>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "model.NotificationService", "SendToAll"))
          .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.cshep4.premierpredictor.notification.Notification.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.google.protobuf.Empty.getDefaultInstance()))
          .build();

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static NotificationServiceStub newStub(io.grpc.Channel channel) {
    return new NotificationServiceStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static NotificationServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new NotificationServiceBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static NotificationServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new NotificationServiceFutureStub(channel);
  }

  /**
   */
  public static abstract class NotificationServiceImplBase implements io.grpc.BindableService {

    /**
     */
    public void saveUser(com.cshep4.premierpredictor.notification.SaveRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_SAVE_USER, responseObserver);
    }

    /**
     */
    public void send(com.cshep4.premierpredictor.notification.SingleRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_SEND, responseObserver);
    }

    /**
     */
    public void sendToGroup(com.cshep4.premierpredictor.notification.GroupRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_SEND_TO_GROUP, responseObserver);
    }

    /**
     */
    public void sendToAll(com.cshep4.premierpredictor.notification.Notification request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_SEND_TO_ALL, responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            METHOD_SAVE_USER,
            asyncUnaryCall(
              new MethodHandlers<
                com.cshep4.premierpredictor.notification.SaveRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_SAVE_USER)))
          .addMethod(
            METHOD_SEND,
            asyncUnaryCall(
              new MethodHandlers<
                com.cshep4.premierpredictor.notification.SingleRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_SEND)))
          .addMethod(
            METHOD_SEND_TO_GROUP,
            asyncUnaryCall(
              new MethodHandlers<
                com.cshep4.premierpredictor.notification.GroupRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_SEND_TO_GROUP)))
          .addMethod(
            METHOD_SEND_TO_ALL,
            asyncUnaryCall(
              new MethodHandlers<
                com.cshep4.premierpredictor.notification.Notification,
                com.google.protobuf.Empty>(
                  this, METHODID_SEND_TO_ALL)))
          .build();
    }
  }

  /**
   */
  public static final class NotificationServiceStub extends io.grpc.stub.AbstractStub<NotificationServiceStub> {
    private NotificationServiceStub(io.grpc.Channel channel) {
      super(channel);
    }

    private NotificationServiceStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected NotificationServiceStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new NotificationServiceStub(channel, callOptions);
    }

    /**
     */
    public void saveUser(com.cshep4.premierpredictor.notification.SaveRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_SAVE_USER, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void send(com.cshep4.premierpredictor.notification.SingleRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_SEND, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void sendToGroup(com.cshep4.premierpredictor.notification.GroupRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_SEND_TO_GROUP, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void sendToAll(com.cshep4.premierpredictor.notification.Notification request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_SEND_TO_ALL, getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class NotificationServiceBlockingStub extends io.grpc.stub.AbstractStub<NotificationServiceBlockingStub> {
    private NotificationServiceBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private NotificationServiceBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected NotificationServiceBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new NotificationServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.google.protobuf.Empty saveUser(com.cshep4.premierpredictor.notification.SaveRequest request) {
      return blockingUnaryCall(
          getChannel(), METHOD_SAVE_USER, getCallOptions(), request);
    }

    /**
     */
    public com.google.protobuf.Empty send(com.cshep4.premierpredictor.notification.SingleRequest request) {
      return blockingUnaryCall(
          getChannel(), METHOD_SEND, getCallOptions(), request);
    }

    /**
     */
    public com.google.protobuf.Empty sendToGroup(com.cshep4.premierpredictor.notification.GroupRequest request) {
      return blockingUnaryCall(
          getChannel(), METHOD_SEND_TO_GROUP, getCallOptions(), request);
    }

    /**
     */
    public com.google.protobuf.Empty sendToAll(com.cshep4.premierpredictor.notification.Notification request) {
      return blockingUnaryCall(
          getChannel(), METHOD_SEND_TO_ALL, getCallOptions(), request);
    }
  }

  /**
   */
  public static final class NotificationServiceFutureStub extends io.grpc.stub.AbstractStub<NotificationServiceFutureStub> {
    private NotificationServiceFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private NotificationServiceFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected NotificationServiceFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new NotificationServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> saveUser(
        com.cshep4.premierpredictor.notification.SaveRequest request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_SAVE_USER, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> send(
        com.cshep4.premierpredictor.notification.SingleRequest request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_SEND, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> sendToGroup(
        com.cshep4.premierpredictor.notification.GroupRequest request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_SEND_TO_GROUP, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> sendToAll(
        com.cshep4.premierpredictor.notification.Notification request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_SEND_TO_ALL, getCallOptions()), request);
    }
  }

  private static final int METHODID_SAVE_USER = 0;
  private static final int METHODID_SEND = 1;
  private static final int METHODID_SEND_TO_GROUP = 2;
  private static final int METHODID_SEND_TO_ALL = 3;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final NotificationServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(NotificationServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_SAVE_USER:
          serviceImpl.saveUser((com.cshep4.premierpredictor.notification.SaveRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_SEND:
          serviceImpl.send((com.cshep4.premierpredictor.notification.SingleRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_SEND_TO_GROUP:
          serviceImpl.sendToGroup((com.cshep4.premierpredictor.notification.GroupRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_SEND_TO_ALL:
          serviceImpl.sendToAll((com.cshep4.premierpredictor.notification.Notification) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static final class NotificationServiceDescriptorSupplier implements io.grpc.protobuf.ProtoFileDescriptorSupplier {
    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.cshep4.premierpredictor.notification.NotificationOuterClass.getDescriptor();
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (NotificationServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new NotificationServiceDescriptorSupplier())
              .addMethod(METHOD_SAVE_USER)
              .addMethod(METHOD_SEND)
              .addMethod(METHOD_SEND_TO_GROUP)
              .addMethod(METHOD_SEND_TO_ALL)
              .build();
        }
      }
    }
    return result;
  }
}
