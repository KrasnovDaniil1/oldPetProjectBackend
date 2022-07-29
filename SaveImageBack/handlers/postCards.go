package handlers

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"saveimage/storage"
	"saveimage/variable"
	"strconv"
	"strings"
	"github.com/gin-gonic/gin"
)

var card struct {
	Tags      []string
	ImageName string
}
var image *multipart.FileHeader
var err error

func PostCard(c *gin.Context) {
	card.Tags = strings.Split(c.PostForm("tags"), ",")
	image, err = c.FormFile("image")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить данные."})
		return
	}

	card.ImageName, err = extensionImageFormat(image)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	err = saveImage(c, image)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Внутренняя ошибка сервера"})
		return
	}

	err = storage.STORAGE.UpdateTable(card.Tags, card.ImageName)
	fmt.Println(card)
	if err != nil {
		deleteImage()
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить картинку"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Картинка и теги успеешны сохранены."})
}

func saveImage(c *gin.Context, image *multipart.FileHeader) error {
	err := c.SaveUploadedFile(image, variable.CREATE_FOLDER+"/"+card.ImageName)
	return err
}

func deleteImage() {
	e := os.Remove(card.ImageName)
	if e != nil {
		log.Fatal("Произошла ошибка, картинка была сохранена, но в таблицу не записалась, а удалить не получилось.")
	}
}

func extensionImageFormat(image *multipart.FileHeader) (string, error) {
	var permission = false
	var extension string

	for _, v := range image.Filename {
		if v == '.' {
			permission = true
		}
		if permission {
			extension += string(v)
		}
	}

	lastId, err := storage.STORAGE.GetLastId()
	intLastId, _ := strconv.Atoi(lastId)
	if err != nil {
		return "", errors.New("Внутренняя ошибка сервера")
	}

	for _, v := range variable.VALID_IMAGE_FORMAT {

		if extension == v {
			return strconv.Itoa(intLastId+1) + extension, nil
		}
	}
	fmt.Println(extension)
	return "", errors.New("Неверный формат картинки")
}
