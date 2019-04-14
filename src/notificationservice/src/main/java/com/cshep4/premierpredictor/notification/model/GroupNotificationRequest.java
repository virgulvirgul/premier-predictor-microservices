package com.cshep4.premierpredictor.notification.model;

import lombok.Builder;
import lombok.Data;

import java.util.List;

@Data
@Builder
public class GroupNotificationRequest {
    private List<String> userIds;
    private Notification notification;

    public static GroupNotificationRequest fromGrpc(com.cshep4.premierpredictor.notification.GroupRequest grpc) {
        return GroupNotificationRequest.builder()
                .userIds(grpc.getUserIdsList())
                .notification(Notification.fromGrpc(grpc.getNotification()))
                .build();
    }
}