package main

import (
	"app/handler"
	"fmt"
	"net/http"
)

func main() {

	http.Handle("/user", http.HandlerFunc(handler.UserGetPost))
	http.Handle("/user/filename", http.HandlerFunc(handler.UserFilenameGetPostPutDelete))

	fmt.Println("Сервер запущен")
	http.ListenAndServe(":8081", nil)

}
