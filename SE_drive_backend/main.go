package main

import (
	"SE_drive_backend/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	router := chi.NewRouter()

	// Apply CORS middleware
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, //
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Define routes
	router.Post("/signin", handlers.SignIn)
	router.Post("/login", handlers.Login)
	router.Post("/uploadphotos", uploadPhoto)
	router.Post("/uploadvideos", uploadVideo)
	router.Post("/uploadaudios", uploadAudio)
	router.Post("/uploadtexts", uploadText)
	router.Post("/uploadpdfs", uploadPdf)
	router.Post("/getsavedmedia", handlers.GetSavedMedia)
	router.Post("/signout", handlers.SignOut)

	// Start the server
	http.ListenAndServe(":8000", router)
}
