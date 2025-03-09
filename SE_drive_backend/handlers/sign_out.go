package handlers

import (
	errors "SE_drive_backend/Errors"
	"SE_drive_backend/functions"
	"SE_drive_backend/global"
	"SE_drive_backend/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func SignOut(w http.ResponseWriter, r *http.Request) {
	//insert all thngs from added maps into db .

	var signOutRequestModel models.SignOutRequestModel
	// var signOutResponseModel models.SignOutResponseModel

	jsonError := json.NewDecoder(r.Body).Decode(&signOutRequestModel)
	if jsonError != nil {
		json.NewEncoder(w).Encode(errors.SetErrorModel(http.StatusBadRequest, "error while decoding the token from request . "))

	}
	db, dbErr := functions.DbConnect(w)
	if dbErr != nil {
		json.NewEncoder(w).Encode(errors.SetErrorModel(http.StatusBadRequest, fmt.Sprintf("error while connecting to db from signout. %s", dbErr)))

		return
	}

	data, exists := global.MediaMap[signOutRequestModel.Token]
	if !exists {
		json.NewEncoder(w).Encode(errors.SetErrorModel(
			http.StatusBadRequest,
			"Token not found in MediaMap."))
		return
	}

	trialsInsert := `insert into trialstable(token,trialsLeft) values(?,?)
	 ON DUPLICATE KEY UPDATE trialsLeft = VALUES(trialsLeft)`
	_, insertTrialErr := db.Exec(trialsInsert, signOutRequestModel.Token, data.TrialsLeft)
	if insertTrialErr != nil {
		json.NewEncoder(w).Encode(errors.SetErrorModel(http.StatusBadRequest, fmt.Sprintf("error while inserting into the trials db , %s", insertTrialErr)))
		return
	}
	delete(global.MediaMap, signOutRequestModel.Token)

	json.NewEncoder(w).Encode(models.SignOutResponseModel{Message: "sign out successful .", StatusCode: 200})

}
