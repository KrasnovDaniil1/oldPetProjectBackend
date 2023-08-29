package deletestamp

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"restapi/global"
	"restapi/minio"
	"strings"

	"github.com/gin-gonic/gin"
	pdf "github.com/pdfcpu/pdfcpu/pkg/api"
)

func DeleteStamp(c *gin.Context) {
	var inFile *multipart.FileHeader
	var selectedPages []string = nil // страницы для изменеия all|тоже-самое что и в документации

	var err error

	if inFile, err = c.FormFile("inFile"); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	inFile.Filename = global.GenerateFileName(inFile.Filename, "removestamp")

	if err = c.SaveUploadedFile(inFile, "./TempFile/"+inFile.Filename); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if c.PostForm("selectedPages") != "all" {
		selectedPages = strings.Split(c.PostForm("selectedPages"), ",")
	}

	if err = pdf.RemoveWatermarksFile("./TempFile/"+inFile.Filename, "", selectedPages, nil); err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	minio.GlobalMinion.SendFile(inFile.Filename)

	var createUrl = minio.GlobalMinion.GetUrlFile(inFile.Filename)

	c.IndentedJSON(http.StatusOK, gin.H{"file": createUrl})

	defer os.Remove("./TempFile/" + inFile.Filename)
}
