package global

import "SE_drive_backend/models"

func SignInInit(tokenKey string) {
	AddedMediaMap[tokenKey] = &models.AddedMediaMap{}
	AddedMediaMap[tokenKey].IsSubscribed = false
	AddedMediaMap[tokenKey].TrialsLeft = 10    //initial value for first signin .
	MediaMapEntry(tokenKey, models.MediaMap{}) //intialize the token in the map
}
