package handlers

import (
	errors "SE_drive_backend/Errors"
	"SE_drive_backend/functions"
	"SE_drive_backend/global"
	"SE_drive_backend/models"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Login(w http.ResponseWriter, r *http.Request) {

	fmt.Print("cors error from login ")
	var logInDetails models.LogInRequestModel

	var logInFailure models.ErrorsModel
	var token string

	err := json.NewDecoder(r.Body).Decode(&logInDetails)
	if err != nil {
		logInFailure = errors.SetErrorModel(http.StatusBadRequest, "Invalid JSON format. Invalid LogIn") //error models sets the error model , nothing else .
		json.NewEncoder(w).Encode(logInFailure)

		return
	}
	//! password not checked ,only email verified in db .

	//check if user is already in globalMap to prevent unnecessary call of db .

	mapModelValue, ok := functions.DoesUserExistInMap(logInDetails.Email)
	if ok {
		fmt.Print("\n\n\n\n from inside map =====\n\n\n\n")
		json.NewEncoder(w).Encode(mapModelValue)
		return

	}

	db, dbErr := functions.DbConnect(w)
	if dbErr != nil {
		logInFailure = errors.SetErrorModel(http.StatusBadRequest, "error while connecting to DB during Login.")
		json.NewEncoder(w).Encode(logInFailure)
		return
	}
	defer db.Close()
	getFileNamesQuery :=
		`SELECT  
    u.userName, 
    u.isSubscribed,
    u.token,  
    CASE 
        WHEN u.isSubscribed = FALSE THEN v.outputVideoFileName
        ELSE v.originalVideoFileName
    END AS VideoFileName,

    CASE 
        WHEN u.isSubscribed = FALSE THEN p.outputPhotoFileName
        ELSE p.originalPhotoFileName
    END AS PhotoFileName,

    CASE 
        WHEN u.isSubscribed = FALSE THEN pdf.outputPdfFileName
        ELSE pdf.originalPdfFileName
    END AS PdfFileName,

    CASE 
        WHEN u.isSubscribed = FALSE THEN a.outputAudioFileName
        ELSE a.originalAudioFileName
    END AS AudioFileName,

    CASE 
        WHEN u.isSubscribed = FALSE THEN t.outputTextFileName
        ELSE t.originalTextFileName
    END AS TextFileName

FROM UserInfoTable u
LEFT JOIN VideoTable v ON u.token = v.token
LEFT JOIN PhotoTable p ON u.token = p.token
LEFT JOIN PdfTable pdf ON u.token = pdf.token
LEFT JOIN AudioTable a ON u.token = a.token
LEFT JOIN TextTable t ON u.token = t.token
WHERE u.email = (?); 
`
	resultRow, err := db.Query(getFileNamesQuery, logInDetails.Email)
	if err != nil {
		logInFailure = errors.SetErrorModel(http.StatusBadRequest, fmt.Sprintf("error while fetching results from db from gethistory.%s", err))
		json.NewEncoder(w).Encode(logInFailure)
		return
	}
	defer resultRow.Close()
	var mediaMapStructureInitialize = models.MediaMap{ //global structure .for inserting in globalMap , thisis the structure.

		VideosList: []string{},
		PhotosList: []string{},
		AudiosList: []string{},
		TextsList:  []string{},
		PdfsList:   []string{},
	}
	if !resultRow.Next() {
		logInFailure = errors.SetErrorModel(http.StatusBadRequest, "No Such Email Registered.New User?Sign Up If You Are Using New Account.")
		json.NewEncoder(w).Encode(logInFailure)
		return
	}

	for {
		var loginDbModel models.LogInDbModel
		if err := resultRow.Scan(
			&loginDbModel.UserName,
			&loginDbModel.IsSubscribed,
			&loginDbModel.Token,
			&loginDbModel.VideoFileName,
			&loginDbModel.PhotoFileName,
			&loginDbModel.PdfFileName,
			&loginDbModel.AudioFileName,
			&loginDbModel.TextFileName,
		); err != nil {
			logInFailure = errors.SetErrorModel(http.StatusBadRequest, fmt.Sprintf("Error scanning row: %s", err))
			json.NewEncoder(w).Encode(logInFailure)
			return
		}
		//@ make funcition to do below 5 + 5 + 5 lines , take logindbmodel as arguement and checknull there , then append then do other things . just improve readability .can use pointersas well when dealing with global things.
		token = functions.CheckDbNullString(&loginDbModel.Token)
		mediaMapStructureInitialize = *functions.CheckDbNullStringAndReturnMap(loginDbModel, &mediaMapStructureInitialize)

		if !resultRow.Next() {
			break
		}
	}
	//fornext end

	//@as name implies .

	mediaMapStructureInitialize = functions.RemoveDuplicatesFromMapModel(mediaMapStructureInitialize)
	mediaMapStructureInitialize.Email = logInDetails.Email
	mediaMapStructureInitialize.Token = token
	//below assignment is right , as if user exists , then we return above step only , used DoesUserExistInMap() for checking .
	global.MediaMap[token] = mediaMapStructureInitialize
	json.NewEncoder(w).Encode(global.MediaMap[token])

}
