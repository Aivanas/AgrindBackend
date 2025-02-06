package main

import (
	"AgringBackend/config"
	"AgringBackend/models"
	"AgringBackend/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
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
		&models.User{},
		&models.UsersRoles{},
		&models.Role{},
		&models.Field{},
		&models.FieldStatus{},
		&models.Crop{},
		&models.WorkType{},
		&models.FieldWorkers{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	// Инициализация Gin
	r := gin.Default()
	routes.RegisterRoutes(r)

	// Запуск сервера
	if err := r.Run(":8085"); err != nil {
		log.Fatal(err)
	}
}
