package global

import (
	"math/rand"
	"os"
	"time"
)

func GenerateFileName(fileName, handler string) string {
	rand.Seed(time.Now().UnixNano())
	letters := ("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var newName string
	for i := 0; i < len(fileName); i++ {
		newName += string(letters[rand.Intn(len(letters))])
	}
	newName += "_" + handler + "_.pdf"
	return newName
}

func GeneratePath(fileName, handler string) string {
	rand.Seed(time.Now().UnixNano())
	letters := ("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var newPath string
	for i := 0; i < len(fileName); i++ {
		newPath += string(letters[rand.Intn(len(letters))])
	}
	newPath += "_" + handler
	return newPath
}

func MakeDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir|0755)
	}
	return nil
}
