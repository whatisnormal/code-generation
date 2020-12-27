package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type generateCodeReq struct {
    ID string `json:id`
 }

 // Ping Just a simple ping
 func Ping(c *gin.Context){
    apiKey := c.Param("api_key")
    c.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("You made it %v",apiKey )})
}

// GenerateCode takes a gin.Context and serves a response as Json
func GenerateCode(c *gin.Context) {
	var json generateCodeReq
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
        }
    
    apiKey := c.Param("api_key")

    msisdn := json.ID
    /*
	country, err := services.LocationsService.GetCountry(c.Param("country_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
    }
    */
    ///c.JSON(http.StatusOK, country)
    c.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("You made it %v and %v",apiKey, msisdn )})
}

//ValidateCode Checks if the provided code is valid for the given identifiacation.
func ValidateCode(c *gin.Context){

    c.JSON(http.StatusOK, gin.H{"status": "You made it"})
}