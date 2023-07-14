package middlewares

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"uji/domain"
	"uji/helpers"
)

func UserAuthorization(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userId, err := strconv.Atoi(c.Param("Id"))
			if err != nil {
				return c.JSON(http.StatusBadRequest, helpers.WebResponse{
					Status: "BAD_REQUEST",
					Code:   400,
					Data:   "Invalid Param!",
				})
			}
			userData := c.Get("userData").(jwt.MapClaims)
			userID := uint32(userData["id"].(float64))

			User := domain.User{}
			err = db.Select("id").First(&User, uint32(userId)).Error

			if err != nil {
				return c.JSON(http.StatusNotFound, helpers.WebResponse{
					Status: "NOT_FOUND",
					Code:   404,
					Data:   "Data Not Found!",
				})
			}

			if User.ID != userID {
				return c.JSON(http.StatusUnauthorized, helpers.WebResponse{
					Status: "UNAUTHORIZED",
					Code:   401,
					Data:   "you're not allowed to access this data!",
				})
			}

			return next(c)
		}
	}
}
