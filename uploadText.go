package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func uploadText(w http.ResponseWriter, r *http.Request) {
	CORSFix(w, r)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		fmt.Print("error while parsing in uploadText")
		log.Fatal(err)
	}

	file, header, err := r.FormFile("Text")
	defer file.Close()
	if err != nil {
		fmt.Print("error while assigning Text file")
		log.Fatal(err)
	}

	newTextFile, er := os.Create("./uploadedTexts/" + header.Filename)
	if er != nil {
		log.Fatal(er)
	}
	defer newTextFile.Close()
	_, er = io.Copy(newTextFile, file)
	if er != nil {
		fmt.Print("error copying")
		log.Fatal(er)
	}

}