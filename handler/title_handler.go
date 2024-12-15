package handler

import (
	"betteryou/view/layout"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct {}

func (h HomeHandler) HomeGetHandler(c echo.Context) error {
	return render(c, layout.Home())
}
