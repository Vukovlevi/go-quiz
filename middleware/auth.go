package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/vukovlevi/kviz-go/routes"
)

func CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        _, err := c.Cookie("token")
        if err != nil {
            return routes.MustLogin(c)
        }
        return next(c)
    }
}
