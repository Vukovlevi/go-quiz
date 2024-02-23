package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/vukovlevi/kviz-go/views"
)

func IndexHandler(c echo.Context) error {
	return render(c, views.Default(views.Index()))
} 