package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

//StartApp Initializes controller routes and server in port 8000
func StartApp() {
	mapURLs()
	if err := router.Run(":8000"); err != nil {
		panic(err)
	}
}