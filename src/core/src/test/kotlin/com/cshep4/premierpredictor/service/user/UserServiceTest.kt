package com.cshep4.premierpredictor.service.user

import com.cshep4.premierpredictor.data.SignUpUser
import com.cshep4.premierpredictor.data.User
import com.cshep4.premierpredictor.data.UserDetails
import com.cshep4.premierpredictor.data.UserPasswords
import com.cshep4.premierpredictor.entity.UserEntity
import com.cshep4.premierpredictor.repository.mongo.UserRepository
import com.nhaarman.mockito_kotlin.times
import com.nhaarman.mockito_kotlin.verify
import com.nhaarman.mockito_kotlin.whenever
import org.hamcrest.CoreMatchers.notNullValue
import org.hamcrest.CoreMatchers.nullValue
import org.hamcrest.MatcherAssert.assertThat
import org.junit.Before
import org.junit.Test
import org.junit.runner.RunWith
import org.mockito.InjectMocks
import org.mockito.Mock
import org.mockito.Mockito.any
import org.mockito.junit.MockitoJUnitRunner
import org.springframework.security.core.userdetails.UsernameNotFoundException
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder
import org.hamcrest.CoreMatchers.`is` as Is

@RunWith(MockitoJUnitRunner::class)
internal class UserServiceTest {
    private val email = "test@email.com"
    private val password = "test password"

    @Mock
    lateinit var bCryptPasswordEncoder: BCryptPasswordEncoder

    @Mock
    private lateinit var userRepository: UserRepository

    @InjectMocks
    private lateinit var userService: UserService

    @Before
    fun init() {
        whenever(bCryptPasswordEncoder.matches(any(), any())).thenReturn(true)
    }

    @Test
    fun `'loadUserByUsername' returns User instance from db`() {
        val user = User(email = email, password = password)
        whenever(userRepository.findByEmail(email)).thenReturn(user)

        val result = userService.loadUserByUsername(email)!!

        assertThat(result.username, Is(email))
        assertThat(result.password, Is(password))
    }

    @Test(expected = UsernameNotFoundException::class)
    fun `'loadUserByUsername' throws UsernameNotFoundException if not user is found`() {
        whenever(userRepository.findByEmail(email)).thenReturn(null)

        userService.loadUserByUsername(email)
    }

    @Test
    fun `'createUser' does not add user to db if password is invalid`() {
        val user = SignUpUser(firstName = "first", surname = "name", email = email, password = "invalidpassword", confirmPassword = "invalidpassword", predictedWinner = "France")

        val result = userService.createUser(user)

        assertThat(result, Is(nullValue()))
        verify(userRepository, times(0)).save(user)
    }

    @Test
    fun `'createUser' does not add user to db if passwords don't match`() {
        val user = SignUpUser(firstName = "first", surname = "name", email = email, password = "Pass123", confirmPassword = "Word123", predictedWinner = "France")

        val result = userService.createUser(user)

        assertThat(result, Is(nullValue()))
        verify(userRepository, times(0)).save(user)
    }

    @Test
    fun `'createUser' does not add user to db if email is invalid`() {
        val user = SignUpUser(firstName = "first", surname = "name", email = "invalid email", password = "Pass123", confirmPassword = "Pass123", predictedWinner = "France")

        val result = userService.createUser(user)

        assertThat(result, Is(nullValue()))
        verify(userRepository, times(0)).save(user)
    }

    @Test
    fun `'createUser' does not add user to db if first name is blank`() {
        val user = SignUpUser(firstName = "", surname = "name", email = email, password = "Pass123", confirmPassword = "Pass123", predictedWinner = "France")

        val result = userService.createUser(user)

        assertThat(result, Is(nullValue()))
        verify(userRepository, times(0)).save(user)
    }

    @Test
    fun `'createUser' does not add user to db if surname is blank`() {
        val user = SignUpUser(firstName = "first", surname = "", email = email, password = "Pass123", confirmPassword = "Pass123", predictedWinner = "France")

        val result = userService.createUser(user)

        assertThat(result, Is(nullValue()))
        verify(userRepository, times(0)).save(user)
    }

    @Test
    fun `'createUser' does not add user to db if user already exists with same email`() {
        val user = SignUpUser(firstName = "first", surname = "name", email = email, password = "Pass123", confirmPassword = "Pass123", predictedWinner = "France")

        whenever(userRepository.findByEmail(email)).thenReturn(UserEntity.fromDto(user).toDto())

        val result = userService.createUser(user)

        assertThat(result, Is(nullValue()))
        verify(userRepository, times(0)).save(user)
    }

    @Test
    fun `'createUser' adds user to db`() {
        val user = SignUpUser(firstName = "first", surname = "name", email = email, password = "Pass123", confirmPassword = "Pass123", predictedWinner = "France")

        whenever(userRepository.save(user)).thenReturn(UserEntity.fromDto(user).toDto())
        whenever(userRepository.findByEmail(email)).thenReturn(null)
        whenever(bCryptPasswordEncoder.encode(user.password)).thenReturn(user.password)

        val result = userService.createUser(user)

        assertThat(result!!.firstName, Is("first"))
        assertThat(result.surname, Is("name"))
        assertThat(result.email, Is(email))
        assertThat(result.password, Is("Pass123"))
        assertThat(result.predictedWinner, Is("France"))
        assertThat(result.joined, Is(notNullValue()))
        verify(userRepository).save(user)
    }

    @Test
    fun `'retrieveUserById' should retrieve user if one is found with id`() {
        val user = User()

        whenever(userRepository.findById("1")).thenReturn(user)

        val result = userService.retrieveUserById("1")

        assertThat(result, Is(user))
    }

    @Test
    fun `'retrieveUserById' should return null if no user is found`() {
        whenever(userRepository.findById("1")).thenReturn(null)

        val result = userService.retrieveUserById("1")

        assertThat(result, Is(nullValue()))
    }

    @Test
    fun `'retrieveUserByEmail' should retrieve user if one is found with email`() {
        val user = User()

        whenever(userRepository.findByEmail("test")).thenReturn(user)

        val result = userService.retrieveUserByEmail("test")

        assertThat(result, Is(user))
    }

    @Test
    fun `'retrieveUserByEmail' should return null if no user is found`() {
        whenever(userRepository.findByEmail("test")).thenReturn(null)

        val result = userService.retrieveUserByEmail("test")

        assertThat(result, Is(nullValue()))
    }

    @Test
    fun `'updateUserPassword' does not update user if password is invalid`() {
        val userPasswords = UserPasswords(id = "1", oldPassword = "Qwerty123", newPassword = "test", confirmPassword = "test")
        val user = User(id = "1", password = "oldPassword")

        whenever(userRepository.findById("1")).thenReturn(user)

        val result = userService.updateUserPassword(userPasswords)

        assertThat(result, Is(nullValue()))
        verify(userRepository, times(0)).save(user)
    }

    @Test
    fun `'updateUserPassword' does not update user if new passwords don't match`() {
        val userPasswords = UserPasswords(id = "1", oldPassword = "Qwerty123", newPassword = "Qwerty123", confirmPassword = "Qwrty1")
        val user = User(id = "1", password = "oldPassword")

        whenever(userRepository.findById("1")).thenReturn(user)

        val result = userService.updateUserPassword(userPasswords)

        assertThat(result, Is(nullValue()))
        verify(userRepository, times(0)).save(user)
    }

    @Test
    fun `'updateUserPassword' does not update user if old password is incorrect`() {
        val userPasswords = UserPasswords(id = "1", oldPassword = "Qwerty123", newPassword = "Qwerty123", confirmPassword = "Qwerty123")
        val user = User(id = "1", password = "oldPassword")

        whenever(userRepository.findById("1")).thenReturn(user)
        whenever(bCryptPasswordEncoder.matches(userPasswords.oldPassword, user.password)).thenReturn(false)

        val result = userService.updateUserPassword(userPasswords)

        assertThat(result, Is(nullValue()))
        verify(userRepository, times(0)).save(user)
    }

    @Test
    fun `'updateUserPassword' does not update if user does not exist`() {
        val userPasswords = UserPasswords(id = "1", oldPassword = "Qwerty123", newPassword = "Qwerty123", confirmPassword = "Qwerty123")

        whenever(userRepository.findById("1")).thenReturn(null)

        val result = userService.updateUserPassword(userPasswords)

        assertThat(result, Is(nullValue()))
    }

    @Test
    fun `'updateUserPassword' updates user password`() {
        val userPasswords = UserPasswords(id = "1", oldPassword = "Qwerty12", newPassword = "Qwerty123", confirmPassword = "Qwerty123")
        val user = User(id = "1", email = "test", password = "oldPassword")

        whenever(userRepository.findById("1")).thenReturn(user)
        whenever(bCryptPasswordEncoder.encode(userPasswords.newPassword)).thenReturn("newPassword")

        val result = userService.updateUserPassword(userPasswords)

        assertThat(result, Is(notNullValue()))
        verify(userRepository).save(user)
        assertThat(result!!.password, Is("newPassword"))
    }

    @Test
    fun `'updateUserDetails' does not update user if email is invalid`() {
        val userDetails = UserDetails(id = "1", firstName = "First", surname = "Last", email = "this is an invalid email")
        val user = User(id = "1")

        whenever(userRepository.findById(userDetails.id)).thenReturn(user)

        val result = userService.updateUserDetails(userDetails)

        assertThat(result, Is(nullValue()))
        verify(userRepository, times(0)).save(user)
    }

    @Test
    fun `'updateUserDetails' does not update user if first name is blank`() {
        val userDetails = UserDetails(id = "1", firstName = "", surname = "Last", email = email)
        val user = User(id = "1")

        whenever(userRepository.findById(userDetails.id)).thenReturn(user)
        whenever(userRepository.findByEmail(userDetails.email)).thenReturn(null)

        val result = userService.updateUserDetails(userDetails)

        assertThat(result, Is(nullValue()))
        verify(userRepository, times(0)).save(user)
    }

    @Test
    fun `'updateUserDetails' does not update user if surname is blank`() {
        val userDetails = UserDetails(id = "1", firstName = "First", surname = "", email = email)
        val user = User(id = "1")

        whenever(userRepository.findById(userDetails.id)).thenReturn(user)
        whenever(userRepository.findByEmail(userDetails.email)).thenReturn(null)

        val result = userService.updateUserDetails(userDetails)

        assertThat(result, Is(nullValue()))
        verify(userRepository, times(0)).save(user)
    }

    @Test
    fun `'updateUserDetails' does not update user if user already exists with same email`() {
        val userDetails = UserDetails(id = "1", firstName = "First", surname = "Last", email = email)
        val user = User(id = "5")

        whenever(userRepository.findById(userDetails.id)).thenReturn(user)
        whenever(userRepository.findByEmail(userDetails.email)).thenReturn(user)

        val result = userService.updateUserDetails(userDetails)

        assertThat(result, Is(nullValue()))
    }

    @Test
    fun `'updateUserPassword' updates user details`() {
        val userDetails = UserDetails(id = "1", firstName = "First", surname = "Last", email = email)
        val user = User(id = "1")

        whenever(userRepository.findById(userDetails.id)).thenReturn(user)
        whenever(userRepository.findByEmail(userDetails.email)).thenReturn(null)

        val result = userService.updateUserDetails(userDetails)

        assertThat(result, Is(notNullValue()))
        verify(userRepository).save(user)
    }
}