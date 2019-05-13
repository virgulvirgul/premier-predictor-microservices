package com.cshep4.premierpredictor.notification.service;

import com.cshep4.premierpredictor.notification.model.*;
import com.cshep4.premierpredictor.notification.repository.NotificationRepository;
import io.reactivex.Observer;
import lombok.SneakyThrows;
import lombok.val;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.MockitoJUnitRunner;

import static com.cshep4.premierpredictor.notification.constant.Constants.NOTIFICATION_LIMIT;
import static java.util.Collections.singletonList;
import static org.hamcrest.CoreMatchers.is;
import static org.hamcrest.MatcherAssert.assertThat;
import static org.mockito.Mockito.*;

@RunWith(MockitoJUnitRunner.class)
public class NotificationServiceTest {
    private static final String ID = "id";
    private static final String NOTIFICATION_ID = "id";
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

        verify(notificationRepository).saveUser(notificationUser);
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
        verify(notificationRepository).saveNotification(notification, ID);
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
        verify(notificationRepository).saveNotification(notification, ID);
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

    @Test
    public void getNotificationsGetsUsersRecentNotifications() {
        val notifications = new CircularQueue<Notification>(NOTIFICATION_LIMIT);

        when(notificationRepository.getRecentNotifications(ID)).thenReturn(notifications);

        val result = notificationService.getNotifications(ID);

        verify(notificationRepository).getRecentNotifications(ID);
        assertThat(result, is(notifications));
    }

    @Test
    public void updateReadNotificationUpdatesDb() {
        notificationService.updateReadNotification(ID, NOTIFICATION_ID);

        verify(notificationRepository).updateReadNotification(ID, NOTIFICATION_ID);
    }

    @Test
    public void subscribeForUpdates() {
        Observer<Notification> notificationObserver = mock(Observer.class);
        notificationService.subscribeToUpdates(ID, notificationObserver);

        verify(notificationRepository).subscribeToUpdates(ID, notificationObserver);
    }
}
