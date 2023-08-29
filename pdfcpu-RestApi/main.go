package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	addstamp "restapi/addStamp"
	"restapi/collect"
	"restapi/crop"
	deletestamp "restapi/deleteStamp"
	"restapi/global"
	"restapi/merge"
	"restapi/minio"
	"restapi/optimize"
	"restapi/rotate"
	"restapi/split"
	"restapi/trim"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type envConfig struct {
	endpoint        string
	accessKeyID     string
	secretAccessKey string
	region          string
	bucketName      string
	port            string
}

/*split - режит файл на страницы*/

func main() {

	/*создаёт временную папку*/
	if err := global.MakeDirectoryIfNotExists("TempFile"); err != nil {
		log.Fatal("Не удалось создать временную папку.")
	}

	/*считывает данные с .env*/
	variable, err := envFile()
	if err != nil {
		log.Fatal(err.Error())
	}

	/*подключаеться к серверу*/
	minio.ConnectionMinion(variable.endpoint, variable.accessKeyID, variable.secretAccessKey, variable.region, variable.bucketName)

	/*маршрутизация*/
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	router.Use(cors.Default())
	router.POST("/addWatermarks", addstamp.AddStamp)
	router.POST("/removeWatermarks", deletestamp.DeleteStamp)
	router.POST("/collect", collect.Collect)
	router.POST("/rotate", rotate.Rotate)
	router.POST("/trim", trim.Trim)
	router.POST("/crop", crop.Crop)
	router.POST("/merge", merge.Merge)
	router.POST("/optimize", optimize.Optimize)
	router.POST("/split", split.Split)

	fmt.Println("Server run on port:" + variable.port)
	router.Run(":" + variable.port)
}

func envFile() (envConfig, error) {

	var varible envConfig
	var exists bool

	if err := godotenv.Load(); err != nil {
		return varible, errors.New("No .env file found")
	}

	if varible.endpoint, exists = os.LookupEnv("ENDPOINT"); !exists {
		return varible, errors.New("ENDPOINT not exists")
	}
	if varible.accessKeyID, exists = os.LookupEnv("ACCESS_KEY_ID"); !exists {
		return varible, errors.New("ACCESS_KEY_ID not exists")
	}
	if varible.secretAccessKey, exists = os.LookupEnv("SECRET_ACCESS_KEY"); !exists {
		return varible, errors.New("SECRET_ACCESS_KEY not exists")
	}
	if varible.region, exists = os.LookupEnv("REGION"); !exists {
		return varible, errors.New("REGION not exists")
	}
	if varible.bucketName, exists = os.LookupEnv("BUCKET_NAME"); !exists {
		return varible, errors.New("BUCKET_NAME not exists")
	}
	if varible.port, exists = os.LookupEnv("PORT"); !exists {
		varible.port = "3000"
		return varible, nil
	}
	return varible, nil
}
