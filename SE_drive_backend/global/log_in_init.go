package global

import "SE_drive_backend/models"

func LogInInit(tokenKey string, trialsLeft int, email string, isSubscribed bool, modelsMediaMap *models.MediaMap) {
	if _, ok := MediaMap[tokenKey]; !ok {

		MediaMap[tokenKey] = modelsMediaMap //here the media is initialized/non-duplicated one .
		MediaMap[tokenKey].Email = email
		MediaMap[tokenKey].Token = tokenKey
		MediaMap[tokenKey].IsSubscribed = isSubscribed
		MediaMap[tokenKey].TrialsLeft = trialsLeft

	}
}
