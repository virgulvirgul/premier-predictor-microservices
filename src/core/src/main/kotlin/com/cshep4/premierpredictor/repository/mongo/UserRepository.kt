package com.cshep4.premierpredictor.repository.mongo

import com.cshep4.premierpredictor.config.MongoConfig.Companion.ID
import com.cshep4.premierpredictor.data.SignUpUser
import com.cshep4.premierpredictor.data.User
import com.cshep4.premierpredictor.entity.UserEntity
import com.mongodb.client.MongoClient
import com.mongodb.client.MongoCollection
import com.mongodb.client.MongoDatabase
import com.mongodb.client.model.Filters.and
import com.mongodb.client.model.Filters.eq
import com.mongodb.client.model.FindOneAndUpdateOptions
import com.mongodb.client.model.Updates.set
import org.bson.Document
import org.bson.types.ObjectId
import org.bson.types.ObjectId.isValid
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Repository

@Repository
class UserRepository {
    companion object {
        const val DATABASE = "user"
        const val COLLECTION = "user"

        const val EMAIL = "email"
        const val SIGNATURE = "signature"
        const val PASSWORD = "password"
    }

    @Autowired
    private lateinit var client: MongoClient

    private fun database(): MongoDatabase {
        return client.getDatabase(DATABASE)
    }

    private fun collection(): MongoCollection<UserEntity> {
        return database().getCollection(COLLECTION, UserEntity::class.java)
    }

    fun findById(id: String): User? {
        if (!isValid(id)) {
            throw IllegalArgumentException("Invalid id")
        }

        return collection()
                .find(eq(ID, ObjectId(id)))
                .first()
                ?.toDto()
    }

    fun findByEmail(email: String): User? {
        return collection()
                .find(eq(EMAIL, email))
                .first()
                ?.toDto()
    }

    fun setUserSignature(signature: String, email: String): Int {
        return collection()
                .updateOne(
                        eq(EMAIL, email),
                        set(SIGNATURE, signature)
                )
                .modifiedCount
                .toInt()
    }

    fun resetUserPassword(password: String, email: String, signature: String): Int {
        return collection()
                .updateOne(
                        and(
                                eq(EMAIL, email),
                                eq(SIGNATURE, signature)
                        ),
                        set(PASSWORD, password)
                )
                .modifiedCount
                .toInt()
    }

    fun save(user: User) {
        val entity = UserEntity.fromDto(user)

        val opts = FindOneAndUpdateOptions()
        opts.upsert(true)

        collection().findOneAndUpdate(
                eq(ID, entity.id),
                Document(
                        mapOf(
                                Pair("\$set", entity)
                        )
                ),
                opts
        )
    }

    fun save(user: SignUpUser): User? {
        val entity = UserEntity.fromDto(user)

        collection().insertOne(entity)

        return entity.toDto()
    }
}