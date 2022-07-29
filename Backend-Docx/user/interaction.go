package user

import (
	"fmt"
	"io/ioutil"
	"os"
)

/*возвращает данные пользователя*/

func GetUser(allUsers []User, login, password string) (string, string, []string) {
	var message, err string
	var file []string
	key, permission := PermissionUser(allUsers, login, password)
	if permission {
		userCreateFolder(login) // навсякий случай создаём папку если ранее не создали
		file = allUsers[key].GetAllFile()
		message = "Авторизация прошла успешна"
		return message, err, file
	}
	err = "Неверный логин или пароль"
	return message, err, file
}

/*создаёт нового пользователя*/

func NewUser(allUsers *[]User, login, password string) (string, string) {
	_, permission := PermissionUser(*allUsers, login, password)
	var message, err string
	if !permission {
		user := User{Login: login, Password: password, File: []string{}}
		*allUsers = append(*allUsers, user)
		userCreateFolder(login)
		message = "Новый пользователь создан"
		return message, err
	}
	err = "Такой пользователь уже есть"
	return message, err
}

func UserChangeFile(allUsers []User, login, password, filename, text string) (string, string) {
	var message, err string
	key, permission := PermissionUser(allUsers, login, password)
	if permission {
		message, err = allUsers[key].ChangeFile(filename, text)
	} else {
		err = "Что то не так, не найден пользователь"
	}
	return message, err
}

func UserCreateNewFile(allUsers []User, login, password, filename, text string) (string, string) {
	var message, err string
	key, permission := PermissionUser(allUsers, login, password)
	if permission {
		message, err = allUsers[key].AddNewFile(filename, text)
	} else {
		err = "Что то не так, не найден пользователь"
	}
	return message, err
}

func UserGetFile(allUsers []User, login, password, filename string) (string, string, string) {
	var message, err, text string
	key, permission := PermissionUser(allUsers, login, password)
	if permission {
		message, err, text = allUsers[key].GetFile(filename)
	} else {
		err = "Что то не так, не найден пользователь"
	}
	return message, err, text
}

/*создаёт папку*/

func userCreateFolder(login string) {
	folderCreate := false
	files, _ := ioutil.ReadDir("userFolder")
	for _, file := range files {
		if file.Name() == login {
			folderCreate = true
		}
	}
	if !folderCreate {
		os.Mkdir(("userFolder/" + login), 0755)
		fmt.Printf("Папка %v создана", login)
	}
}

/*проверяет есть ли такой пользователь*/

func PermissionUser(allUsers []User, login, password string) (int, bool) {
	for i, v := range allUsers {
		if v.Login == login && v.Password == password {
			return i, true
		}
	}
	return -1, false
}

func UserDeleteFile(allUsers []User, login, password, filename string) (string, string) {
	var message, err string
	key, permission := PermissionUser(allUsers, login, password)
	if permission {
		message, err = allUsers[key].DeleteFile(filename)
	} else {
		err = "Что то не так, не найден пользователь"
	}
	return message, err
}
