package app

import (
	"github.com/whatisnormal/code-generation/controllers"
)



func mapURLs() {
    //TODO Change

    router.GET("/code-generation/ping/:api_key",controllers.Ping)
    router.POST("/code-generation/generate/:api_key", controllers.GenerateCode)
    router.GET("/code-generation/validate/:api_key",controllers.ValidateCode)
}