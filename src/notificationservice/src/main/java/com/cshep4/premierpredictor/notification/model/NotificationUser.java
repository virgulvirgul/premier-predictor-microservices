package com.cshep4.premierpredictor.notification.model;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class NotificationUser {
    private String id;
    private String notificationToken;

    public static NotificationUser fromGrpc(com.cshep4.premierpredictor.notification.SaveRequest grpc) {
        return NotificationUser.builder()
                .id(grpc.getUserId())
                .notificationToken(grpc.getNotificationToken())
                .build();
    }
}
