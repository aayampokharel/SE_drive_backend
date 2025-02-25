package main

import (
	"SE_drive_backend/handlers"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func uploadPdf(w http.ResponseWriter, r *http.Request) {
	handlers.CORSFix(w, r)
	er := r.ParseMultipartForm(10 << 20)
	if er != nil {
		fmt.Print("eror while parsing pdf")

		log.Fatal(er)

	}
	file, header, er := r.FormFile("PDF")
	if er != nil {
		fmt.Print("error while parsing form field ")
		log.Fatal(er)
	}
	fileName := replaceSpaceInFileName(header.Filename)
	newPdfFile, er := os.Create("./uploadedPdfs/" + fileName)
	if er != nil {
		fmt.Print("error while creating file")
		log.Fatal(er)
	}
	_, err := io.Copy(newPdfFile, file)
	if err != nil {
		fmt.Print("error while copying pdf")
		log.Fatal(err)
	}

}
