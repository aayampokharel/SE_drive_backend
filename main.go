package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	router.Post("/uploadphoto", func(w http.ResponseWriter, r *http.Request) {
		//Cors
		CORSFix(w, r)
		uploadPhoto(w, r)
	})

	http.ListenAndServe(":8000", router)
}
