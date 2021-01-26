package domain

import "log"

type Notifier interface {
	Notify(id string, code string) error
}

type dummyNotifier struct {
}

func (d dummyNotifier) Notify(id string, code string) error {

	log.Printf("Notified id: %v with code: %v", id, code)

	return nil
}
