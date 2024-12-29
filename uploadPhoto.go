package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func uploadPhoto(w http.ResponseWriter, r *http.Request) {
	CORSFix(w, r)
	err := r.ParseMultipartForm(20 << 20)
	if err != nil {

		log.Fatal(err)
	}

	file, header, er := r.FormFile("Photo")
	if er != nil {

		log.Fatal(er)
	}
	defer file.Close()

	fileName := "photo_" + replaceSpaceInFileName(header.Filename)
	inputPhotoFileStr := "./uploadedPhotos/" + fileName
	newPhotoFile, er := os.Create(inputPhotoFileStr)
	if er != nil {

		log.Fatal(er)
	}
	defer newPhotoFile.Close()
	_, er = io.Copy(newPhotoFile, file)
	if er != nil {
		fmt.Print("4")
		log.Fatal(er)
	}
	outputPhotoFileStr := "./uploadedPhotos/" + "output_" + fileName
	//ffmpeg -i input.png output.jpg
	//@ extensions checking left ....! undone !
	cmdStr := fmt.Sprintf("ffmpeg -i %s -qscale:v 31 %s", newPhotoFile.Name(), outputPhotoFileStr)
	cmd := exec.Command("cmd", "/C", cmdStr)

	er = cmd.Run()
	if er != nil {
		fmt.Print("\n💦💦💦💦💦\n")
		fmt.Print(cmd)
		fmt.Print("\n💦💦💦💦💦\n")
		log.Fatal(er)
	}
	defer fmt.Print("done")

}
