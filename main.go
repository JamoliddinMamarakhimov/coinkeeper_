package main

import (
	"coinkeeper/configs"
	"coinkeeper/db"
	"coinkeeper/logger"
	"coinkeeper/pkg/controllers"
	"coinkeeper/server"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title COIN_KEEPER API
// @version 1.0
// @description Coin-Keeper: Finance.

// @host localhost:8181
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Ошибка загрузки .env файла: %s", err)
	}

	if err := configs.ReadSettings(); err != nil {
		log.Fatal("Ошибка чтения настроек: %s", err)
	}

	if err := logger.Init(); err != nil {
		log.Fatal("Ошибка инициализации логгера: %s", err)
	}

	var err error
	err = db.ConnectToDB()
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных: %s", err)
	}

	if err = db.Migrate(); err != nil {
		log.Fatal("Ошибка миграции базы данных: %s", err)
	}

	mainServer := new(server.Server)
	go func() {
		if err = mainServer.Run(configs.AppSettings.AppParams.PortRun, controllers.InitRoutes()); err != nil {
			log.Println("Ошибка при запуске HTTP сервера: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	fmt.Printf("\nНачало завершение программ\n")

	//Close DB
	if sqlDB, err := db.GetDBConn().DB(); err == nil {
		if err := sqlDB.Close(); err != nil {
			log.Fatal("Ошибка при закрытии соединения с БД: %s", err)
		}
	} else {
		log.Fatal("Ошибка при получении *sql.DB из GORM: %s", err)
	}
	fmt.Println("Соединение с БД успешно закрыто")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = mainServer.Shutdown(ctx); err != nil {
		log.Fatal("Ошибка при завершении работы сервера: %s", err)
	}

	fmt.Println("HTTP-сервис успешно выключен")
	fmt.Println("Конец завершения программы")
}
