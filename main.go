package main

import (
	"AgringBackend/config"
	"AgringBackend/models"
	"AgringBackend/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Загрузка переменных окружения
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Подключение к базе данных
	config.ConnectDB()

	// Миграции
	err = config.DB.AutoMigrate(
		&models.Vehicle{},
		&models.VehicleCategory{},
		&models.VehicleStatus{},
		&models.Employee{},
		&models.EmployeeRole{},
		&models.Role{},
		&models.Field{},
		&models.FieldStatusInfo{},
		&models.FieldStatusList{},
		&models.Crop{},
	)
	if err != nil {
		return
	}

	// Настройка маршрутов
	r := gin.Default()
	routes.RegisterRoutes(r)

	// Запуск сервера
	err = r.Run("0.0.0.0:8082")
	if err != nil {
		return
	}
}
