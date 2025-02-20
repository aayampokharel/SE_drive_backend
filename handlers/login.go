package handlers

import (
	"SE_drive_backend/functions"
	"SE_drive_backend/models"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func login(w http.ResponseWriter, r *http.Request) {
	var logInModel models.LogInModel

	// Decode the incoming JSON request into logInModel
	err := json.NewDecoder(r.Body).Decode(&logInModel)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	dsn := functions.GetDSN()         //gets the required credentials.
	db, err := sql.Open("mysql", dsn) // Connect to the MySQL database using the credentials.
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close() //@ necessary always .

	query := "SELECT id FROM users WHERE email = ? AND password = ?"
	row := db.QueryRow(query, logInModel.Email, logInModel.Password)

	// Check if the query returns a result
	var userID int
	err = row.Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			// If no rows are returned, email or password is incorrect
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		} else {
			// Other error (e.g., query failure)
			http.Error(w, "Error querying the database", http.StatusInternalServerError)
		}
		return
	}

	// Successfully logged in
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"status": "success", "user_id": string(userID)}
	json.NewEncoder(w).Encode(response)
}
