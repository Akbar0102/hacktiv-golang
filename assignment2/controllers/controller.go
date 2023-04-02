package controllers

import (
	"assignment2/models"
	"assignment2/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderResponse struct {
	Data []models.Order `json:"data"`
}

func GetOrders(ctx *gin.Context) {
	var orders []models.Order

	err := services.GetOrders(&orders)
	orderResponse := OrderResponse{Data: orders}
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, orderResponse)

}

func CreateOrder(ctx *gin.Context) {
	var newOrder models.Order

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "error json request",
		})
		return
	}

	orderResult, err := services.CreateOrder(newOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, orderResult)
}

func UpdateOrderById(ctx *gin.Context) {
	var newOrder models.Order
	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "error json request",
		})
		return
	}

	orderId := ctx.Param("orderId")
	id, err := strconv.Atoi(orderId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "error order id not recognized",
		})
		return
	}

	newOrder.ID = id
	err = services.UpdateOrderById(&newOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    newOrder,
		"message": "Update data success",
		"success": true,
	})

}

func DeleteOrderById(ctx *gin.Context) {
	orderId := ctx.Param("orderId")

	id, err := strconv.Atoi(orderId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "error order id not recognized",
		})
		return
	}

	err = services.DeleteOrderById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete data success",
		"success": true,
	})
}
