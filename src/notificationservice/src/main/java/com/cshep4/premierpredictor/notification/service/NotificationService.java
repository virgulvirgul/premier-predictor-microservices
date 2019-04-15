package com.cshep4.premierpredictor.notification.service;

import com.cshep4.premierpredictor.notification.model.GroupNotificationRequest;
import com.cshep4.premierpredictor.notification.model.Notification;
import com.cshep4.premierpredictor.notification.model.NotificationUser;
import com.cshep4.premierpredictor.notification.model.SingleNotificationRequest;
import com.cshep4.premierpredictor.notification.repository.NotificationRepository;
import com.google.firebase.messaging.FirebaseMessagingException;
import lombok.val;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import static java.util.stream.Collectors.toList;

@Service
public class NotificationService {
    @Autowired
    private NotificationRepository notificationRepository;

    @Autowired
    private FirebaseService firebaseService;

    public void saveUser(NotificationUser notificationUser) {
        notificationRepository.save(notificationUser);
    }

    public void send(SingleNotificationRequest req) throws FirebaseMessagingException {
        val user = notificationRepository.getById(req.getUserId());

        if (user == null) {
            return;
        }

        firebaseService.sendNotification(req.getNotification(), user.getNotificationToken());
    }

    public void send(GroupNotificationRequest req) throws FirebaseMessagingException {
        val tokens = notificationRepository.getAllByIds(req.getUserIds())
                .stream()
                .map(NotificationUser::getNotificationToken)
                .collect(toList());

        if (tokens.size() == 0) {
            return;
        }

        firebaseService.sendNotification(req.getNotification(), tokens);
    }

    public void send(Notification notification) throws FirebaseMessagingException {
        val tokens = notificationRepository.getAll()
                .stream()
                .map(NotificationUser::getNotificationToken)
                .collect(toList());

        if (tokens.size() == 0) {
            return;
        }

        firebaseService.sendNotification(notification, tokens);
    }
}
