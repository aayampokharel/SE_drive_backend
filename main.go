package main

import (
	"SE_drive_backend/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	router := chi.NewRouter()

	router.Post("/signin", func(w http.ResponseWriter, r *http.Request) {
		handlers.SignIn(w, r)
	})
	router.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.Login(w, r)
	})

	router.Post("/uploadphotos", func(w http.ResponseWriter, r *http.Request) {
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
	//http.ListenAndServe(":41114", router) //for my testing in android studio
}
