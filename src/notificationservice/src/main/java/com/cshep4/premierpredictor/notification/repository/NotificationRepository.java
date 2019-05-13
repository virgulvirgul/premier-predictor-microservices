package com.cshep4.premierpredictor.notification.repository;

import com.cshep4.premierpredictor.notification.entity.NotificationUserEntity;
import com.cshep4.premierpredictor.notification.model.Notification;
import com.cshep4.premierpredictor.notification.model.NotificationUser;
import com.mongodb.client.MongoClient;
import com.mongodb.client.MongoCollection;
import com.mongodb.client.MongoDatabase;
import com.mongodb.client.model.Aggregates;
import com.mongodb.client.model.changestream.ChangeStreamDocument;
import io.reactivex.Observer;
import lombok.val;
import org.bson.Document;
import org.bson.types.ObjectId;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;

import java.util.ArrayList;
import java.util.List;
import java.util.Queue;
import java.util.function.Consumer;

import static com.mongodb.client.model.Filters.*;
import static java.util.Arrays.asList;
import static java.util.Collections.singletonList;
import static java.util.stream.Collectors.toList;
import static org.bson.types.ObjectId.isValid;

@Repository
public class NotificationRepository {
    private static final String ID = "_id";

    static final String DATABASE = "notification";
    static final String COLLECTION = "notification";

    @Autowired
    private MongoClient client;

    private MongoDatabase database() {
        return client.getDatabase(DATABASE);
    }

    private MongoCollection<NotificationUserEntity> collection() {
        return database().getCollection(COLLECTION, NotificationUserEntity.class);
    }

    public NotificationUser getById(String id) {
        val entity = collection()
                .find(eq(ID, new ObjectId(id)))
                .first();

        if (entity == null) {
            return null;
        }

        return entity.toDto();
    }

    public List<NotificationUser> getAllByIds(List<String> ids) {
        List<NotificationUser> notificationUsers = new ArrayList<>();

        val objectIds = ids.stream()
                .map(ObjectId::new)
                .collect(toList());

        return collection()
                .find(in(ID, objectIds))
                .map(NotificationUserEntity::toDto)
                .into(notificationUsers);
    }

    public List<NotificationUser> getAll() {
        List<NotificationUser> notificationUsers = new ArrayList<>();

        return collection()
                .find()
                .map(NotificationUserEntity::toDto)
                .into(notificationUsers);
    }

    public void saveUser(NotificationUser notification) {
        val entity = NotificationUserEntity.fromDto(notification);

        collection().insertOne(entity);
    }

    public void saveNotification(Notification notification, String userId) {
        val user = getById(userId);

        user.getNotifications().offer(notification);

        val entity = NotificationUserEntity.fromDto(user);

        collection().findOneAndReplace(
                eq(ID, new ObjectId(userId)),
                entity
        );
    }

    public Queue<Notification> getRecentNotifications(String id) {
        return getById(id).getNotifications();
    }

    public void updateReadNotification(String userId, String notificationId) {
        if (!isValid(userId) || !isValid(notificationId)) {
            return;
        }

        collection().updateOne(
                eq(
                        "_id",
                        new ObjectId(userId)
                ),
                new Document(
                        "$set",
                        new Document(
                                "lastReadNotification",
                                new ObjectId(notificationId)
                        )
                )
        );
    }

    public void subscribeToUpdates(String id, Observer<Notification> notificationObserver) {
        Consumer<ChangeStreamDocument<NotificationUserEntity>> streamDocument = (ChangeStreamDocument<NotificationUserEntity> d) -> notificationObserver.onNext(
                d.getFullDocument()
                        .toDto()
                        .getNotifications()
                        .peek()
        );

        collection()
                .watch(
                        singletonList(
                                Aggregates.match(
                                        and(
                                                asList(
                                                        in("operationType", asList("update", "replace")),
                                                        eq("fullDocument._id", new ObjectId(id))
                                                )
                                        )
                                )
                        )
                )
                .forEach(streamDocument);
    }
}
