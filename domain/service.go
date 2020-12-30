package domain

import "errors"

var (
	errRedirectNotFound = errors.New("Redirect Not Found")
	errRedirectInvalid  = errors.New("Redirect Invalid")
)

//CodeManagerService TODO 
type CodeManagerService interface {
    Generate(r *GenerateReq) (string, error)

    Validate(r *ValidateReq) (bool, error)
}

//NewCodeManagerService TODO
func NewCodeManagerService(c CodeManagerRepository) CodeManagerService{
    return &codeManager{
        c,
    }
}

type codeManager struct {
    codeManagerRepository CodeManagerRepository
}

func (c *codeManager) Generate(r *GenerateReq) (string, error){
    code := "ABCD"
    return c.codeManagerRepository.Save(r.Context, r.ID, code)
}

func (c *codeManager) Validate(r *ValidateReq) (bool, error){
    return c.codeManagerRepository.Find(r.Context, r.ID)
}