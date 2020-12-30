package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	domain "github.com/whatisnormal/code-generation/domain"
)

type generateCodeReq struct {
    ID string `json:id`
 }

type ControllerHandler interface {
    Ping(c *gin.Context)
    GenerateCode(c *gin.Context)
    ValidateCode(c *gin.Context)
}

 type handler struct{
     codeManagerService domain.CodeManagerService
 }

 func NewHandler(c domain.CodeManagerService) ControllerHandler{
     return &handler{
        c,
     }
 }

 // Ping Just a simple ping
 func (h *handler) Ping(c *gin.Context){
    apiKey := c.Param("api_key")
    c.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("You made it %v",apiKey )})
}

// GenerateCode takes a gin.Context and serves a response as Json
func (h *handler) GenerateCode(c *gin.Context) {
	var json generateCodeReq
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
        }
    
    apiKey := c.Param("api_key")

    msisdn := json.ID

    code, err := h.codeManagerService.Generate(&domain.GenerateReq{
        APIKey: apiKey,
        Context: "default",
        ID: msisdn,
    })

    if err != nil{
		c.JSON(http.StatusInternalServerError, err)
		return        
    }

    ///c.JSON(http.StatusOK, country)
    c.JSON(http.StatusOK, gin.H{"code": code})
}

//ValidateCode Checks if the provided code is valid for the given identifiacation.
func (h *handler)  ValidateCode(c *gin.Context){
//    h.codeManagerService.
    c.JSON(http.StatusOK, gin.H{"status": "You made it"})
}