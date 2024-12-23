package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	//"path/filepath"
)

func uploadAudio(w http.ResponseWriter, r *http.Request) {
	CORSFix(w, r)
	const directory = "./uploadedAudio/"
	var fileName string //@ later stores the filename after arrival.
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Fatal(err)
	}
	file, header, er := r.FormFile("Audio")
	if er != nil {
		log.Fatal(er)
	}
	defer file.Close() //@ ??
	fileName = header.Filename

	createdFile, er := os.Create(directory + "audio_" + fileName)
	if er != nil {
		log.Fatal(er)
	}
	defer createdFile.Close()
	io.Copy(createdFile, file)
	cmd := exec.Command("ffmpeg", "-i", createdFile.Name(), directory+"audio"+fileName+".mp3")
	er = cmd.Run()
	if er != nil {
		log.Fatal(er)
	}

}
