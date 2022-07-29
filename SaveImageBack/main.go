package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"saveimage/handlers"
	"saveimage/storage"
	"saveimage/variable"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	createFolder()
	storage.STORAGE = storage.OpenPostgres()
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.Use(cors.Default())
	router.POST("/cards", handlers.PostCard)
	router.GET("/cards", handlers.GetCards)
	router.DELETE("/cards/:id", handlers.DeleteCard)
	router.GET("/tags", handlers.GetTags)

	router.StaticFS("/images", http.Dir("images"))
	router.Run("localhost:8080")
}

// создаёт папку для хранения картинок
func createFolder() {
	err := os.MkdirAll(variable.CREATE_FOLDER, 0755)
	if err != nil {
		log.Fatal(errors.New("createFolder - Папка не создана"))
	}
	fmt.Println("Папка готова для работы")
}
