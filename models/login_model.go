package models

type LogInRequestModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogInResponseModel struct {
	MessageStatus string
}
