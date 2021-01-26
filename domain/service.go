package domain

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

var (
	errRedirectNotFound = errors.New("Redirect Not Found")
	errRedirectInvalid  = errors.New("Redirect Invalid")
)

var letters = []rune("123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//CodeManagerService TODO
type CodeManagerService interface {
	Generate(r *GenerateReq) (string, error)

	Validate(r *ValidateReq) (bool, error)
}

//NewCodeManagerService TODO
func NewCodeManagerService(c CodeManagerRepository,
	n Notifier) CodeManagerService {

	return &codeManager{
		codeManagerRepository: c,
		notifier:              n,
	}
}

type codeManager struct {
	codeManagerRepository CodeManagerRepository
	notifier              Notifier
}

//Generates a code for a specific ID depending on the context,
//followed by a notification.
func (c *codeManager) Generate(r *GenerateReq) (string, error) {

	//Generate CODE
	code := randSeq(4)

	// persist
	saveErr := c.codeManagerRepository.Save(r.Context, r.ID, code)
	if saveErr != nil {
		return "", fmt.Errorf("Error while persisting code %v for context %v and id %v", code, r.Context, r.ID)
	}

	// send notification
	notifyErr := c.notifier.Notify(r.ID, code)
	if notifyErr != nil {
		return "", fmt.Errorf("Error while notifying code %v for context %v and id %v. Try generating again", code, r.Context, r.ID)
	}

	return code, nil
}

func (c *codeManager) Validate(r *ValidateReq) (bool, error) {
	// Fetches the code from the repo with the specified arguments.
	code, err := c.codeManagerRepository.Find(r.Context, r.ID)

	//If there was an errhor while trying to find the record for the given argument, then return.
	if err != nil {
		return false, fmt.Errorf("Error while validating code %v for context: %v and id: %v", r.Code, r.Context, r.ID)
	}

	matches := r.Code == code

	log.Printf("Does given code %v match %v for context %v and id %v? %t", r.Code, code, r.Context, r.ID, matches)

	return matches, nil
}
