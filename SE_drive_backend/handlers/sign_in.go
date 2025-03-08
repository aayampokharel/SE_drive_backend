package handlers

import (
	errors "SE_drive_backend/Errors"
	"SE_drive_backend/functions"
	"SE_drive_backend/global"
	"database/sql"
	"fmt"

	"SE_drive_backend/models"
	"encoding/json"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {

	//______define required datasets .
	var signUpRequestDetails models.SignUpRequestModel
	var signInFailure models.ErrorsModel
	//# ____decode
	err := json.NewDecoder(r.Body).Decode(&signUpRequestDetails)
	if err != nil {
		signInFailure = errors.SetErrorModel(http.StatusBadRequest, "Invalid JSON format. Invalid SignIn.")
		json.NewEncoder(w).Encode(signInFailure)
		return
	}
	//#_______db connection .
	db, err := functions.DbConnect(w)
	if err != nil {
		signInFailure = errors.SetErrorModel(http.StatusBadRequest, "Failed Db connection during SignIn.Bad Request")
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
	//# checks if the email is already registered or not ______
	userCheckQuery := `SELECT email FROM UserInfoTable WHERE email=(?);`
	userInsertQuery := `Insert into UserInfoTable(email,userName,userPassword,isSubscribed,token) values (?,?,?,false,?);`

	er := db.QueryRow(userCheckQuery, signUpRequestDetails.Email).Scan(&dbNameCheck)
	if er != nil {
		if er == sql.ErrNoRows {
			//# new user is verified inside here i.e. no existing email in db .
			_, insertErr := db.Exec(userInsertQuery, signUpRequestDetails.Email, signUpRequestDetails.Name, signUpRequestDetails.Password, token)
			if insertErr != nil {
				json.NewEncoder(w).Encode(errors.SetErrorModel(http.StatusInternalServerError, "error while sigin insertion of data into db . "))
			}

			//# below one is necessary as in signin , the person is never subscribed .
			global.SignInInit(token)

			//-successful reponse .
			signUpResponseDetails := models.SignUpResponseModel{Message: "SignIn Successful.", TokenId: token}
			json.NewEncoder(w).Encode(signUpResponseDetails)
			return
			//! error as why 2 different and i think belwo should be user exists and last one should be removed .
		} else {
			signInFailure = errors.SetErrorModel(http.StatusBadRequest, fmt.Sprintf("error occured during signin. %s", er))
			json.NewEncoder(w).Encode(signInFailure)
			return
		}

	}
	signInFailure = errors.SetErrorModel(http.StatusBadRequest, "Username already exists.")
	json.NewEncoder(w).Encode(signInFailure)

}
