package domain



type CodeManagerRepository interface {
    Save(context string, id string, code string) (error)

    Find(context string, id string) (string, error)
}