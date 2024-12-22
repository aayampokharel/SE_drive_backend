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

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Fatal(err)
	}

	file, header, er := r.FormFile("file")
	if er != nil {
		log.Fatal(er)
	}
	defer file.Close()

	fmt.Print("the size of file is: ", header.Size, "\n")
	fmt.Print("the header of file is: ", header.Header, "\n")
	fmt.Print("the name of file is: ", header.Filename, "\n")
	directoryToStoreImages, er := os.Create("./uploads/" + "input.png")
	if er != nil {
		log.Fatal(er)
	}
	_, er = io.Copy(directoryToStoreImages, file)
	if er != nil {
		log.Fatal(er)
	}
	//ffmpeg -i input.png output.jpg

	cmd := exec.Command("ffmpeg", "-i", "./uploads/input.png", "./uploads/output.jpg")

	er = cmd.Run()
	if er != nil {
		log.Fatal(er)
	}
	defer fmt.Print("done")
	defer directoryToStoreImages.Close()

}
