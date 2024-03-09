package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vukovlevi/kviz-go/db"
	"github.com/vukovlevi/kviz-go/views"
	"golang.org/x/crypto/bcrypt"
)

func HandleLoginIndex(c echo.Context) error {
    return render(c, views.Default(views.LoginIndex()))
}

func HandleLogin(c echo.Context) error {
    username := c.FormValue("username")
    password := c.FormValue("password")

    user := db.GetUserByUsername(username)
    if user == nil {
        return render(c, views.Error("Nem létező felhasználó!"))
    }
    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

    if err != nil {
        return render(c, views.Error("Helytelen jelszó"))
    }

    c.SetCookie(&http.Cookie{Name: "token", Value: "asdfghjkl", HttpOnly: true})
    return render(c, views.Default(views.LoginSuccess()))
}

func MustLogin(c echo.Context) error {
    return render(c, views.MustLogin())
}
