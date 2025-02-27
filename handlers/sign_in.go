package handlers

import (
	"SE_drive_backend/functions"
	"database/sql"
	"fmt"

	"SE_drive_backend/models"
	"encoding/json"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	CORSFix(w, r)
	//define required datasets .
	var signUpRequestDetails models.SignUpRequestModel
	var signInFailure models.ErrorsModel
	err := json.NewDecoder(r.Body).Decode(&signUpRequestDetails)
	if err != nil {
		signInFailure = functions.SetErrorModel(http.StatusBadRequest, "Invalid JSON format. Invalid SignIn.")
		json.NewEncoder(w).Encode(signInFailure)
		return
	}
	//db connection .
	db, err := functions.DbConnect(w)
	if err != nil {
		signInFailure = functions.SetErrorModel(http.StatusBadRequest, "Failed Db connection during SignIn.Bad Request")
		json.NewEncoder(w).Encode(signInFailure)
		return
	}
	// create table if not exists UserInfoTable(
	// 	email varchar(60) Primary key,
	// 	userName varchar(60) not null,
	// 	userPassword varchar(60) not null,
	// 	isSubscribed bool not null,
	// 	token varchar(50) unique not null
	//);
	//============_________===========
	//token
	token := functions.GenerateUUID()
	var dbNameCheck string
	userCheckQuery := `SELECT email FROM UserInfoTable WHERE email=(?);`
	userInsertQuery := `Insert into UserInfoTable(email,userName,userPassword,isSubscribed,token) values (?,?,?,false,?);`

	er := db.QueryRow(userCheckQuery, signUpRequestDetails.Email).Scan(&dbNameCheck)
	if er != nil {
		if er == sql.ErrNoRows {
			//no rows meaning user doesnot exist .
			_, err = db.Exec(userInsertQuery, signUpRequestDetails.Email, signUpRequestDetails.Name, signUpRequestDetails.Password, token)
			if err != nil {
				signInFailure = functions.SetErrorModel(http.StatusBadRequest, "error while insertion in db during SignIn.")

			}
			//-successful reponse .
			signUpResponseDetails := models.SignUpResponseModel{Message: "SignIn Successful.", TokenId: token}
			json.NewEncoder(w).Encode(signUpResponseDetails)

		} else {
			signInFailure = functions.SetErrorModel(http.StatusBadRequest, fmt.Sprintf("error occured during signin. %s", er))
		}

		json.NewEncoder(w).Encode(signInFailure)
		return
	}
	signInFailure = functions.SetErrorModel(http.StatusBadRequest, "Username already exists.")
	json.NewEncoder(w).Encode(signInFailure)

}
