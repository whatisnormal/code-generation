package repositories

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	domain "github.com/whatisnormal/code-generation/domain"
)

type fileRepository struct {
}

//NewFileRepo TODO
func NewFileRepo() domain.CodeManagerRepository {
	return &fileRepository{}
}

func (fr *fileRepository) Save(context string,
	id string, code string) error {

	dirPath := fmt.Sprintf("/tmp/context/%v", context)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, 0700) // Create your file
	}
	filePath := fmt.Sprintf("%v/%v.txt", dirPath, id)

	err := ioutil.WriteFile(filePath, []byte(code), 0644)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (fr *fileRepository) Find(context string, id string) (string, error) {
	dirPath := fmt.Sprintf("/tmp/context/%v", context)
	filePath := fmt.Sprintf("%v/%v.txt", dirPath, id)

	file, openError := os.Open(filePath)
	defer file.Close()

	if openError != nil {
		return "", openError
	}

	b, readError := ioutil.ReadAll(file)
	if readError != nil {
		log.Fatal(readError)
	}

	return string(b), nil
}
