package app

import (
	"fmt"
	api "github.com/whatisnormal/code-generation/controllers"
	domain "github.com/whatisnormal/code-generation/domain"
	"github.com/whatisnormal/code-generation/notifiers"
	"github.com/whatisnormal/code-generation/repositories"
	"flag"
	"log"
)

// repo <- service -> serializer  -> http

func defineRoutes() {
	originMsisdn := flag.String("TWILLIO_FROM_MSISDN","", "Your Twillio Origin phone number." )
	accountSid := flag.String("TWILLIO_ACCOUNT_SID","", "Your Twillio Account ID." )
	authToken := flag.String("TWILLIO_AUTH_TOKEN","", "Your Twillio Auth Token." )

	flag.Parse()

	repo := chooseRepo()
	notifier, notifierError := chooseNotifier(*originMsisdn,
							  *accountSid,
							  *authToken)
	if notifierError != nil {
		log.Fatal("Notifier could not be loaded.")
		return
	}

	service := domain.NewCodeManagerService(repo, notifier)

	handler := api.NewHandler(service)

	router.GET("/code-generation/ping/:api_key", handler.Ping)
	router.POST("/code-generation/generate/:api_key", handler.GenerateCode)
	router.GET("/code-generation/validate/:api_key", handler.ValidateCode)
}

func chooseRepo() domain.CodeManagerRepository {
	return repositories.NewFileRepo()
}

func chooseNotifier(originMsisdn string,
				  accountSid string,
				  authToken string) (domain.Notifier, error) {
	if originMsisdn == "" || accountSid == "" || authToken == "" {
		return nil, fmt.Errorf("empty parameter value for sms notification")
	}
	return notifiers.NewSmsNotifier(originMsisdn, accountSid, authToken), nil
}
