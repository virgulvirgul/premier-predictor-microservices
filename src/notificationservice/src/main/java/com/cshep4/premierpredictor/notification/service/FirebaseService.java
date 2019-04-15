package com.cshep4.premierpredictor.notification.service;

import com.cshep4.premierpredictor.notification.model.Notification;
import com.google.auth.oauth2.GoogleCredentials;
import com.google.firebase.FirebaseApp;
import com.google.firebase.FirebaseOptions;
import com.google.firebase.messaging.FirebaseMessaging;
import com.google.firebase.messaging.FirebaseMessagingException;
import com.google.firebase.messaging.Message;
import com.google.firebase.messaging.MulticastMessage;
import lombok.SneakyThrows;
import lombok.val;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class FirebaseService {
    private static final String FCM_CONFIG_FILE = "firebase.json";

    @SneakyThrows
    public FirebaseService() {
        val classloader = Thread.currentThread().getContextClassLoader();

        try (val serviceAccountFile = classloader.getResourceAsStream(FCM_CONFIG_FILE)) {
            FirebaseOptions options = new FirebaseOptions.Builder()
                    .setCredentials(GoogleCredentials.fromStream(serviceAccountFile))
                    .build();

            FirebaseApp.initializeApp(options);
        }
    }

    public void sendNotification(Notification notification, String token) throws FirebaseMessagingException {
        val message = Message.builder()
                .setToken(token)
                .setNotification(new com.google.firebase.messaging.Notification(notification.getTitle(), notification.getMessage()))
                .build();

        FirebaseMessaging.getInstance().send(message);
    }

    public void sendNotification(Notification notification, List<String> tokens) throws FirebaseMessagingException {
        val message = MulticastMessage.builder()
                .addAllTokens(tokens)
                .setNotification(new com.google.firebase.messaging.Notification(notification.getTitle(), notification.getMessage()))
                .build();

        FirebaseMessaging.getInstance().sendMulticast(message);
    }
}
