// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net/url"
// 	"time"

// 	"github.com/minio/minio-go/v7"
// 	"github.com/minio/minio-go/v7/pkg/credentials"
// 	"github.com/minio/minio-go/v7/pkg/lifecycle"
// )

// func main() {

// 	// Подключение к minio

// 	ctx := context.Background()
// 	endpoint := "127.0.0.1:9000"
// 	accessKeyID := "admin"
// 	secretAccessKey := "password"
// 	useSSL := false

// 	minioClient, err := minio.New(endpoint, &minio.Options{
// 		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
// 		Secure: useSSL,
// 	})

// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	// Создание корзины

// 	bucketName := "myfile"
// 	location := "us-east-1"

// 	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
// 	if err != nil {
// 		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
// 		if errBucketExists == nil && exists {
// 			log.Printf("We already own %s\n", bucketName)
// 		} else {
// 			log.Fatalln(err)
// 		}
// 	} else {
// 		log.Printf("Successfully created %s\n", bucketName)
// 	}

// 	// Загрузка файлов в корзину
// 	objectName := "dtest.pdf"
// 	filePath := "./file/test.pdf"
// 	contentType := "application/pdf"

// 	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)

// 	/*установка времени жизни файлов*/
// 	config := lifecycle.NewConfiguration()
// 	config.Rules = []lifecycle.Rule{
// 		{
// 			ID:     "expire-bucket",
// 			Status: "Enabled",
// 			Expiration: lifecycle.Expiration{
// 				Days: 2,
// 			},
// 		},
// 	}

// 	err = minioClient.SetBucketLifecycle(context.Background(), "myfile", config)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	// Set request parameters for content-disposition.
// 	reqParams := make(url.Values)
// 	reqParams.Set("response-content-disposition", "attachment; filename=\"dtest.pdf\"")

// 	// Generates a presigned url which expires in a day.
// 	presignedURL, err := minioClient.PresignedGetObject(context.Background(), "myfile", "dtest.pdf", time.Second*24*60*60, reqParams)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("Successfully generated presigned URL", presignedURL)
// }

// /* Получить ссылку на файл */
// /* Удалять файл через 3 месяца */

package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7/pkg/lifecycle"
)

func main() {
	ConnectionMinion("127.0.0.1:9000", "admin", "password", false)
	GlobalMinion.CreateBucket("other", "us-east-1", "application/pdf", context.Background())
	GlobalMinion.BucketLifecycle(90, "testId")
	GlobalMinion.SendFile("my.pdf", "file/my.pdf")
	fmt.Println(GlobalMinion.GetUrlFile("my.pdf"))

}

var GlobalMinion MinioMain

/*подключение к minio*/
func ConnectionMinion(
	endpoint,
	accessKeyID,
	secretAccessKey string, useSSL bool) {
	var err error

	GlobalMinion.endpoint = endpoint
	GlobalMinion.accessKeyID = accessKeyID
	GlobalMinion.secretAccessKey = secretAccessKey
	GlobalMinion.useSSL = useSSL

	GlobalMinion.minioClient, err = minio.New(GlobalMinion.endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(GlobalMinion.accessKeyID, GlobalMinion.secretAccessKey, ""),
		Secure: GlobalMinion.useSSL,
	})

	if err != nil {
		log.Fatalln(err)
		fmt.Println("Не удалось подключиться к minio. По причине", err.Error())
	}
}

type MinioMain struct {
	ctx             context.Context
	endpoint        string
	accessKeyID     string
	secretAccessKey string
	useSSL          bool
	minioClient     *minio.Client
	bucketName      string
	location        string
	contentType     string
}

/*создание корзины*/
func (obj *MinioMain) CreateBucket(bucketName, location, contentType string, ctx context.Context) {
	obj.bucketName = bucketName
	obj.location = location
	obj.contentType = contentType
	obj.ctx = ctx

	err := obj.minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := obj.minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			fmt.Println("Такая корзина уже есть", bucketName)
			return
		}
		fmt.Println("Не удалось создать/подключиться к", bucketName, "По причине", err.Error())
		return
	}
	fmt.Println("Корзина создана", bucketName)
}

/*установка времени жизни файлов*/
func (obj *MinioMain) BucketLifecycle(lifecycleDay int, idLifecycle string) {
	config := lifecycle.NewConfiguration()
	config.Rules = []lifecycle.Rule{
		{
			ID:     idLifecycle,
			Status: "Enabled",
			Expiration: lifecycle.Expiration{
				Days: lifecycle.ExpirationDays(lifecycleDay),
			},
		},
	}

	err := obj.minioClient.SetBucketLifecycle(obj.ctx, obj.bucketName, config)
	if err != nil {
		fmt.Println("Жизненый цикл не получилось создать. Причина", err.Error())
		return
	}
	fmt.Println("Жизненый цикл успешно создан в размере", lifecycleDay)

}

/*Загрузка файла в корзину*/
func (obj *MinioMain) SendFile(fileName, filePath string) {
	info, err := obj.minioClient.FPutObject(obj.ctx, obj.bucketName, fileName, filePath, minio.PutObjectOptions{ContentType: obj.contentType})
	if err != nil {
		fmt.Println("Не удалось загрузить файл", fileName)
	}
	fmt.Println("Успешно загрузилось", fileName, info.Size)
}

/*Создание ссылки на файл*/
func (obj *MinioMain) GetUrlFile(fileName string) *url.URL {
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\""+fileName+"\"")
	presignedURL, err := obj.minioClient.PresignedGetObject(context.Background(), obj.bucketName, fileName, time.Second*24*60*60*7, reqParams)
	if err != nil {
		fmt.Println("Не удалось создать ссылку причина в", err.Error())
		return nil
	}
	fmt.Println("Ссылка на файл успешно создана", presignedURL)
	return presignedURL
}
