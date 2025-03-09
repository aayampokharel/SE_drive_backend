package global

import (
	"SE_drive_backend/models"
	"fmt"
	"net/http"
)

//var AddedMediaMap = make(map[string]*models.AddedMediaMap)

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

		MediaMap[tokenKey].PhotosList = append(MediaMap[tokenKey].PhotosList, singleMedia)
		fmt.Print(MediaMap[tokenKey])
		// MediaMapEntry(tokenKey, models.MediaMap{PhotosList: []string{singleMedia}})
		// fmt.Print(global.MediaMap[tokenKey])
	case "Video":
		// MediaMap[tokenKey].VideosList = append(MediaMap[tokenKey].VideosList, singleMedia)
		MediaMapEntry(tokenKey, models.MediaMap{VideosList: []string{singleMedia}})
	case "Audio":
		// MediaMap[tokenKey].AudiosList = append(MediaMap[tokenKey].AudiosList, singleMedia)
		MediaMapEntry(tokenKey, models.MediaMap{AudiosList: []string{singleMedia}})

	case "Pdf":
		// MediaMap[tokenKey].PdfsList = append(MediaMap[tokenKey].PdfsList, singleMedia)
		MediaMapEntry(tokenKey, models.MediaMap{PdfsList: []string{singleMedia}})
	case "Text":
		// MediaMap[tokenKey].TextsList = append(MediaMap[tokenKey].PdfsList, singleMedia)
		MediaMapEntry(tokenKey, models.MediaMap{TextsList: []string{singleMedia}})

	}
	return models.ErrorsModel{}, true
}
