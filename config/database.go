package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Получение параметров подключения из переменных окружения
	dbUser := os.Getenv("DB_USER")           // Имя пользователя
	dbPassword := os.Getenv("DB_PASSWORD")   // Пароль
	dbHost := os.Getenv("DB_HOST")           // Хост (например, localhost)
	dbPort := os.Getenv("DB_PORT")           // Порт (по умолчанию 5432)
	dbName := os.Getenv("DB_NAME")           // Имя базы данных
	sslMode := os.Getenv("DB_SSLMODE")       // SSL режим (например, disable, require)
	searchPath := os.Getenv("DB_SEARCHPATH") // Schema Search Path (если требуется)

	// Формирование строки подключения PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s search_path=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, sslMode, searchPath,
	)

	// Подключение к PostgreSQL через GORM
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Database connection established successfully")

}
