package addstamp

import (
	"mime/multipart"
	"net/http"
	"os"
	"restapi/global"
	"restapi/minio"

	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	pdf "github.com/pdfcpu/pdfcpu/pkg/api"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func AddStamp(c *gin.Context) {
	var inFile *multipart.FileHeader // сам файл до изменения
	var watermark *pdfcpu.Watermark  // значения по умалчанию

	var mode string                  // тип штампа img|text|pdf
	var update bool = false          // обновление stamp
	var onTop bool = false           // поверх контента или нет stamp|watermarks
	var selectedPages []string = nil // страницы для изменеия all|тоже-самое что и в документации

	var fileMode *multipart.FileHeader // сам файл pdf|img для stamp в случае если mode указан как img|pdf
	var pagePdfMode string             // какую страницу взять с pdf в случае если mode указан как pdf
	var textMode string = "demo"       // текст для stamp в случае если mode указан как text
	var description string             // настройки для stamp

	var err error // ошибка

	/*получение всего*/
	if inFile, err = c.FormFile("inFile"); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	inFile.Filename = global.GenerateFileName(inFile.Filename, "addstamp")

	mode = c.PostForm("mode")
	if mode == "img" || mode == "pdf" {
		if fileMode, err = c.FormFile("fileMode"); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fileMode.Filename = global.GenerateFileName(fileMode.Filename, "stamp")
		if err = c.SaveUploadedFile(fileMode, "./TempFile/"+fileMode.Filename); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if mode == "pdf" {
			pagePdfMode = c.PostForm("pagePdfMode")
		}
	}

	if onTop, err = strconv.ParseBool(c.PostForm("onTop")); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if update, err = strconv.ParseBool(c.PostForm("update")); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if c.PostForm("selectedPages") != "all" {
		selectedPages = strings.Split(c.PostForm("selectedPages"), ",")
	}
	textMode = c.PostForm("text")
	description = c.PostForm("description")
	// fmt.Println("file", inFile.Filename, "mode", mode, "onTop", onTop, "update", update, "pages", selectedPages, "text", textMode, "desc", description)

	/*сохранение всего*/
	if err = c.SaveUploadedFile(inFile, "./TempFile/"+inFile.Filename); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	/*добавляет конфигурации к stamp*/
	if mode == "text" {
		watermark, err = pdf.TextWatermark(textMode, description, onTop, update, pdfcpu.POINTS)
	} else if mode == "img" {
		watermark, err = pdf.ImageWatermark("./TempFile/"+fileMode.Filename, description, onTop, update, pdfcpu.POINTS)
	} else if mode == "pdf" {
		watermark, err = pdf.PDFWatermark("./TempFile/"+fileMode.Filename+":"+pagePdfMode, description, onTop, update, pdfcpu.POINTS)
	}
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	/*добавление stamp*/
	if err = pdf.AddWatermarksFile("./TempFile/"+inFile.Filename, "", selectedPages, watermark, nil); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	minio.GlobalMinion.SendFile(inFile.Filename)

	var createUrl = minio.GlobalMinion.GetUrlFile(inFile.Filename)

	c.IndentedJSON(http.StatusOK, gin.H{"file": createUrl})

	defer os.Remove("./TempFile/" + inFile.Filename)
	if mode != "text" {
		defer os.Remove("./TempFile/" + fileMode.Filename)
	}
}
