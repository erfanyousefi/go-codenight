package utils

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveUploadFile(ctx *gin.Context, fieldName string, destination string) (string, error) {
	file, err := ctx.FormFile(fieldName)
	if err == nil {
		var filename string
		ext := path.Ext(file.Filename)
		root, _ := os.Getwd()
		filename = fmt.Sprintf("%d%s", time.Now().UnixMilli(), ext)
		address := fmt.Sprintf("/uploads/%s/%s", destination, filename)
		if err := ctx.SaveUploadedFile(file, path.Join(root, "uploads", destination, filename)); err != nil {
			return "", err
		}
		return address, nil
	}
	return "", err
}
