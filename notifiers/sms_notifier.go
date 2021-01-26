package notifiers

import (
	domain "github.com/whatisnormal/code-generation/domain"
	"log"
)

type smsNotifier struct {
}

//NewFileRepo TODO
func NewSmsNotifier() domain.Notifier {
	return &smsNotifier{}
}

func (s smsNotifier) Notify(id string, code string) error {

	log.Printf("Notified id: %v with code: %v", id, code)

	return nil
}
