package models

type SignUpRequestModel struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type SignUpResponseModel struct {
	Message       string   `json:"message"`
	MediaMapModel MediaMap `json:"map_model"`
}
