package services

import (
	"assignment2/database"
	"assignment2/models"
	"errors"
	"fmt"
)

func GetOrders(orders *[]models.Order) error {
	db := database.GetDB()

	err := db.Model(&models.Order{}).Preload("Items").Order("order_id").Find(&orders).Error
	if err != nil {
		fmt.Println("error get orders", err)
		return errors.New("error get orders")
	}

	return nil
}

func GetOrderById(order *models.Order, orderId int) error {
	db := database.GetDB()
	order.ID = orderId

	err := db.Model(&order).Preload("Items").Find(&order).Error
	if err != nil {
		fmt.Println("error get data by id", err)
		return errors.New("error get order by id")
	}

	return nil
}

func CreateOrder(order models.Order) (models.Order, error) {
	db := database.GetDB()

	err := db.Create(&order).Error
	if err != nil {
		return models.Order{}, err
	}

	orderId := order.ID
	data := models.Order{}
	err = GetOrderById(&data, orderId)

	return data, err
}

func UpdateOrderById(newOrder *models.Order) error {
	db := database.GetDB()

	err := db.First(&models.Order{}, newOrder.ID).Error
	if err != nil {
		return errors.New("order with id not found")
	}

	err = db.Model(&models.Order{}).Where("order_id = ?", newOrder.ID).Updates(&newOrder).Error
	if err != nil {
		return errors.New("error update order data")
	}

	return nil
}

func DeleteOrderById(orderId int) error{
	db := database.GetDB()

	err := db.First(&models.Order{}, orderId).Error
	if err != nil {
		return errors.New("order with id not found")
	}

	err = db.Delete(&models.Order{}, orderId).Error
	if err != nil {
		return errors.New("error delete data order")
	}

	return nil
}
