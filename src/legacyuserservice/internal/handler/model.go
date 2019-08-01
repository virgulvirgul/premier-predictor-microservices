package handler

type legacyIdRequest struct {
	Id int `json:"id"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
