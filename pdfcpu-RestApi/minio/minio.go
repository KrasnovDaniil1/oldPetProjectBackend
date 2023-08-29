package minio

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var GlobalMinion MinioMain

type MinioMain struct {
	ctx             context.Context
	endpoint        string
	accessKeyID     string
	secretAccessKey string
	useSSL          bool
	minioClient     *minio.Client
	bucketName      string
	region          string
	contentType     string
}

/*подключение к minio*/
func ConnectionMinion(
	endpoint,
	accessKeyID,
	secretAccessKey, region, bucketName string) {
	var err error

	GlobalMinion.endpoint = endpoint
	GlobalMinion.accessKeyID = accessKeyID
	GlobalMinion.secretAccessKey = secretAccessKey
	GlobalMinion.region = region
	GlobalMinion.bucketName = bucketName
	GlobalMinion.useSSL = false
	GlobalMinion.ctx = context.Background()

	GlobalMinion.minioClient, err = minio.New(GlobalMinion.endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(GlobalMinion.accessKeyID, GlobalMinion.secretAccessKey, ""),
		Secure: GlobalMinion.useSSL,
	})

	if err != nil {
		log.Fatalln(err)
		fmt.Println("Не удалось подключиться к серверу. По причине", err.Error())
	}
}

/*Загрузка файла в корзину*/
func (obj *MinioMain) SendFile(fileName string) {
	info, err := obj.minioClient.FPutObject(obj.ctx, obj.bucketName, fileName, "./TempFile/"+fileName, minio.PutObjectOptions{ContentType: obj.contentType})
	if err != nil {
		fmt.Println("Не удалось загрузить файл", fileName)
	}
	fmt.Println("Успешно загрузилось", fileName, info.Size)
}

/*Создание ссылки на файл*/
func (obj *MinioMain) GetUrlFile(fileName string) string {
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\""+fileName+"\"")
	presignedURL, err := obj.minioClient.PresignedGetObject(context.Background(), obj.bucketName, fileName, time.Second*24*60*60*7, reqParams)
	if err != nil {
		fmt.Println("Не удалось создать ссылку по причине", err.Error())
		return ""
	}
	return presignedURL.String()
}

// func (obj *MinioMain) FolderSend() {
// 	n, err := obj.minioClient.PutObject(context.Background(), obj.bucketName, "my-objectname/sdfs", bytes.NewReader([]byte("Hello, World")), bytes.NewReader([]byte("Hello, World")).Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
// 	if err != nil {
// 		log.Fatalln(err)
// 	} else {
// 		fmt.Println(n)
// 	}
// }
