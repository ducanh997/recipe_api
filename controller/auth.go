package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
}

func (*AuthController) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func NewAuthController() *AuthController {
	return &AuthController{}
}
