package com.cshep4.premierpredictor.notification.grpc;

import com.cshep4.premierpredictor.notification.*;
import com.cshep4.premierpredictor.notification.NotificationServiceGrpc.NotificationServiceImplBase;
import com.cshep4.premierpredictor.notification.interceptor.AuthInterceptor;
import com.cshep4.premierpredictor.notification.model.GroupNotificationRequest;
import com.cshep4.premierpredictor.notification.model.NotificationUser;
import com.cshep4.premierpredictor.notification.model.SingleNotificationRequest;
import com.cshep4.premierpredictor.notification.service.NotificationService;
import com.cshep4.premierpredictor.request.IdRequest;
import com.google.firebase.messaging.FirebaseMessagingException;
import com.google.protobuf.Empty;
import io.grpc.stub.StreamObserver;
import io.reactivex.observers.DisposableObserver;
import lombok.val;
import org.lognet.springboot.grpc.GRpcService;
import org.springframework.beans.factory.annotation.Autowired;

import static com.cshep4.premierpredictor.notification.model.Notification.fromGrpc;

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

    @Override
    public void getNotifications(IdRequest request, StreamObserver<NotificationResponse> responseObserver) {
        notificationService.getNotifications(request.getId())
                .stream()
                .map(com.cshep4.premierpredictor.notification.model.Notification::toGrpc)
                .forEach(responseObserver::onNext);

        val notificationObserver = new DisposableObserver<com.cshep4.premierpredictor.notification.model.Notification>() {
            @Override
            public void onNext(com.cshep4.premierpredictor.notification.model.Notification notification) {
                responseObserver.onNext(notification.toGrpc());
            }

            @Override
            public void onError(Throwable e) {
                responseObserver.onError(e);
            }

            @Override
            public void onComplete() {
                responseObserver.onCompleted();
            }
        };

        notificationService.subscribeToUpdates(request.getId(), notificationObserver);
    }

    @Override
    public void updateReadNotification(UpdateReadRequest request, StreamObserver<Empty> responseObserver) {
        notificationService.updateReadNotification(request.getUserId(), request.getNotificationId());

        successResponse(responseObserver);
    }

    private void successResponse(StreamObserver<Empty> responseObserver) {
        responseObserver.onNext(EMPTY);
        responseObserver.onCompleted();
    }
}
