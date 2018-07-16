package service

import (
	"crypto/md5"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/0pen-source/Carpooling/dao"
	"github.com/0pen-source/Carpooling/models"
	"github.com/gin-gonic/gin"
)

type Header struct {
	multipart.FileHeader
}

func Upload(c *gin.Context) {
	payload := models.UserMessage{}
	if err := c.Bind(&payload); err != nil {
		return
	}
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	filename := fmt.Sprintf("%x%s", md5.Sum([]byte(fmt.Sprintf("%s_%s", payload.Phone, payload.ImageType))), path.Ext(header.Filename))
	out, err := os.Create(filename)
	defer out.Close()
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	_, err = io.Copy(out, file)

	go dao.SaveImage(path, filename, payload.ImageType,payload.Phone)

	c.String(http.StatusOK, "upload successful")
}

func GetFileNewName() {

}
