package models

type SignUpRequestModel struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type SignUpResponseModel struct {
	TokenId string `json:"token_id"`
	Message string `json:"message"`
}
