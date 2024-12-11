package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Получение параметров подключения из переменных окружения
	dbUser := os.Getenv("DB_USER")         // Имя пользователя
	dbPassword := os.Getenv("DB_PASSWORD") // Пароль
	dbHost := os.Getenv("DB_HOST")         // Хост (например, localhost)
	dbPort := os.Getenv("DB_PORT")         // Порт (по умолчанию 3306)
	dbName := os.Getenv("DB_NAME")         // Имя базы данных

	// Формирование строки подключения
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Подключение к MySQL
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Database connection established successfully")
}
