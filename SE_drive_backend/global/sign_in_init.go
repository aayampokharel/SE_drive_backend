package global

import "SE_drive_backend/models"

func SignInInit(tokenKey string, email string) {
	MediaMap[tokenKey] = &models.MediaMap{}

	MediaMap[tokenKey].Email = email
	MediaMap[tokenKey].Token = tokenKey

	MediaMap[tokenKey].IsSubscribed = false
	MediaMap[tokenKey].TrialsLeft = 10         //initial value for first signin .
	MediaMapEntry(tokenKey, models.MediaMap{}) //intialize the token in the map
}
