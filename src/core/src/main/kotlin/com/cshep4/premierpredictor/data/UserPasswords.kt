package com.cshep4.premierpredictor.data

data class UserPasswords (val id: String = "", var oldPassword: String = "", var newPassword: String = "", var confirmPassword: String = "")