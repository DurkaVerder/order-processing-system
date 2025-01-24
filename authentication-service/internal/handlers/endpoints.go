package handlers

import (
	"fmt"
	"log"
	"net/http"

	common "github.com/DurkaVerder/common-for-order-processing-system/models"
	"github.com/gin-gonic/gin"
)

func (h *HandlersManager) Login(c *gin.Context) {
	dataLogin := common.AuthDataLogin{}
	if err := c.BindJSON(&dataLogin); err != nil {
		log.Printf("Invalid data: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	token, err := h.service.Login(dataLogin)
	if err != nil {
		fmt.Printf("Error login: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, token)
}

func (h *HandlersManager) Register(c *gin.Context) {
	dateRegister := common.AuthDataRegister{}
	if err := c.BindJSON(&dateRegister); err != nil {
		log.Printf("Error unmarshal data: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if err := h.service.Register(dateRegister); err != nil {
		log.Printf("Error register: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}

func (h *HandlersManager) Logout(c *gin.Context) {
	token := common.Token{
		Token: c.Query("token"),
	}
	if err := h.service.Logout(token); err != nil {
		log.Printf("Error logout: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (h *HandlersManager) ValidateToken(c *gin.Context) {
	token := common.Token{
		Token: c.Query("token"),
	}

	if err := h.service.ValidateToken(token); err != nil {
		log.Printf("Error validate token: %s", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Token is valid"})
}
