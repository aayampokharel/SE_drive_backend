package models

//@ CHANGE TYPE OF LOGINMODEL AND ADD MAP FOR LIST OF VIDFEOS ,PHOTOS, ETC .
type MediaMap struct {
	Email      string   `json:"email"`
	Token      string   `json:"token_id"`
	VideosList []string `json:"videos_list"`
	PhotosList []string `json:"photos_list"`
	AudiosList []string `json:"audios_list"`
	TextsList  []string `json:"texts_list"`
	PdfsList   []string `json:"pdfs_list"`
}
