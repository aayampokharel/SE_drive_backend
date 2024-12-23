package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func uploadVideo(w http.ResponseWriter, r *http.Request) {
	CORSFix(w, r)
	fmt.Print("video upload occuring===")
	reader, err := r.MultipartReader()

	if err != nil {
		fmt.Print("error while parting")
		log.Fatal(err)
	}
	for { //part is a stream of data .
		part, er := reader.NextPart()
		if er == io.EOF {
			fmt.Print("==end of file==")
			log.Fatal(er)
		}
		if part.FormName() == "video" {
			//create a small buffer , then supply this to others .
			file, er := os.Create("./uploadedVideos/output.mp4")
			if er != nil {
				log.Fatal(er)
			}
			defer file.Close()
			_, err = io.Copy(file, part) // Stream file data to disk
			if err != nil {
				log.Println("Error saving file:", err)
			}

		}
	}

}
