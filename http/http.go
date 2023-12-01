package http

import (
	"connectly/http/request"
	"connectly/repository"
	"connectly/service/bot"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yggbrazil/go-toolbox/api"
	"github.com/yggbrazil/go-toolbox/handler"
)

type Message struct {
	Message string `json:"message"`
}

var (
	errStartOnStartNewFlow = errors.New(`error on start a flow`)
)

func CreateHTTPServer(chatbot bot.Service, res repository.Repository) {
	s := api.Make()

	api.ProvideEchoInstance(FrontViews)

	s.POST("/api/callback/start-chatbot-flow", func(c echo.Context) error {
		p := c.Get(handler.PARAMETERS).(*request.SendMessageRequest)

		err := chatbot.StartFlow(p.Flow, p.Parameters)
		if err != nil {
			return errStartOnStartNewFlow
		}

		return c.JSON(http.StatusOK, Message{"Flow started"})
	}, handler.MiddlewareBindAndValidate(&request.SendMessageRequest{}))

	s.GET("/api/reviews", func(c echo.Context) error {
		reviews, err := res.GetAllReviews()
		if err != nil {
			return errStartOnStartNewFlow
		}

		return c.JSON(http.StatusOK, reviews)
	})

	api.Run()
}
