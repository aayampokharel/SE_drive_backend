package models

type SignOutRequestModel struct {
	Token string `json:"token_id"`
}

type SignOutResponseModel struct {
	Message    string `json:"messsage"`
	StatusCode string `json:"statuscode"`
}
