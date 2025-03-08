package global

import "SE_drive_backend/models"

func LogInInit(tokenKey string, trialsLeft int, isSubscribed bool) {
	if _, ok := AddedMediaMap[tokenKey]; !ok {

		AddedMediaMap[tokenKey] = &models.AddedMediaMap{}
		AddedMediaMap[tokenKey].IsSubscribed = isSubscribed
		AddedMediaMap[tokenKey].TrialsLeft = trialsLeft

	}
}
