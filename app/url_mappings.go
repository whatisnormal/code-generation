package app

import (
	api "github.com/whatisnormal/code-generation/controllers"
	domain "github.com/whatisnormal/code-generation/domain"
	"github.com/whatisnormal/code-generation/notifiers"
	"github.com/whatisnormal/code-generation/repositories"
)

// repo <- service -> serializer  -> http

func defineRoutes() {
	//TODO Change
	repo := chooseRepo()
	notifier := chooseNotifier()
	service := domain.NewCodeManagerService(repo, notifier)

	handler := api.NewHandler(service)

	router.GET("/code-generation/ping/:api_key", handler.Ping)
	router.POST("/code-generation/generate/:api_key", handler.GenerateCode)
	router.GET("/code-generation/validate/:api_key", handler.ValidateCode)
}

func chooseRepo() domain.CodeManagerRepository {
	return repositories.NewFileRepo()
}

func chooseNotifier() domain.Notifier {
	return notifiers.NewSmsNotifier()
}
