package collect

import (
	"mime/multipart"
	"net/http"
	"os"
	"restapi/global"
	"restapi/minio"
	"strings"

	pdf "github.com/pdfcpu/pdfcpu/pkg/api"

	"github.com/gin-gonic/gin"
)

func Collect(c *gin.Context) {
	var inFile *multipart.FileHeader // сам файл до изменения
	var selectedPages []string = nil // страницы для изменеия all|тоже-самое что и в документации
	var err error
	/*получение всего*/
	if inFile, err = c.FormFile("inFile"); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить файл для изменения."})
		return
	}
	inFile.Filename = global.GenerateFileName(inFile.Filename, "collect")

	selectedPages = strings.Split(c.PostForm("selectedPages"), ",")

	/*сохранение всего*/
	if err = c.SaveUploadedFile(inFile, "./TempFile/"+inFile.Filename); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить файл, возможно с ним что-то не так."})
		return
	}

	/*добавление collect*/
	if err = pdf.CollectFile("./TempFile/"+inFile.Filename, "", selectedPages, nil); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Не удалось изменить файл."})
		return
	}
	minio.GlobalMinion.SendFile(inFile.Filename)
	var createUrl = minio.GlobalMinion.GetUrlFile(inFile.Filename)
	c.IndentedJSON(http.StatusOK, gin.H{"file": createUrl})
	defer os.Remove("./TempFile/" + inFile.Filename)

}
