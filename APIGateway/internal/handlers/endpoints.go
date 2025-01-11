package handlers

import (
	"APIGateway/common"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HandlersManager) HandlerLogin(c *gin.Context) {
	loginData := &common.AuthData{}

	if err := c.BindJSON(&loginData); err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	url := StartURL + h.cfg.Authentication.Server.Port + h.cfg.Authentication.Route.Base + h.cfg.Authentication.Route.Endpoints["login"]
	res, err := h.requester.SendRequest(url, http.MethodPost, loginData)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	if res.StatusCode == http.StatusBadRequest {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid login or password"})
		return
	}

	if res.StatusCode != http.StatusOK {
		c.JSON(res.StatusCode, gin.H{"error": "internal error"})
		return
	}

	token := common.Token{}
	if err := h.requester.UnmarshalResponse(res, &token); err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.SetCookie("jwt", token.Token, 3600*72, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "login successful"})

}

func (h *HandlersManager) HandlerRegister(c *gin.Context) {
	registerData := &common.AuthData{}

	if err := c.BindJSON(&registerData); err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	url := StartURL + h.cfg.Authentication.Server.Port + h.cfg.Authentication.Route.Base + h.cfg.Authentication.Route.Endpoints["register"]
	res, err := h.requester.SendRequest(url, http.MethodPost, registerData)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	if res.StatusCode == http.StatusBadRequest {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid login or password"})
		return
	}

	if res.StatusCode != http.StatusOK {
		c.JSON(res.StatusCode, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "register successful"})

}

func (h *HandlersManager) HandlerLogout(c *gin.Context) {
	token, err := c.Cookie("jwt")
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	url := StartURL + h.cfg.Authentication.Server.Port + h.cfg.Authentication.Route.Base + h.cfg.Authentication.Route.Endpoints["logout"] + "?token=" + token
	res, err := h.requester.SendRequest(url, http.MethodGet, nil)
	if err != nil || res.StatusCode != http.StatusOK {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}

func (h *HandlersManager) HandlerCreateOrder(c *gin.Context) {

}

func (h *HandlersManager) HandlerGetOrders(c *gin.Context) {

}

func (h *HandlersManager) HandlerGetOrder(c *gin.Context) {

}

func (h *HandlersManager) HandlerDeleteOrder(c *gin.Context) {

}

func (h *HandlersManager) HandlerStatusOrder(c *gin.Context) {

}

func (h *HandlersManager) HandlerHistoryOrder(c *gin.Context) {

}
