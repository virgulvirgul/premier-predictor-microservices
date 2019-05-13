package com.cshep4.premierpredictor.notification.repository;

import com.cshep4.premierpredictor.notification.entity.NotificationUserEntity;
import com.cshep4.premierpredictor.notification.model.CircularQueue;
import com.cshep4.premierpredictor.notification.model.Notification;
import com.cshep4.premierpredictor.notification.model.NotificationUser;
import com.mongodb.ConnectionString;
import com.mongodb.MongoClientSettings;
import com.mongodb.client.MongoClient;
import com.mongodb.client.MongoClients;
import lombok.val;
import org.bson.codecs.pojo.PojoCodecProvider;
import org.bson.types.ObjectId;
import org.junit.After;
import org.junit.Before;
import org.junit.Test;
import org.springframework.test.util.ReflectionTestUtils;

import static com.cshep4.premierpredictor.notification.constant.Constants.NOTIFICATION_LIMIT;
import static com.cshep4.premierpredictor.notification.repository.NotificationRepository.COLLECTION;
import static com.cshep4.premierpredictor.notification.repository.NotificationRepository.DATABASE;
import static com.mongodb.client.model.Filters.eq;
import static java.util.Arrays.asList;
import static org.bson.codecs.configuration.CodecRegistries.fromProviders;
import static org.bson.codecs.configuration.CodecRegistries.fromRegistries;
import static org.hamcrest.CoreMatchers.*;
import static org.hamcrest.MatcherAssert.assertThat;

public class NotificationRepositoryTest {
    private static final String ID = new ObjectId().toHexString();
    private static final String NOTIFICATION_KEY = "notificationToken";
    private static final int NUM_NOTIFICATIONS = 50;

    private NotificationRepository notificationRepository;

    private MongoClient client;

    @Before
    public void init() {
        val pojoCodecRegistry = fromRegistries(
                MongoClientSettings.getDefaultCodecRegistry(),
                fromProviders(
                        PojoCodecProvider.builder()
                                .automatic(true)
                                .build()
                )
        );

        val settings = MongoClientSettings.builder()
                .codecRegistry(pojoCodecRegistry)
                .applyConnectionString(new ConnectionString("mongodb://localhost:27017"))
                .build();

        client = MongoClients.create(settings);

        notificationRepository = new NotificationRepository();
        ReflectionTestUtils.setField(notificationRepository, "client", client);
    }

    @Test
    public void getByIdRetrievesTheNotificationObject() {
        val notificationUserEntity = NotificationUserEntity.builder()
                .id(new ObjectId(ID))
                .notificationToken(NOTIFICATION_KEY)
                .build();

        client.getDatabase(DATABASE)
                .getCollection(COLLECTION, NotificationUserEntity.class)
                .insertOne(notificationUserEntity);

        val result = notificationRepository.getById(ID);

        assertThat(result.getId(), is(ID));
        assertThat(result.getNotificationToken(), is(NOTIFICATION_KEY));
    }

    @Test
    public void getAllByIdsRetrievesAllTheNotificationUserObjectsWithTheSpecifiedIds() {
        val notificationUserEntity = NotificationUserEntity.builder()
                .id(new ObjectId(ID))
                .notificationToken(NOTIFICATION_KEY)
                .build();

        String id = new ObjectId().toHexString();

        val notificationUserEntity2 = NotificationUserEntity.builder()
                .id(new ObjectId(id))
                .notificationToken(NOTIFICATION_KEY)
                .build();

        client.getDatabase(DATABASE)
                .getCollection(COLLECTION, NotificationUserEntity.class)
                .insertMany(asList(notificationUserEntity, notificationUserEntity2));

        val result = notificationRepository.getAllByIds(asList(ID, id));

        assertThat(result.get(0).getId(), is(ID));
        assertThat(result.get(0).getNotificationToken(), is(NOTIFICATION_KEY));
        assertThat(result.get(1).getId(), is(id));
        assertThat(result.get(1).getNotificationToken(), is(NOTIFICATION_KEY));
    }

    @Test
    public void getAllRetrievesAllTheNotificationUserObjects() {
        val notificationUserEntity = NotificationUserEntity.builder()
                .id(new ObjectId(ID))
                .notificationToken(NOTIFICATION_KEY)
                .build();

        String id = new ObjectId().toHexString();

        val notificationUserEntity2 = NotificationUserEntity.builder()
                .id(new ObjectId(id))
                .notificationToken(NOTIFICATION_KEY)
                .build();

        client.getDatabase(DATABASE)
                .getCollection(COLLECTION, NotificationUserEntity.class)
                .insertMany(asList(notificationUserEntity, notificationUserEntity2));

        val result = notificationRepository.getAll();

        assertThat(result.get(0).getId(), is(ID));
        assertThat(result.get(0).getNotificationToken(), is(NOTIFICATION_KEY));
        assertThat(result.get(1).getId(), is(id));
        assertThat(result.get(1).getNotificationToken(), is(NOTIFICATION_KEY));
    }

    @Test
    public void getByIdReturnsNullIfIdCannotBeFound() {
        val notificationUserEntity = NotificationUserEntity.builder()
                .id(new ObjectId())
                .build();

        client.getDatabase(DATABASE)
                .getCollection(COLLECTION, NotificationUserEntity.class)
                .insertOne(notificationUserEntity);

        val result = notificationRepository.getById(ID);

        assertThat(result, is(nullValue()));
    }

    @Test
    public void saveUserWillSaveObjectInDatabase() {
        val notifications = new CircularQueue<Notification>(NOTIFICATION_LIMIT);
        notifications.offer(Notification.builder().build());

        val notificationUser = NotificationUser.builder()
                .id(ID)
                .notificationToken(NOTIFICATION_KEY)
                .notifications(notifications)
                .build();

        notificationRepository.saveUser(notificationUser);

        val user = client.getDatabase(DATABASE)
                .getCollection(COLLECTION, NotificationUserEntity.class)
                .find(eq("_id", new ObjectId(ID)))
                .first()
                .toDto();

        assertThat(user.getId(), is(ID));
        assertThat(user.getNotificationToken(), is(NOTIFICATION_KEY));
        assertThat(user.getNotifications().peek().getId(), is(notNullValue()));
    }

    @Test
    public void getRecentNotificationsGetsTheUsersLast20Notifications() {
        val notificationUser = NotificationUser.builder()
                .id(ID)
                .notificationToken(NOTIFICATION_KEY)
                .notifications(new CircularQueue<>(NOTIFICATION_LIMIT))
                .build();

        int num = 0;

        for (int i = 0; i < NUM_NOTIFICATIONS; i++) {
            val notification = Notification.builder()
                    .title(String.valueOf(num))
                    .build();

            notificationUser.getNotifications().offer(notification);
            num++;
        }

        client.getDatabase(DATABASE)
                .getCollection(COLLECTION, NotificationUserEntity.class)
                .insertOne(NotificationUserEntity.fromDto(notificationUser));

        val result = notificationRepository.getRecentNotifications(ID);

        int i = NUM_NOTIFICATIONS - NOTIFICATION_LIMIT;

        assertThat(result.size(), is(NOTIFICATION_LIMIT));
        for (val n : result) {
            assertThat(n.getTitle(), is(String.valueOf(i)));
            i++;
        }
    }

    @Test
    public void updateReadNotificationUpdatesTheUsersLastReadNotification() {
        val notificationUser = NotificationUser.builder()
                .id(ID)
                .notificationToken(NOTIFICATION_KEY)
                .build();

        notificationRepository.saveUser(notificationUser);

        val notificationId = new ObjectId().toHexString();

        notificationRepository.updateReadNotification(ID, notificationId);

        val user = client.getDatabase(DATABASE)
                .getCollection(COLLECTION, NotificationUserEntity.class)
                .find(eq("_id", new ObjectId(ID)))
                .first()
                .toDto();

        assertThat(user.getId(), is(ID));
        assertThat(user.getLastReadNotification(), is(notificationId));
    }

    @Test
    public void saveNotificationWillSaveNotificationForTheUser() {
        val notificationUser = NotificationUser.builder()
                .id(ID)
                .notificationToken(NOTIFICATION_KEY)
                .build();

        notificationRepository.saveUser(notificationUser);

        val title = "title";
        val message = "test";
        val notification = Notification.builder()
                .title(title)
                .message(message)
                .build();

        notificationRepository.saveNotification(notification, ID);

        val user = notificationRepository.getById(ID);

        val notificationResult = user.getNotifications().peek();
        assertThat(notificationResult.getId(), is(notNullValue()));
        assertThat(notificationResult.getTitle(), is(title));
        assertThat(notificationResult.getMessage(), is(message));
    }

    @After
    public void tearDown() {
        client.getDatabase(DATABASE)
                .drop();

        client.close();
    }
}
