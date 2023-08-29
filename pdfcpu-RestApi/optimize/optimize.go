package optimize

import (
	"mime/multipart"
	"net/http"
	"os"
	"restapi/global"
	"restapi/minio"

	pdf "github.com/pdfcpu/pdfcpu/pkg/api"

	"github.com/gin-gonic/gin"
)

func Optimize(c *gin.Context) {
	var inFile *multipart.FileHeader // сам файл до изменения
	var err error
	if inFile, err = c.FormFile("inFile"); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	inFile.Filename = global.GenerateFileName(inFile.Filename, "optimize")

	/*сохранение всего*/
	if err = c.SaveUploadedFile(inFile, "./TempFile/"+inFile.Filename); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	/*добавление collect*/
	if err = pdf.OptimizeFile("./TempFile/"+inFile.Filename, "", nil); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	minio.GlobalMinion.SendFile(inFile.Filename)
	var createUrl = minio.GlobalMinion.GetUrlFile(inFile.Filename)
	c.IndentedJSON(http.StatusOK, gin.H{"file": createUrl})
	defer os.Remove("./TempFile/" + inFile.Filename)

}
