package merge

import (
	"mime/multipart"
	"net/http"
	"os"
	"restapi/global"
	"restapi/minio"

	"github.com/gin-gonic/gin"
	pdf "github.com/pdfcpu/pdfcpu/pkg/api"
)

func Merge(c *gin.Context) {
	var inFiles []*multipart.FileHeader // сам файл до изменения
	var form *multipart.Form
	var fileNames []string
	var resultFileName string = global.GenerateFileName("resultFileName.pdf", "merge")
	var err error

	if form, err = c.MultipartForm(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	inFiles = form.File["inFiles"]
	for i, file := range inFiles {
		fileNames = append(fileNames, "./TempFile/"+global.GenerateFileName(file.Filename, "merge"))
		if err = c.SaveUploadedFile(file, fileNames[i]); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	resultFileName = global.GenerateFileName("resultFileName.pdf", "merge")
	if err = pdf.MergeCreateFile(fileNames, "./TempFile/"+resultFileName, nil); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	minio.GlobalMinion.SendFile(resultFileName)
	var createUrl = minio.GlobalMinion.GetUrlFile(resultFileName)

	c.IndentedJSON(http.StatusOK, gin.H{"file": createUrl})

	defer os.Remove("./TempFile/" + resultFileName)
	for _, pathFile := range fileNames {
		defer os.Remove(pathFile)
	}

}
