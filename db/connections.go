package db

import (
	"coinkeeper/configs"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var dbConn *gorm.DB

func ConnectToDB() error {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Вывод логов в стандартный вывод (os.Stdout)
		logger.Config{
			SlowThreshold: time.Second, // Порог для медленных запросов
			LogLevel:      logger.Info, // Уровень логирования (вывод всех SQL-запросов)
			Colorful:      true,        // Цветной вывод логов
		},
	)

	connStr := fmt.Sprintf(
		`host=%s port=%s user=%s dbname=%s password=%s`,
		configs.AppSettings.PostgresParams.Host,
		configs.AppSettings.PostgresParams.Port,
		configs.AppSettings.PostgresParams.User,
		configs.AppSettings.PostgresParams.Database,
		os.Getenv("DB_PASSWORD"),
	)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return err
	}
	fmt.Println("Successfully connected to DB")
	dbConn = db
	return nil
}

func CloseDBConn() error {
	return nil
}

func GetDBConn() *gorm.DB {
	return dbConn
}
