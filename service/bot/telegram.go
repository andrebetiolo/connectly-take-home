package bot

import (
	"bytes"
	"connectly/model"
	"connectly/service/products"
	"connectly/service/reviews"
	"connectly/service/users"
	"encoding/json"
	"fmt"
	"log"
	"text/template"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type service struct {
	config          Config
	telegramBot     *tgbotapi.BotAPI
	usersService    users.Service
	productsService products.Service
	reviewsService  reviews.Service
}

func NewTelegram(
	c Config,
	usersService users.Service,
	productsService products.Service,
	reviewsService reviews.Service,
) (Service, error) {
	return &service{
		config:          c,
		usersService:    usersService,
		productsService: productsService,
		reviewsService:  reviewsService,
	}, nil
}

func (s *service) ListenMessages() error {
	bot, err := tgbotapi.NewBotAPI(s.config.ApiToken)
	if err != nil {
		log.Println("Error on init bot", err)
		return err
	}

	bot.Debug = true

	s.telegramBot = bot

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := s.telegramBot.GetUpdatesChan(u)

	// Loop through each update.
	for update := range updates {
		var message string

		if update.Message != nil {
			if update.Message == nil {
				continue
			}

			isCommand := update.Message.IsCommand()

			if isCommand {
				switch update.Message.Command() {
				case "start":
					user := messageToUser(*update.Message)
					_, err := s.usersService.CreateOrUpdate(user)
					if err != nil {
						log.Println("Error on create or update user", err, user)
					}

					message = fmt.Sprintf(`Welcome to ConnectlyBot %s!🎉`, user.FirstName)

					s.SendMessage(user.BotID, message)
					continue
				default:
					message = "I don't know that command"
					s.SendMessage(update.Message.From.ID, message)
				}
			}
		}

		if update.CallbackQuery != nil {
			userID := update.CallbackQuery.Message.Chat.ID
			callbackData := update.CallbackQuery.Data

			data := map[string]interface{}{}

			if err := json.Unmarshal([]byte(callbackData), &data); err != nil {
				log.Print(`Error on ummarshal callback data`, err)
			}

			switch data["flow"] {
			case FlowReviewProduct:
				err := s.reviewsService.CreateOrUpdate(data["user_id"].(float64), data["product_id"].(float64), data["rate"].(float64))
				if err != nil {
					log.Println(`error on save review`, err)
				}

				s.SendMessage(userID, `Thanks for you feedback!`)
				continue
			case FlowReturnOrder:
				s.StartFlow(FlowReturnOrder, data)
				continue
			}
		}
	}

	return nil
}

func (s *service) StartFlow(flow string, parameters map[string]interface{}) error {
	var user model.User
	var err error
	if parameters["user_id"] != nil {
		user, err = s.usersService.Get(parameters["user_id"].(float64))
		if err != nil {
			log.Println(`error on get user`, err)
		}
	}

	var product model.Product
	if parameters["product_id"] != nil {
		product, err = s.productsService.Get(parameters["user_id"].(float64))
		if err != nil {
			log.Println(`error on get user`, err)
		}
	}

	switch flow {
	case FlowReviewProduct:
		params := map[string]interface{}{
			"user_name":    user.FirstName,
			"product_name": product.Name,
		}

		message, err := renderTemplate(`review_product.tpt`, params)
		if err != nil {
			return err
		}

		buttons := map[string]string{
			"Liked very much": fmt.Sprintf(`{"flow":"review_product","user_id":%v,"product_id":%v,"rate":3}`, user.ID, product.ID),
			"Liked":           fmt.Sprintf(`{"flow":"review_product","user_id":%v,"product_id":%v,"rate":2}`, user.ID, product.ID),
			"Not Liked":       fmt.Sprintf(`{"flow":"review_product","user_id":%v,"product_id":%v,"rate":1}`, user.ID, product.ID),
			"Return order":    fmt.Sprintf(`{"flow":"return_order","user_id":%v,"product_id":%v}`, user.ID, product.ID),
		}

		keyboard := createKeyboardWithKeyValue(buttons)

		return s.sendMessageWithKeyboard(user.BotID, message, keyboard)
	case FlowReturnOrder:
		message := `We deeply regret any inconvenience experienced. Please contact us to process your order refund and ensure your complete satisfaction. We are here to resolve any issues and appreciate your understanding. 🙏 #CustomerService`

		return s.SendMessage(user.BotID, message)
	case FlowProductRecommendations:
		params := map[string]interface{}{
			"user_name": user.FirstName,
			"product_a": product.Name,
			"product_b": `Product Recomendation`,
		}

		message, err := renderTemplate(`product_recommendations.tpt`, params)
		if err != nil {
			return err
		}

		var keyboard = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonURL("Know more", "https://www.connectly.ai"),
			),
		)

		return s.sendMessageWithKeyboard(user.BotID, message, keyboard)
	default:
		return fmt.Errorf(`flow %s not defined`, flow)
	}
}

func (s *service) SendMessage(userId interface{}, message string) error {
	msg := tgbotapi.NewMessage(userId.(int64), message)

	_, err := s.telegramBot.Send(msg)

	return err
}

func (s *service) sendMessageWithKeyboard(userId interface{}, message string, keyboard tgbotapi.InlineKeyboardMarkup) error {
	msg := tgbotapi.NewMessage(userId.(int64), message)

	msg.ReplyMarkup = keyboard

	_, err := s.telegramBot.Send(msg)

	return err
}

func messageToUser(message tgbotapi.Message) model.User {
	return model.User{
		Username:  message.From.UserName,
		FirstName: message.From.FirstName,
		LastName:  message.From.LastName,
		BotID:     message.From.ID,
		BotName:   "telegram",
	}
}

func renderTemplate(tmplName string, params any) (string, error) {
	templatePath := "templates"
	tempBuffer := new(bytes.Buffer)

	fullPath := fmt.Sprintf(`%s/%s`, templatePath, tmplName)

	tpl, _ := template.ParseFiles(
		fullPath,
	)

	err := tpl.Execute(tempBuffer, params)
	if err != nil {
		return "", err
	}

	return tempBuffer.String(), nil
}

func createKeyboardWithKeyValue(options map[string]string) tgbotapi.InlineKeyboardMarkup {
	var optBottons []tgbotapi.InlineKeyboardButton

	for key, val := range options {
		optBottons = append(optBottons, tgbotapi.NewInlineKeyboardButtonData(key, val))
	}

	return tgbotapi.NewInlineKeyboardMarkup(
		// tgbotapi.NewInlineKeyboardRow(
		optBottons,
	// ),
	)
}
