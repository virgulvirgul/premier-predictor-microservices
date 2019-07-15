package com.cshep4.premierpredictor.repository.mongo

import com.cshep4.premierpredictor.config.MongoConfig.Companion.ID
import com.cshep4.premierpredictor.data.SignUpUser
import com.cshep4.premierpredictor.data.User
import com.cshep4.premierpredictor.entity.UserEntity
import com.cshep4.premierpredictor.repository.mongo.UserRepository.Companion.COLLECTION
import com.cshep4.premierpredictor.repository.mongo.UserRepository.Companion.DATABASE
import com.cshep4.premierpredictor.repository.mongo.UserRepository.Companion.EMAIL
import com.cshep4.premierpredictor.utils.TestUtils
import com.mongodb.client.MongoClient
import com.mongodb.client.model.Filters
import com.mongodb.client.model.FindOneAndUpdateOptions
import org.bson.Document
import org.bson.types.ObjectId
import org.hamcrest.CoreMatchers.`is`
import org.hamcrest.CoreMatchers.nullValue
import org.hamcrest.MatcherAssert.assertThat
import org.junit.After
import org.junit.Before
import org.junit.Test
import org.springframework.test.util.ReflectionTestUtils

internal class UserRepositoryTest {
    companion object {
        val ID_1 = ObjectId()

        const val EMAIL_1 = "✉️"
        const val SIGNATURE = "✍️"
        const val PASSWORD = "🛡"
    }
    private lateinit var userRepository: UserRepository

    private lateinit var mongoClient: MongoClient

    @Before
    fun init() {
        mongoClient = TestUtils.mongoClient()

        userRepository = UserRepository()
        ReflectionTestUtils.setField(userRepository, "client", mongoClient)
    }

    @Test(expected = IllegalArgumentException::class)
    fun `'findById' will throw IllegalArgumentException if id is not valid`() {
        userRepository.findById("invalid id 😰")
    }

    @Test
    fun `'findById' will get a user for given id`() {
        val entity = UserEntity(id = ID_1, email = EMAIL_1)

        val opts = FindOneAndUpdateOptions()
        opts.upsert(true)

        mongoClient.getDatabase(DATABASE)
                .getCollection(UserRepository.COLLECTION, UserEntity::class.java)
                .findOneAndUpdate(
                Filters.eq(ID, entity.id),
                Document(
                        mapOf(
                                Pair("\$set", entity)
                        )
                ),
                opts
        )

        val result = userRepository.findById(ID_1.toHexString())

        assertThat(result, `is`(entity.toDto()))
    }

    @Test
    fun `'findById' will return null for given id if none exist`() {
        val result = userRepository.findById(ID_1.toHexString())

        assertThat(result, `is`(nullValue()))
    }

    @Test
    fun `'findByEmail' will get a user for given email`() {
        val entity = UserEntity(id = ID_1, email = EMAIL_1)

        val opts = FindOneAndUpdateOptions()
        opts.upsert(true)

        mongoClient.getDatabase(DATABASE)
                .getCollection(UserRepository.COLLECTION, UserEntity::class.java)
                .findOneAndUpdate(
                        Filters.eq(ID, entity.id),
                        Document(
                                mapOf(
                                        Pair("\$set", entity)
                                )
                        ),
                        opts
                )

        val result = userRepository.findByEmail(EMAIL_1)

        assertThat(result, `is`(entity.toDto()))
    }

    @Test
    fun `'findByEmail' will return null for given email if none exist`() {
        val result = userRepository.findByEmail(EMAIL_1)

        assertThat(result, `is`(nullValue()))
    }

    @Test
    fun `'setUserSignature' will set signature for user with given email and return number of modified documents`() {
        val entity = UserEntity(id = ID_1, email = EMAIL_1)

        val opts = FindOneAndUpdateOptions()
        opts.upsert(true)

        mongoClient.getDatabase(DATABASE)
                .getCollection(UserRepository.COLLECTION, UserEntity::class.java)
                .findOneAndUpdate(
                        Filters.eq(ID, entity.id),
                        Document(
                                mapOf(
                                        Pair("\$set", entity)
                                )
                        ),
                        opts
                )

        val result = userRepository.setUserSignature(SIGNATURE, EMAIL_1)

        assertThat(result, `is`(1))

        val storedUser = mongoClient.getDatabase(DATABASE)
                .getCollection(UserRepository.COLLECTION, UserEntity::class.java)
                .find(Filters.eq(EMAIL, EMAIL_1))
                .first()!!

        assertThat(storedUser.signature, `is`(SIGNATURE))
    }

    @Test
    fun `'setUserSignature' will return 0 if no user is found for given email`() {
        val result = userRepository.setUserSignature(SIGNATURE, EMAIL_1)

        assertThat(result, `is`(0))
    }

    @Test
    fun `'resetUserPassword' will set password for user with given email and signature and return number of modified documents`() {
        val entity = UserEntity(id = ID_1, email = EMAIL_1, signature = SIGNATURE)

        val opts = FindOneAndUpdateOptions()
        opts.upsert(true)

        mongoClient.getDatabase(DATABASE)
                .getCollection(UserRepository.COLLECTION, UserEntity::class.java)
                .findOneAndUpdate(
                        Filters.eq(ID, entity.id),
                        Document(
                                mapOf(
                                        Pair("\$set", entity)
                                )
                        ),
                        opts
                )

        val result = userRepository.resetUserPassword(PASSWORD, EMAIL_1, SIGNATURE)

        assertThat(result, `is`(1))

        val storedUser = mongoClient.getDatabase(DATABASE)
                .getCollection(UserRepository.COLLECTION, UserEntity::class.java)
                .find(Filters.eq(EMAIL, EMAIL_1))
                .first()!!

        assertThat(storedUser.signature, `is`(SIGNATURE))
    }

    @Test
    fun `'resetUserPassword' will return 0 if no user is found for given email and signature`() {
        val entity = UserEntity(id = ID_1, email = EMAIL_1)

        val opts = FindOneAndUpdateOptions()
        opts.upsert(true)

        mongoClient.getDatabase(DATABASE)
                .getCollection(COLLECTION, UserEntity::class.java)
                .findOneAndUpdate(
                        Filters.eq(ID, entity.id),
                        Document(
                                mapOf(
                                        Pair("\$set", entity)
                                )
                        ),
                        opts
                )

        val result = userRepository.resetUserPassword(PASSWORD, EMAIL_1, SIGNATURE)

        assertThat(result, `is`(0))
    }

    @Test
    fun `'save' will store a user to the database`() {
        val user = User(id = ID_1.toHexString())

        userRepository.save(user)

        val storedUser = mongoClient.getDatabase(DATABASE)
                .getCollection(COLLECTION, UserEntity::class.java)
                .find(Filters.eq(ID, ID_1))
                .first()
                ?.toDto()

        assertThat(storedUser, `is`(user))
    }

    @Test
    fun `'save' will update a user in the database if it exists`() {
        val user = User(id = ID_1.toHexString())

        userRepository.save(user)

        user.firstName = "first name"
        user.surname = "surname"

        userRepository.save(user)

        val storedUser = mongoClient.getDatabase(DATABASE)
                .getCollection(COLLECTION, UserEntity::class.java)
                .find(Filters.eq(ID, ID_1))
                .first()
                ?.toDto()!!

        assertThat(storedUser, `is`(user))
        assertThat(storedUser.firstName, `is`(user.firstName))
        assertThat(storedUser.surname, `is`(user.surname))
    }

    @Test
    fun `'save' will store initial sign up user to database`() {
        val user = SignUpUser(
                firstName = "",
                surname = "name",
                email = "email",
                password = "Pass123",
                confirmPassword = "Pass123",
                predictedWinner = "Liverpool"
        )

        val result = userRepository.save(user)

        val storedUser = mongoClient.getDatabase(DATABASE)
                .getCollection(COLLECTION, UserEntity::class.java)
                .find()
                .first()!!
                .toDto()

        assertThat(storedUser.id, `is`(result!!.id))
        assertThat(storedUser.firstName, `is`(user.firstName))
        assertThat(storedUser.surname, `is`(user.surname))
        assertThat(storedUser.email, `is`(user.email))
        assertThat(storedUser.password, `is`(user.password))
        assertThat(storedUser.predictedWinner, `is`(user.predictedWinner))
    }

    @After
    fun tearDown() {
        mongoClient.getDatabase(DATABASE)
                .drop()

        mongoClient.close()
    }
}