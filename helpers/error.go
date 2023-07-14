package helpers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ErrorHandler(ctx echo.Context, err error) error {
	//_, ok := err.(govalidator.Errors)
	isValid := IsDuplicateError(err)
	switch {
	//case !ok:
	//	fmt.Println(err.(govalidator.Errors))
	//	return ctx.JSON(http.StatusBadRequest, WebResponse{
	//		Code:   400,
	//		Status: "BAD_REQUEST",
	//		Data:   err.Error(),
	//	})
	case !isValid:
		return ctx.JSON(http.StatusBadRequest, WebResponse{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	default:
		return ctx.JSON(http.StatusInternalServerError, WebResponse{
			Code:   500,
			Status: "INTERNAL_SERVER_ERROR",
			Data:   err.Error(),
		})
	}

}
