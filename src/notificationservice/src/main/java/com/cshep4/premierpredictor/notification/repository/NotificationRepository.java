package com.cshep4.premierpredictor.notification.repository;

import com.cshep4.premierpredictor.notification.entity.NotificationUserEntity;
import com.cshep4.premierpredictor.notification.model.NotificationUser;
import com.mongodb.client.MongoClient;
import com.mongodb.client.MongoCollection;
import com.mongodb.client.MongoDatabase;
import lombok.val;
import org.bson.types.ObjectId;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;

import java.util.ArrayList;
import java.util.List;

import static com.mongodb.client.model.Filters.eq;
import static com.mongodb.client.model.Filters.in;
import static java.util.stream.Collectors.toList;

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
                .find(eq("_id", new ObjectId(id)))
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

    public void save(NotificationUser notification) {
        val entity = NotificationUserEntity.fromDto(notification);

        collection().insertOne(entity);
    }
}
