package domain



type Notifier interface {    

    Notify(id string, code string) error
    
}