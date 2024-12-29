package main

import (
	"fmt"
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
	fileName = "audio_" + replaceSpaceInFileName(header.Filename)

	createdFile, er := os.Create(directory + fileName)
	if er != nil {
		log.Fatal(er)
	}
	defer createdFile.Close()
	io.Copy(createdFile, file)
	outputAudioFile := directory + "output_" + fileName
	cmdStr := fmt.Sprintf("ffmpeg -i %s -b:a 48k -ar 22050 -ac 1 %s", createdFile.Name(), outputAudioFile)
	cmd := exec.Command("cmd", "/C", cmdStr)
	er = cmd.Run()
	if er != nil {
		log.Fatal(er)
	}

}
