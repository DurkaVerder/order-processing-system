package handlers

import (
	"log"
	"net/http"
	"strconv"

	common "github.com/DurkaVerder/common-for-order-processing-system/models"
	"github.com/gin-gonic/gin"
)

// HandlerLogin is a handler for login
func (h *HandlersManager) HandlerLogin(c *gin.Context) {
	loginData := common.AuthDataLogin{}

	if err := c.BindJSON(&loginData); err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	url := StartURLauth + h.cfg.Authentication.Server.Port + h.cfg.Authentication.Route.Base + h.cfg.Authentication.Route.Endpoints["login"]
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

// HandlerRegister is a handler for register
func (h *HandlersManager) HandlerRegister(c *gin.Context) {
	registerData := common.AuthDataRegister{}

	if err := c.BindJSON(&registerData); err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	url := StartURLauth + h.cfg.Authentication.Server.Port + h.cfg.Authentication.Route.Base + h.cfg.Authentication.Route.Endpoints["register"]
	res, err := h.requester.SendRequest(url, http.MethodPost, registerData)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error send request"})
		return
	}
	defer res.Body.Close()

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

// HandlerLogout is a handler for logout
func (h *HandlersManager) HandlerLogout(c *gin.Context) {
	token, err := c.Cookie("jwt")
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	url := StartURLauth + h.cfg.Authentication.Server.Port + h.cfg.Authentication.Route.Base + h.cfg.Authentication.Route.Endpoints["logout"] + "?token=" + token
	res, err := h.requester.SendRequest(url, http.MethodGet, nil)
	if err != nil || res.StatusCode != http.StatusOK {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}

// HandlerCreateOrder is a handler for creating order
func (h *HandlersManager) HandlerCreateOrder(c *gin.Context) {
	newOrder := common.Order{}
	if err := c.BindJSON(&newOrder); err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	userId, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userIdStr := strconv.Itoa(userId.(int))

	url := StartURLorder + h.cfg.Order.Server.Port + h.cfg.Order.Route.Base + h.cfg.Order.Route.Endpoints["create_order"] + "?user_id=" + userIdStr
	res, err := h.requester.SendRequest(url, http.MethodPost, newOrder)
	if err != nil || res.StatusCode != http.StatusCreated {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "order created",
		"order":   newOrder,
	})
}

// HandlerGetOrders is a handler for getting orders
func (h *HandlersManager) HandlerGetOrders(c *gin.Context) {
	userId, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userIdStr := strconv.Itoa(userId.(int))

	url := StartURLorder + h.cfg.Order.Server.Port + h.cfg.Order.Route.Base + h.cfg.Order.Route.Endpoints["get_orders"] + "?user_id=" + userIdStr
	res, err := h.requester.SendRequest(url, http.MethodGet, nil)
	if err != nil || res.StatusCode != http.StatusOK {
		log.Println("Error send request: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	orders := []common.Order{}
	if err := h.requester.UnmarshalResponse(res, &orders); err != nil {
		log.Println("Error unmarshal response: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// HandlerGetOrder is a handler for getting order
func (h *HandlersManager) HandlerGetOrder(c *gin.Context) {
	orderId := c.Param("order_id")

	url := StartURLorder + h.cfg.Order.Server.Port + h.cfg.Order.Route.Base + h.cfg.Order.Route.Endpoints["get_order"] + "?order_id=" + orderId
	res, err := h.requester.SendRequest(url, http.MethodGet, nil)
	if err != nil || res.StatusCode != http.StatusOK {
		log.Println("Error send request: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	order := common.Order{}
	if err := h.requester.UnmarshalResponse(res, &order); err != nil {
		log.Println("Error unmarshal response: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// HandlerDeleteOrder is a handler for deleting order
func (h *HandlersManager) HandlerDeleteOrder(c *gin.Context) {
	orderId := c.Param("order_id")

	url := StartURLorder + h.cfg.Order.Server.Port + h.cfg.Order.Route.Base + h.cfg.Order.Route.Endpoints["delete_order"] + "?order_id=" + orderId
	res, err := h.requester.SendRequest(url, http.MethodDelete, nil)
	if err != nil || res.StatusCode != http.StatusOK {
		log.Println("Error send request: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order deleted"})
}

// HandlerHistoryOrder is a handler for getting order history
func (h *HandlersManager) HandlerHistoryOrder(c *gin.Context) {
	orderId := c.Param("order_id")

	url := StartURLhistory + h.cfg.History.Server.Port + h.cfg.History.Route.Base + h.cfg.History.Route.Endpoints["history_order"] + "?order_id=" + orderId

	res, err := h.requester.SendRequest(url, http.MethodGet, nil)
	if err != nil || res.StatusCode != http.StatusOK {
		log.Println("Error send request: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	history := common.HistoryOrder{}

	if err := h.requester.UnmarshalResponse(res, &history); err != nil {
		log.Println("Error unmarshal response: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, history)
}
