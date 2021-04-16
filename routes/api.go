package routes
import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user": "DanhHan",
		"emai": "danhhm@antoree.com",
	})
}
func HandleSignUp(c echo.Context) error {
	type User struct {
		email string
		fullName string
	}
	user := User {
		email: "Danhhm1@antoree.com",
		fullName: "Danh Han 1",
	}
	// fmt.Println(user)
	return c.JSON(http.StatusOK, echo.Map{
		"user": user.email,
		"emai": user.fullName,
	})
}