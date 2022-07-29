package handler

import (
	"app/user"
	"encoding/json"
	"io"
	"net/http"
)

type UserResp struct {
	Login    string   `json:"login"`
	Password string   `json:"password"`
	Filename string   `json:"filename"`
	Filetext string   `json:"filetext"`
	File     []string `json:"file"`
	Status   string   `json:"status"`
	Error    string   `json:"error"`
}

func UserGetPost(w http.ResponseWriter, r *http.Request) {
	var userBody UserResp

	bodyBytes, _ := io.ReadAll(r.Body)

	json.Unmarshal([]byte(bodyBytes), &userBody)

	if r.Method == http.MethodGet {
		userBody.Status, userBody.Error, userBody.File = user.GetUser(user.AllUsers, userBody.Login, userBody.Password)
	} else if r.Method == http.MethodPost {
		userBody.Status, userBody.Error = user.NewUser(&user.AllUsers, userBody.Login, userBody.Password)
	} else {
		userBody.Error = "Невозможно обработать запрос, ошибка в сервере"
	}
	userWrite, _ := json.Marshal(userBody)
	w.Header().Set("Content-Type", "application/json")
	w.Write(userWrite)

}

func UserFilenameGetPostPutDelete(w http.ResponseWriter, r *http.Request) {

	var userBody UserResp

	bodyBytes, _ := io.ReadAll(r.Body)

	json.Unmarshal([]byte(bodyBytes), &userBody)

	if r.Method == http.MethodGet {
		userBody.Status, userBody.Error, userBody.Filetext = user.UserGetFile(user.AllUsers, userBody.Login, userBody.Password, userBody.Filename)
	} else if r.Method == http.MethodPost {
		userBody.Status, userBody.Error = user.UserCreateNewFile(user.AllUsers, userBody.Login, userBody.Password, userBody.Filename, userBody.Filetext)
	} else if r.Method == http.MethodPut {
		userBody.Status, userBody.Error = user.UserChangeFile(user.AllUsers, userBody.Login, userBody.Password, userBody.Filename, userBody.Filetext)
	} else if r.Method == http.MethodDelete {
		userBody.Status, userBody.Error = user.UserDeleteFile(user.AllUsers, userBody.Login, userBody.Password, userBody.Filename)
	} else {
		userBody.Error = "Невозможно обработать запрос, ошибка в сервере"
	}

	userWrite, _ := json.Marshal(userBody)
	w.Header().Set("Content-Type", "application/json")
	w.Write(userWrite)
}
