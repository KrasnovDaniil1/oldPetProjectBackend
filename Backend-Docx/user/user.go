package user

import (
	"fmt"
	"io/ioutil"
	"os"
)

var AllUsers = []User{} // все пользователи
type User struct {
	Login    string   `json:"login"`
	Password string   `json:"password"`
	File     []string `json:"filename"`
}

/*создаст новый файл*/

func (user *User) AddNewFile(filename, filetext string) (string, string) {
	var message, err string
	for _, v := range user.File {
		if v == filename {
			err = "Файл с таким именем уже есть"
			return message, err
		}
	}
	user.File = append(user.File, filename)

	file, e := os.Create("userFolder/" + user.Login + "/" + filename + ".txt")

	if e != nil {
		fmt.Println("Unable to create file:", err)
		err = "Файл не был создан причина в железе"
	} else {
		file.WriteString(filetext)
		message = "Файл создан"
	}
	defer file.Close()

	return message, err
}

func (user *User) GetFile(filename string) (string, string, string) {
	var message, err, text string
	for _, v := range user.File {
		if v == filename {
			message = "Файл найден"
			fContent, _ := ioutil.ReadFile("userFolder/" + user.Login + "/" + filename + ".txt")
			text = string(fContent)
			return message, err, text
		}
	}
	err = "Нет такого файла"
	return message, err, text
}

func (user *User) ChangeFile(filename, filetext string) (string, string) {
	var message, err string
	for k, v := range user.File {
		if v == filename {
			os.Remove("userFolder/" + user.Login + "/" + filename + ".txt")
			user.File = append(user.File[:k], user.File[k+1:]...)
			user.AddNewFile(filename, filetext)
			message = "Файл изменён"
			return message, err
		}
	}
	err = "Нет такого файла"
	return message, err
}

func (user *User) DeleteFile(filename string) (string, string) {
	var message, err string
	for k, v := range user.File {
		if v == filename {
			os.Remove("userFolder/" + user.Login + "/" + filename + ".txt")
			user.File = append(user.File[:k], user.File[k+1:]...)
			message = "Файл удалён"
			return message, err
		}
	}
	err = "Нет такого файла"
	return message, err
}

func (user *User) GetAllFile() []string {
	return user.File
}
