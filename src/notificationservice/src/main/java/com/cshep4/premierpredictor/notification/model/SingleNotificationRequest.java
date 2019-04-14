package com.cshep4.premierpredictor.notification.model;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class SingleNotificationRequest {
    private String userId;
    private Notification notification;

    public static SingleNotificationRequest fromGrpc(com.cshep4.premierpredictor.notification.SingleRequest grpc) {
        return SingleNotificationRequest.builder()
                .userId(grpc.getUserId())
                .notification(Notification.fromGrpc(grpc.getNotification()))
                .build();
    }
}
