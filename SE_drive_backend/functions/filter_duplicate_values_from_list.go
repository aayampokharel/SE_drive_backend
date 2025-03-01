package functions

import "SE_drive_backend/models"

func RemoveDuplicatesFromList(strList []string) []string {
	// Create a map to track seen elements
	seen := make(map[string]struct{})
	list := []string{}
	for _, value := range strList {
		seen[value] = struct{}{}
	}
	for key := range seen {
		list = append(list, key)
	}
	return list
}

func RemoveDuplicatesFromMapModel(mapModel models.MediaMap) models.MediaMap {
	mapModel.AudiosList = RemoveDuplicatesFromList(mapModel.AudiosList)
	mapModel.VideosList = RemoveDuplicatesFromList(mapModel.VideosList)
	mapModel.PdfsList = RemoveDuplicatesFromList(mapModel.PdfsList)
	mapModel.PhotosList = RemoveDuplicatesFromList(mapModel.PhotosList)
	mapModel.TextsList = RemoveDuplicatesFromList(mapModel.TextsList)

	return mapModel

}
