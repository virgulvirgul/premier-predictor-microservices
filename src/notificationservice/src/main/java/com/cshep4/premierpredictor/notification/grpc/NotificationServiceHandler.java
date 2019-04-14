package com.cshep4.premierpredictor.notification.grpc;

import com.cshep4.premierpredictor.notification.GroupRequest;
import com.cshep4.premierpredictor.notification.Notification;
import com.cshep4.premierpredictor.notification.NotificationServiceGrpc.NotificationServiceImplBase;
import com.cshep4.premierpredictor.notification.SaveRequest;
import com.cshep4.premierpredictor.notification.SingleRequest;
import com.cshep4.premierpredictor.notification.interceptor.AuthInterceptor;
import com.cshep4.premierpredictor.notification.model.GroupNotificationRequest;
import com.cshep4.premierpredictor.notification.model.NotificationUser;
import com.cshep4.premierpredictor.notification.model.SingleNotificationRequest;
import com.cshep4.premierpredictor.notification.service.NotificationService;
import com.google.firebase.messaging.FirebaseMessagingException;
import com.google.protobuf.Empty;
import io.grpc.stub.StreamObserver;
import lombok.val;
import org.lognet.springboot.grpc.GRpcService;
import org.springframework.beans.factory.annotation.Autowired;

import static com.cshep4.premierpredictor.notification.model.Notification.*;

@GRpcService(interceptors = AuthInterceptor.class)
public class NotificationServiceHandler extends NotificationServiceImplBase {
    private static final Empty EMPTY = Empty.newBuilder().build();

    @Autowired
    private NotificationService notificationService;

    @Override
    public void saveUser(SaveRequest request, StreamObserver<Empty> responseObserver) {
        val notificationUser = NotificationUser.fromGrpc(request);

        notificationService.saveUser(notificationUser);

        successResponse(responseObserver);
    }

    @Override
    public void send(SingleRequest request, StreamObserver<Empty> responseObserver) {
        val notificationRequest = SingleNotificationRequest.fromGrpc(request);

        try {
            notificationService.send(notificationRequest);

            successResponse(responseObserver);
        } catch (FirebaseMessagingException e) {
            responseObserver.onError(e);
        }
    }

    @Override
    public void sendToGroup(GroupRequest request, StreamObserver<Empty> responseObserver) {
        val notificationRequest = GroupNotificationRequest.fromGrpc(request);

        try {
            notificationService.send(notificationRequest);

            successResponse(responseObserver);
        } catch (FirebaseMessagingException e) {
            responseObserver.onError(e);
        }
    }

    @Override
    public void sendToAll(Notification request, StreamObserver<Empty> responseObserver) {
        val notificationRequest = fromGrpc(request);

        try {
            notificationService.send(notificationRequest);

            successResponse(responseObserver);
        } catch (FirebaseMessagingException e) {
            responseObserver.onError(e);
        }
    }

    private void successResponse(StreamObserver<Empty> responseObserver) {
        responseObserver.onNext(EMPTY);
        responseObserver.onCompleted();
    }
}
