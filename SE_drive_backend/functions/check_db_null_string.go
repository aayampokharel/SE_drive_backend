package functions

import (
	"SE_drive_backend/models"
	"database/sql"
)

func CheckDbNullString(dbNullStringType *sql.NullString) string {
	if dbNullStringType.Valid {
		return dbNullStringType.String

	} else {
		return ""
	}
}

// ! can use pointer for parameter .
func CheckDbNullStringAndReturnMap(loginDbModel models.LogInDbModel, mediaMapPtr *models.MediaMap) *models.MediaMap {
	//var *mediaMapPtr models.MediaMap
	pdfFileName := CheckDbNullString(&loginDbModel.PdfFileName)
	audioFileName := CheckDbNullString(&loginDbModel.AudioFileName)
	photoFileName := CheckDbNullString(&loginDbModel.PhotoFileName)
	videoFileName := CheckDbNullString(&loginDbModel.VideoFileName)
	textFileName := CheckDbNullString(&loginDbModel.TextFileName)
	//-------------------------appending ----------------------||
	mediaMapPtr.AudiosList = append(mediaMapPtr.AudiosList, audioFileName)
	mediaMapPtr.PdfsList = append(mediaMapPtr.PdfsList, pdfFileName)
	mediaMapPtr.PhotosList = append(mediaMapPtr.PhotosList, photoFileName)
	mediaMapPtr.VideosList = append(mediaMapPtr.VideosList, videoFileName)
	mediaMapPtr.TextsList = append(mediaMapPtr.TextsList, textFileName)

	return mediaMapPtr
}
