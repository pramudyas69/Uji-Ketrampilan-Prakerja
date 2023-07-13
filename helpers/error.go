package helpers

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ErrorHandler(ctx echo.Context, err error) error {

	_, ok := err.(govalidator.Errors)
	if !ok {
		return ctx.JSON(http.StatusBadRequest, WebResponse{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	return ctx.JSON(http.StatusInternalServerError, WebResponse{
		Code:   500,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})
}
