package http

import (
	"bytes"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

const templatePath = "templates/front"

func renderTemplates(tmplNames []string) string {
	tempBuffer := new(bytes.Buffer)
	var templatesPath []string

	for _, tempName := range tmplNames {
		templatesPath = append(templatesPath, templatePath+"/"+tempName)
	}

	tpl, _ := template.ParseFiles(
		templatesPath...,
	)

	_ = tpl.Execute(tempBuffer, nil)

	return tempBuffer.String()
}

func FrontViews(e *echo.Echo) {
	e.Static("/static", "static")

	e.GET("/", renderCallback)
	e.GET("/reviews", renderReviews)
}

func renderCallback(c echo.Context) error {
	return c.HTML(http.StatusOK, renderTemplates([]string{`app-shell.html`, `simulate-callback.html`}))
}

func renderReviews(c echo.Context) error {
	return c.HTML(http.StatusOK, renderTemplates([]string{`app-shell.html`, `reviews.html`}))
}
