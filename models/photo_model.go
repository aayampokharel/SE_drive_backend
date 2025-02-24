package models

type PhotoRequestModel struct {
	Token string `json:"token_id"`
}

//to include the photofile as well . as it is streamed so not included in struct portion .
