package handler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component, status int) error {
	c.Response().WriteHeader(status)
	return component.Render(c.Request().Context(), c.Response())
}
