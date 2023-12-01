package main

import (
	"connectly/http"
	"connectly/repository"
	"connectly/service/bot"
	productsService "connectly/service/products"
	reviewsService "connectly/service/reviews"
	usersService "connectly/service/users"
	"os"

	"log"
)

func main() {
	cfgBot := bot.Config{
		Type:     "telegram",
		ApiToken: os.Getenv("CONNECTLY_TELEGRAM_API_TOKEN"),
	}

	cfgRepository := repository.Config{
		Type:         os.Getenv("DB_TYPE"),
		PathToDBFile: os.Getenv("DB_PATH_TO_FILE"),
	}

	res, err := repository.New(cfgRepository)
	if err != nil {
		log.Println("Error on init chatbot", err)
	}

	usersService := usersService.New(res)

	productsService := productsService.New(res)

	reviewsService := reviewsService.New(res)

	chabot, err := bot.NewTelegram(cfgBot, usersService, productsService, reviewsService)
	if err != nil {
		log.Println("Error on init chatbot", err)
	}

	go func() {
		chabot.ListenMessages()
	}()

	http.CreateHTTPServer(chabot, res)
}
