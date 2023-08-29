package split

import (
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"restapi/global"
	"restapi/minio"
	"strconv"

	pdf "github.com/pdfcpu/pdfcpu/pkg/api"

	"github.com/gin-gonic/gin"
)

func Split(c *gin.Context) {
	var inFile *multipart.FileHeader // сам файл до изменения
	var splitPath string = global.GeneratePath("pathsplit", "split")
	var allUrl []string

	var span int
	var err error
	if inFile, err = c.FormFile("inFile"); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	inFile.Filename = global.GenerateFileName(inFile.Filename, "split")

	if span, err = strconv.Atoi(c.PostForm("span")); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	/*сохранение всего*/
	if err = c.SaveUploadedFile(inFile, "./TempFile/"+inFile.Filename); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	splitPath = global.GeneratePath("pathsplit", "split")
	if err := global.MakeDirectoryIfNotExists("./TempFile/" + splitPath); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	/*добавление split*/
	if err = pdf.SplitFile("./TempFile/"+inFile.Filename, "./TempFile/"+splitPath, span, nil); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	files, _ := ioutil.ReadDir("./TempFile/" + splitPath)
	for _, file := range files {
		minio.GlobalMinion.SendFile(splitPath + "/" + file.Name())
		allUrl = append(allUrl, minio.GlobalMinion.GetUrlFile(splitPath+"/"+file.Name()))
	}

	// var createUrl = minio.GlobalMinion.GetUrlFile(splitPath)
	c.IndentedJSON(http.StatusOK, gin.H{"file": allUrl})

	defer os.RemoveAll("./TempFile/" + splitPath)
	defer os.Remove("./TempFile/" + inFile.Filename)

}
