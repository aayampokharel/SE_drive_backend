package main

import (
	"SE_drive_backend/handlers"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func uploadVideo(w http.ResponseWriter, r *http.Request) {
	var fileName string
	var videoDirectory string = "./uploadedVideos/"
	handlers.CORSFix(w, r)

	reader, err := r.MultipartReader()

	if err != nil {
		fmt.Print("error while parting")

		log.Fatal(err)
	}
	for { //part is a stream of data .
		part, er := reader.NextPart()
		if er == io.EOF {
			fmt.Print("==end of file==")
			break

		}

		if part.FormName() == "Video" {
			fileName = "video_" + replaceSpaceInFileName(part.FileName())

			newVideoFile, err := os.Create(videoDirectory + fileName)

			if err != nil {
				fmt.Print("1st error ")
				log.Fatal(er)
			}
			//create a small buffer , then supply this to others .

			_, err = io.Copy(newVideoFile, part) // Stream file data to disk
			if err != nil {
				fmt.Print("2nd error ")
				log.Println("Error saving file:", err)
			}

		}
	}

	//@ eg : vid1.mp4 then ["vid1","mp4"];

	nameWithoutExtension := strings.Split(fileName, ".")[0] //@ext=vid1;
	outputFileName := nameWithoutExtension + ".mp4"

	videoOutputFile := videoDirectory + "output_" + outputFileName
	cmd := exec.Command("ffmpeg", "-i", videoDirectory+fileName, "-vf", "scale=1280:-1", "-c:v", "libx264", "-preset", "slow", "-crf", "23",
		"-c:a", "aac", "-b:a", "128k", videoOutputFile)

	// cmd := exec.Command("ffmpeg", "-i", "./uploadedVideos/gold.mp4", "-vf", "scale=1280:-1", "-c:v", "libx264", "-preset", "slow", "-crf", "23", "-c:a", "aac", "-b:a", "128k", "./uploadedVideos/output2.mp4")
	e := cmd.Run()
	if e != nil {
		fmt.Print("3rd error ")
		log.Fatal(e)
	}
	defer uploadFromResponse(w, videoOutputFile, "video", 1024*500)

}
