package global

import "SE_drive_backend/models"

//history map should update afterwards by adding the filename directly here ,
//dutai filename huncha inside uploadAudio,uploadVideo ,and esto methods chalnu bhanda agadi map is ever ready .
//! send token forupload as well . based on it add it to the map, and store in db as well .
var MediaMap = make(map[string]models.MediaMap)

//defining map of structure:
//{
//"1234-5678-910":{
//  "listOfVideos":["",""],
// "listOfPhotos":["",""],
//        }
//
//
//}

func MediaMapEntry(tokenKey string, mediaMapModelToAdd models.MediaMap) {
	//listOfMap structure : key:{
	// listOfVideos:[],
	//listOfPhotos:[],*3 more .
	//}
	if val, ok := MediaMap[tokenKey]; ok {
		// val represents inside map of MediaMap.
		val.VideosList = append(val.VideosList, mediaMapModelToAdd.VideosList...)
		val.PhotosList = append(val.VideosList, mediaMapModelToAdd.PhotosList...)
		val.AudiosList = append(val.VideosList, mediaMapModelToAdd.AudiosList...)
		val.PdfsList = append(val.VideosList, mediaMapModelToAdd.PdfsList...)
		val.TextsList = append(val.VideosList, mediaMapModelToAdd.TextsList...)
		MediaMap[tokenKey] = val
	} else {
		MediaMap[tokenKey] = mediaMapModelToAdd
	}

}
