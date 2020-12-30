package domain

type GenerateReq struct {
    APIKey string
    Context string
    ID string
}

type ValidateReq struct {
    APIKey string
    Context string
    ID string
    Code string
}