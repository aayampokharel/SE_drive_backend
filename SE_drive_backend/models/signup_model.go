package models

type SignUpRequestModel struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type SignUpResponseModel struct {
	TokenId      string `json:"token_id"`
	IsSubscribed bool   `json:"is_subscribed"`
	TrialsLeft   int    `json:"trials_left"`
	Message      string `json:"message"`
}
