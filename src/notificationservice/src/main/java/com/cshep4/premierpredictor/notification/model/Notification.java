package com.cshep4.premierpredictor.notification.model;

import com.cshep4.premierpredictor.notification.NotificationResponse;
import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class Notification {
    private String id;
    private String title;
    private String message;

    public static Notification fromGrpc(com.cshep4.premierpredictor.notification.Notification notification) {
        return Notification.builder()
                .title(notification.getTitle())
                .message(notification.getMessage())
                .build();
    }

    public NotificationResponse toGrpc() {
        return NotificationResponse.newBuilder()
                .setId(id)
                .setTitle(title)
                .setMessage(message)
                .build();
    }
}
