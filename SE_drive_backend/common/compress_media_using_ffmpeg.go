package common

import (
	errors "SE_drive_backend/Errors"
	"SE_drive_backend/models"
	"fmt"
	"net/http"
	"os/exec"
)

func CompressMediaUsingFfmpeg(inputPhotoFilePath string, outputPhotoFilePath string, types string) (models.ErrorsModel, bool) {

	var cmdStr string
	fmt.Print(inputPhotoFilePath)
	fmt.Print(outputPhotoFilePath)
	switch types {
	default:
		cmdStr = fmt.Sprintf("ffmpeg -i %s -qscale:v 31 -f image2 -vcodec mjpeg %s", inputPhotoFilePath, outputPhotoFilePath)

	}
	cmd := exec.Command("cmd", "/C", cmdStr)

	er := cmd.Run()
	if er != nil {
		return errors.SetErrorModel(http.StatusBadRequest, "error while running the ffmpeg."), false
	}
	return models.ErrorsModel{}, true
}
