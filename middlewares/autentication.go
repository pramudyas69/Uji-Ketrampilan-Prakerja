package middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"uji/helpers"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		verifyToken, err := helpers.VerifyToken(c)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, helpers.WebResponse{
				Status: "UNAUTHORIZED",
				Code:   401,
				Data:   err.Error(),
			})

		}

		c.Set("userData", verifyToken)
		return next(c)
	}
}
