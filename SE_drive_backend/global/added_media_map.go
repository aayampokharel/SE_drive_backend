package global

import (
	"SE_drive_backend/functions"
	"SE_drive_backend/models"
	"net/http"
)

func AddNewMedia(tokenKey string, singleMedia string, types string) (models.ErrorsModel, bool) {

	//@ deduce trail ,
	//@ makie login .

	if !MediaMap[tokenKey].IsSubscribed {
		if MediaMap[tokenKey].TrialsLeft <= 0 {

			return models.ErrorsModel{Err: "Trials period is over . Subscribe. ", StatusCode: http.StatusBadRequest}, false
		} else {
			MediaMap[tokenKey].TrialsLeft = MediaMap[tokenKey].TrialsLeft - 1
		}

	}
	switch types {
	case "Photo":

		MediaMap[tokenKey].PhotosList =
			functions.RemoveDuplicatesFromList(append(MediaMap[tokenKey].PhotosList, singleMedia))

	case "Video":
		MediaMap[tokenKey].VideosList = functions.RemoveDuplicatesFromList(append(MediaMap[tokenKey].VideosList, singleMedia))

	case "Audio":
		MediaMap[tokenKey].AudiosList = functions.RemoveDuplicatesFromList(append(MediaMap[tokenKey].AudiosList, singleMedia))

	case "Pdf":
		MediaMap[tokenKey].PdfsList = functions.RemoveDuplicatesFromList(append(MediaMap[tokenKey].PdfsList, singleMedia))

	case "Text":
		MediaMap[tokenKey].TextsList = functions.RemoveDuplicatesFromList(append(MediaMap[tokenKey].PdfsList, singleMedia))

	}
	return models.ErrorsModel{}, true
}
