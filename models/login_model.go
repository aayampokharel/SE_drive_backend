package models

import "database/sql"

type LogInRequestModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogInResponseModel struct {
	MessageStatus string
}

type LogInDbModel struct {
	Token         sql.NullString
	UserName      string
	IsSubscribed  bool
	VideoFileName sql.NullString
	PhotoFileName sql.NullString
	PdfFileName   sql.NullString
	AudioFileName sql.NullString
	TextFileName  sql.NullString
}
