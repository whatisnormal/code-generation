package app

import (
	api "github.com/whatisnormal/code-generation/controllers"
	domain "github.com/whatisnormal/code-generation/domain"
	repo "github.com/whatisnormal/code-generation/repository/file"
)

// repo <- service -> serializer  -> http


func defineRoutes() {
    //TODO Change
    repo := chooseRepo()
    service := domain.NewCodeManagerService(repo)

	handler := api.NewHandler(service)
	
    router.GET("/code-generation/ping/:api_key",handler.Ping)
    router.POST("/code-generation/generate/:api_key", handler.GenerateCode)
    router.GET("/code-generation/validate/:api_key",handler.ValidateCode)
}

func chooseRepo() domain.CodeManagerRepository {
    return repo.NewFileRepo()
}