package domain

import (
	"fmt"
)

var inMemStore = make(map[string]string)

type CodeManagerRepository interface {
	Save(context string, id string, code string) error

	Find(context string, id string) (string, error)
}

type inMemRepository struct {
}

//NewFileRepo TODO
func NewInMemRepo() CodeManagerRepository {
	return &inMemRepository{}
}

func (ir *inMemRepository) Save(context string,
	id string, code string) error {

	if context == "" || id == "" || code == "" {
		return fmt.Errorf("Repository requires non empty arguments")
	}

	key := fmt.Sprintf("%v/%v", context, id)

	inMemStore[key] = code

	return nil
}

func (ir *inMemRepository) Find(context string, id string) (string, error) {

	if context == "" || id == "" {
		return "", fmt.Errorf("Repository requires non empty arguments")
	}
	return "", nil
}
