package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	router.Post("/uploadphotos", func(w http.ResponseWriter, r *http.Request) {
		//Cors

		uploadPhoto(w, r)
	})
	router.Post("/uploadvideos", func(w http.ResponseWriter, r *http.Request) {
		uploadVideo(w, r)
	})
	router.Post("/uploadaudios", func(w http.ResponseWriter, r *http.Request) {
		uploadAudio(w, r)
	})
	router.Post("/uploadtexts", func(w http.ResponseWriter, r *http.Request) {
		uploadText(w, r)
	})
	router.Post("/uploadpdfs", func(w http.ResponseWriter, r *http.Request) {
		uploadPdf(w, r)
	})
	http.ListenAndServe(":8080", router)
}
