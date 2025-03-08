package global

import (
	"SE_drive_backend/models"
	"encoding/json"
	"net/http"
)

var AddedMediaMap = make(map[string]*models.AddedMediaMap)

func AddNewMedia(w http.ResponseWriter, tokenKey string, singleMedia string, types string) {

	//@ deduce trail ,
	//@ makie login .

	if !AddedMediaMap[tokenKey].IsSubscribed {
		if AddedMediaMap[tokenKey].TrialsLeft <= 0 {
			json.NewEncoder(w).Encode(models.ErrorsModel{StatusCode: http.StatusInternalServerError, Err: "Trials period is over . Subscribe. "})
			return
		} else {
			AddedMediaMap[tokenKey].TrialsLeft = AddedMediaMap[tokenKey].TrialsLeft - 1
		}

	}
	switch types {
	case "Photo":
		//syntactic sugar for dereference .
		AddedMediaMap[tokenKey].PhotosList = append(AddedMediaMap[tokenKey].PhotosList, singleMedia)
		MediaMapEntry(tokenKey, models.MediaMap{PhotosList: []string{singleMedia}})
	case "Video":
		AddedMediaMap[tokenKey].VideosList = append(AddedMediaMap[tokenKey].VideosList, singleMedia)
		MediaMapEntry(tokenKey, models.MediaMap{VideosList: []string{singleMedia}})
	case "Audio":
		AddedMediaMap[tokenKey].AudiosList = append(AddedMediaMap[tokenKey].AudiosList, singleMedia)
		MediaMapEntry(tokenKey, models.MediaMap{AudiosList: []string{singleMedia}})

	case "Pdf":
		AddedMediaMap[tokenKey].PdfsList = append(AddedMediaMap[tokenKey].PdfsList, singleMedia)
		MediaMapEntry(tokenKey, models.MediaMap{PdfsList: []string{singleMedia}})
	case "Text":
		AddedMediaMap[tokenKey].TextsList = append(AddedMediaMap[tokenKey].PdfsList, singleMedia)
		MediaMapEntry(tokenKey, models.MediaMap{TextsList: []string{singleMedia}})

	}

}
