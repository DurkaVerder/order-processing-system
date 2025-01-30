package handlers

import (
	"log"
	"net/http"
	"strconv"

	common "github.com/DurkaVerder/common-for-order-processing-system/models"
	"github.com/gin-gonic/gin"
)

func (h *HandlersManager) HandlerAddOrder(c *gin.Context) {
	order := common.Order{}
	if err := c.BindJSON(&order); err != nil {
		log.Printf("Error while binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while binding JSON"})
		return
	}
	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		log.Printf("Error while parsing userId: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while parsing userId"})
		return
	}

	order.UserId = userId
	if err := h.service.AddOrder(order); err != nil {
		log.Printf("Error while adding order: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while adding order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order added"})

}

func (h *HandlersManager) HandlerGetOrder(c *gin.Context) {
	orderId := c.Query("order_id")
	order, err := h.service.GetOrder(orderId)
	if err != nil {
		log.Printf("Error while getting order: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting order"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *HandlersManager) HandlerGetAllOrders(c *gin.Context) {
	userId := c.Query("user_id")
	orders, err := h.service.GetAllOrders(userId)
	if err != nil {
		log.Printf("Error while getting orders: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *HandlersManager) HandlerDeleteOrder(c *gin.Context) {
	orderId := c.Query("order_id")
	if err := h.service.DeleteOrder(orderId); err != nil {
		log.Printf("Error while deleting order: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while deleting order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
}
