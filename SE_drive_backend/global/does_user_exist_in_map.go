package global

import (
	"SE_drive_backend/models"
)

func DoesUserExistInMap(userMail string) (mapModelValue models.MediaMap, ok bool) {
	for _, mapModelValue := range MediaMap {
		if (*mapModelValue).Email == userMail {
			return *mapModelValue, true
		}
	}
	return models.MediaMap{}, false
}
