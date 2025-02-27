package handlers

import (
	"net/http"
)

func GetSavedPhotos(w http.ResponseWriter, r *http.Request) {
	CORSFix(w, r)
	//use stored values from map
	//call this after data available in frontend from log_in.go only after that .

}
