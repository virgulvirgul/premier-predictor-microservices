package com.cshep4.premierpredictor.notification.model;

import lombok.Builder;
import lombok.Data;

import java.util.Queue;

import static com.cshep4.premierpredictor.notification.constant.Constants.NOTIFICATION_LIMIT;

@Data
@Builder
public class NotificationUser {
    private String id;
    private String notificationToken;
    private String lastReadNotification;
    private Queue<Notification> notifications;

    public static NotificationUser fromGrpc(com.cshep4.premierpredictor.notification.SaveRequest grpc) {
        return NotificationUser.builder()
                .id(grpc.getUserId())
                .notificationToken(grpc.getNotificationToken())
                .notifications(new CircularQueue<>(NOTIFICATION_LIMIT))
                .build();
    }
}
