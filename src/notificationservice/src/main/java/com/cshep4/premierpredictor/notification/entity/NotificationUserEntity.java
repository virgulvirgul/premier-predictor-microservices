package com.cshep4.premierpredictor.notification.entity;

import com.cshep4.premierpredictor.notification.model.CircularQueue;
import com.cshep4.premierpredictor.notification.model.Notification;
import com.cshep4.premierpredictor.notification.model.NotificationUser;
import lombok.*;
import org.bson.types.ObjectId;

import java.util.ArrayList;
import java.util.List;
import java.util.Queue;

import static com.cshep4.premierpredictor.notification.constant.Constants.NOTIFICATION_LIMIT;
import static java.util.stream.Collectors.toCollection;
import static org.bson.types.ObjectId.isValid;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class NotificationUserEntity {
    private ObjectId id;
    private String notificationToken;
    private ObjectId lastReadNotification;
    private List<NotificationEntity> notifications;

    public NotificationUser toDto() {
        Queue<Notification> n = new CircularQueue<>(NOTIFICATION_LIMIT);

        if (notifications != null) {
            n = notifications.stream()
                    .map(NotificationEntity::toDto)
                    .collect(toCollection(() -> new CircularQueue<>(NOTIFICATION_LIMIT)));
        }

        val notificationUser = NotificationUser.builder()
                .id(id.toHexString())
                .notificationToken(notificationToken)
                .notifications(n)
                .build();

        if (lastReadNotification != null) {
            notificationUser.setLastReadNotification(lastReadNotification.toHexString());
        }

        return notificationUser;
    }

    public static NotificationUserEntity fromDto(NotificationUser notification) throws IllegalArgumentException {
        if (!isValid(notification.getId())) {
            throw new IllegalArgumentException("Invalid id");
        }

        List<NotificationEntity> notifications = new CircularQueue<>(NOTIFICATION_LIMIT);

        if (notification.getNotifications() != null) {
            notifications = notification.getNotifications()
                    .stream()
                    .map(NotificationEntity::of)
                    .collect(toCollection(() -> new CircularQueue<>(NOTIFICATION_LIMIT)));
        }

        val notificationUserEntity = NotificationUserEntity.builder()
                .id(new ObjectId(notification.getId()))
                .notificationToken(notification.getNotificationToken())
                .notifications(notifications)
                .build();

        if (notification.getLastReadNotification() != null && isValid(notification.getLastReadNotification())) {
            notificationUserEntity.setLastReadNotification(new ObjectId(notification.getLastReadNotification()));
        }

        return notificationUserEntity;
    }
}
