package com.cshep4.premierpredictor.notification.entity;

import com.cshep4.premierpredictor.notification.model.NotificationUser;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.bson.types.ObjectId;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class NotificationUserEntity {
    private ObjectId id;
    private String notificationToken;

    public NotificationUser toDto() {
        return NotificationUser.builder()
                .id(id.toHexString())
                .notificationToken(notificationToken)
                .build();
    }

    public static NotificationUserEntity fromDto(NotificationUser notification) throws IllegalArgumentException {
        if (!ObjectId.isValid(notification.getId())) {
            throw new IllegalArgumentException("Invalid id");
        }

        return NotificationUserEntity.builder()
                .id(new ObjectId(notification.getId()))
                .notificationToken(notification.getNotificationToken())
                .build();
    }
}
