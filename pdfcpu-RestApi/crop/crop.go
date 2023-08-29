package crop

import (
	"mime/multipart"
	"net/http"
	"os"
	"restapi/global"
	"restapi/minio"
	"strings"

	"github.com/gin-gonic/gin"
	pdf "github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func Crop(c *gin.Context) {
	var inFile *multipart.FileHeader // сам файл до изменения
	var selectedPages []string = nil // страницы для изменеия all|тоже-самое что и в документации
	var description string
	var box *pdfcpu.Box
	var err error

	/*получение всего*/
	if inFile, err = c.FormFile("inFile"); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	inFile.Filename = global.GenerateFileName(inFile.Filename, "crop")

	if description = c.PostForm("description"); description == "" {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "description не указан"})
		return
	}

	if c.PostForm("selectedPages") != "all" {
		selectedPages = strings.Split(c.PostForm("selectedPages"), ",")
	}
	/*сохранение всего*/
	if err = c.SaveUploadedFile(inFile, "./TempFile/"+inFile.Filename); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conf := pdfcpu.NewDefaultConfiguration()

	if box, err = pdf.Box(description, conf.Unit); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err = pdf.CropFile("./TempFile/"+inFile.Filename, "", selectedPages, box, conf); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	minio.GlobalMinion.SendFile(inFile.Filename)
	var createUrl = minio.GlobalMinion.GetUrlFile(inFile.Filename)
	c.IndentedJSON(http.StatusOK, gin.H{"file": createUrl})
	defer os.Remove("./TempFile/" + inFile.Filename)

}
