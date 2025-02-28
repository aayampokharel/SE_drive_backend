package handlers

import (
	errors "SE_drive_backend/Errors"
	"SE_drive_backend/models"
	"encoding/json"
	"net/http"
)

func GetHistory(w http.ResponseWriter, r *http.Request) {
	//know if response can be multipart
	//multipart writer .
	//write .

	CORSFix(w, r)

	var failure models.ErrorsModel
	if mediaType := r.Header.Get("Accept"); mediaType == "multipart/form-data" {
		failure = errors.SetErrorModel(http.StatusBadRequest, "Header not able to accept multipart/form-data mediaTypes.")
		json.NewEncoder(w).Encode(failure)
		return

	}

}

// 	///	mw := multipart.NewWriter(w)
// 	//! join operation use garera get , subscription if true get original file names list , else get output_file names list .
// 	//! typ list ma bhako map with the files in the folder and write in the multipart .
// 	// 	getFileNamesQuery := `SELECT

// 	//     u.userName,
// 	//     u.isSubscribed,

// 	//     CASE
// 	//         WHEN u.isSubscribed = FALSE THEN v.outputVideoFileName
// 	//         ELSE v.originalVideoFileName
// 	//     END AS VideoFileName,

// 	//     CASE
// 	//         WHEN u.isSubscribed = FALSE THEN p.outputPhotoFileName
// 	//         ELSE p.originalPhotoFileName
// 	//     END AS PhotoFileName,

// 	//     CASE
// 	//         WHEN u.isSubscribed = FALSE THEN pdf.outputPdfFileName
// 	//         ELSE pdf.originalPdfFileName
// 	//     END AS PdfFileName,

// 	//     CASE
// 	//         WHEN u.isSubscribed = FALSE THEN a.outputAudioFileName
// 	//         ELSE a.originalAudioFileName
// 	//     END AS AudioFileName,

// 	//     CASE
// 	//         WHEN u.isSubscribed = FALSE THEN t.outputTextFileName
// 	//         ELSE t.originalTextFileName
// 	//     END AS TextFileName

// 	// FROM UserInfoTable u
// 	// LEFT JOIN VideoTable v ON u.token = v.token
// 	// LEFT JOIN PhotoTable p ON u.token = p.token
// 	// LEFT JOIN PdfTable pdf ON u.token = pdf.token
// 	// LEFT JOIN AudioTable a ON u.token = a.token
// 	// LEFT JOIN TextTable t ON u.token = t.token
// 	// WHERE  u.email = (?);
// 	// `

// }
