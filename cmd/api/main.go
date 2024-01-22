package main

import (
	"fmt"
	"github.com/GoGerman/geo-task/run"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	godotenv.Load()
	// инициализация приложения
	app := run.NewApp()
	// запуск приложения
	err := app.Run()
	// в случае ошибки выводим ее в лог и завершаем работу с кодом 2
	if err != nil {
		log.Println(fmt.Sprintf("error: %s", err))
		os.Exit(2)
	}
}
