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

	var logInDetails models.LogInRequestModel
	var logInResponseDetails models.LogInResponseModel
	var logInFailure models.ErrorsModel
	var token string
	//# decode process ...........
	err := json.NewDecoder(r.Body).Decode(&logInDetails)
	if err != nil {
		logInFailure = errors.SetErrorModel(http.StatusBadRequest, "Invalid JSON format. Invalid LogIn") //error models sets the error model , nothing else .
		json.NewEncoder(w).Encode(logInFailure)

		return
	}
	//! password not checked ,only email verified in db .

	//check if user is already in globalMap to prevent unnecessary call of db .
	//# check if user exists in map , that is the connection is alive in some somewhere else as well at the same time ..............
	mapModelValue, ok := functions.DoesUserExistInMap(logInDetails.Email)
	if ok {
		//! error in login to be fixed .
		//  redundant or duplicates .
		// global.AddAllToMediaMap(mapModelValue.Token)
		fmt.Print("yes from insdie okay ")
		json.NewEncoder(w).Encode(models.LogInResponseModel{
			IsSubscribed: global.AddedMediaMap[mapModelValue.Token].IsSubscribed,
			TrialsLeft:   global.AddedMediaMap[mapModelValue.Token].TrialsLeft,

			MediaList: &mapModelValue,
		})
		return

	}
	//# if not in map ,i.e. logging in for first time .................
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

	var loginDbModel models.LogInDbModel
	for {
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

	var trialsLeft int
	if !loginDbModel.IsSubscribed {
		queryToFetchTrialsNumber := `Select trialsLeft from trialstable where token=(?)`
		er := db.QueryRow(queryToFetchTrialsNumber, loginDbModel.Token).Scan(&trialsLeft)
		if er != nil {

			logInFailure = errors.SetErrorModel(http.StatusBadRequest, fmt.Sprintf("error while returning trialsleft in login . contact server man . %s", er))
			json.NewEncoder(w).Encode(logInFailure)
			return
		}
		global.LogInInit(token, trialsLeft, loginDbModel.IsSubscribed)

	} else {
		global.LogInInit(token, -1, loginDbModel.IsSubscribed)

	}

	//@as name implies .

	mediaMapStructureInitialize = functions.RemoveDuplicatesFromMapModel(mediaMapStructureInitialize)
	mediaMapStructureInitialize.Email = logInDetails.Email
	mediaMapStructureInitialize.Token = token
	//below assignment is right , as if user exists , then we return above step only , used DoesUserExistInMap() for checking .
	global.MediaMap[token] = &mediaMapStructureInitialize
	logInResponseDetails = models.LogInResponseModel{
		IsSubscribed: loginDbModel.IsSubscribed,
		TrialsLeft:   trialsLeft,

		MediaList: global.MediaMap[token],
	}
	json.NewEncoder(w).Encode(logInResponseDetails)

}
