package models

//@ CHANGE TYPE OF LOGINMODEL AND ADD MAP FOR LIST OF VIDFEOS ,PHOTOS, ETC .
type MediaMap struct {
	VideosList []string
	PhotosList []string
	AudiosList []string
	TextsList  []string
	PdfsList   []string
}
