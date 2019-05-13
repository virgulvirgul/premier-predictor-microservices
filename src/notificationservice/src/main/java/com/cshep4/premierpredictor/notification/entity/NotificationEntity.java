package com.cshep4.premierpredictor.notification.entity;

import com.cshep4.premierpredictor.notification.model.Notification;
import lombok.*;
import org.bson.types.ObjectId;

import static org.bson.types.ObjectId.isValid;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class NotificationEntity {
    private ObjectId id;
    private String title;
    private String message;

    public Notification toDto() {
        return Notification.builder()
                .id(id.toHexString())
                .title(title)
                .message(message)
                .build();
    }

    public static NotificationEntity of(Notification notification) {
        return NotificationEntity.builder()
                .id(new ObjectId())
                .title(notification.getTitle())
                .message(notification.getMessage())
                .build();
    }

    public static NotificationEntity fromDto(Notification notification) {
        if (isValid(notification.getId())) {
            throw new IllegalArgumentException("Invalid id");
        }

        val entity = NotificationEntity.builder()
                .id(new ObjectId(notification.getId()))
                .title(notification.getTitle())
                .message(notification.getMessage())
                .build();

        return entity;
    }
}
