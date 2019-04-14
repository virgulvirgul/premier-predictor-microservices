package com.cshep4.premierpredictor.notification.service;

import com.cshep4.premierpredictor.notification.model.GroupNotificationRequest;
import com.cshep4.premierpredictor.notification.model.Notification;
import com.cshep4.premierpredictor.notification.model.NotificationUser;
import com.cshep4.premierpredictor.notification.model.SingleNotificationRequest;
import com.cshep4.premierpredictor.notification.repository.NotificationRepository;
import lombok.SneakyThrows;
import lombok.val;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.MockitoJUnitRunner;

import static java.util.Collections.singletonList;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

@RunWith(MockitoJUnitRunner.class)
public class NotificationServiceTest {
    private static final String ID = "id";
    private static final String NOTIFICATION_TOKEN = "token";
    private static final String TITLE = "title";
    private static final String MESSAGE = "message";

    @Mock
    private NotificationRepository notificationRepository;

    @Mock
    private FirebaseService firebaseService;

    @InjectMocks
    private NotificationService notificationService;

    @Test
    public void saveUserToDatabase() {
        val notificationUser = NotificationUser.builder().build();

        notificationService.saveUser(notificationUser);

        verify(notificationRepository).save(notificationUser);
    }

    @Test
    @SneakyThrows
    public void sendSingleNotification() {
        val notificationUser = NotificationUser.builder()
                .id(ID)
                .notificationToken(NOTIFICATION_TOKEN)
                .build();

        when(notificationRepository.getById(ID)).thenReturn(notificationUser);

        val notification = Notification.builder()
                .title(TITLE)
                .message(MESSAGE)
                .build();

        val request = SingleNotificationRequest.builder()
                .userId(ID)
                .notification(notification)
                .build();

        notificationService.send(request);

        verify(notificationRepository).getById(ID);
        verify(firebaseService).sendNotification(notification, NOTIFICATION_TOKEN);
    }

    @Test
    @SneakyThrows
    public void sendMultipleNotifications() {
        val ids = singletonList(ID);

        val notificationUser = NotificationUser.builder()
                .id(ID)
                .notificationToken(NOTIFICATION_TOKEN)
                .build();

        val notificationUsers = singletonList(notificationUser);

        val tokens = singletonList(NOTIFICATION_TOKEN);

        when(notificationRepository.getAllByIds(ids)).thenReturn(notificationUsers);

        val notification = Notification.builder()
                .title(TITLE)
                .message(MESSAGE)
                .build();

        val request = GroupNotificationRequest.builder()
                .userIds(ids)
                .notification(notification)
                .build();

        notificationService.send(request);

        verify(notificationRepository).getAllByIds(ids);
        verify(firebaseService).sendNotification(notification, tokens);
    }

    @Test
    @SneakyThrows
    public void sendNotificationsToAll() {
        val notificationUser = NotificationUser.builder()
                .id(ID)
                .notificationToken(NOTIFICATION_TOKEN)
                .build();

        val notificationUsers = singletonList(notificationUser);

        val tokens = singletonList(NOTIFICATION_TOKEN);

        when(notificationRepository.getAll()).thenReturn(notificationUsers);

        val notification = Notification.builder()
                .title(TITLE)
                .message(MESSAGE)
                .build();

        notificationService.send(notification);

        verify(notificationRepository).getAll();
        verify(firebaseService).sendNotification(notification, tokens);
    }
}
