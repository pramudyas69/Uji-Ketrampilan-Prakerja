package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"uji/domain"
	"uji/helpers"
)

type UserHandler struct {
	userUseCase domain.UserUseCase
}

func NewUserHandler(e *echo.Echo, userUc domain.UserUseCase) {
	handler := UserHandler{
		userUseCase: userUc,
	}
	router := e.Group("/users")
	router.POST("/register", handler.RegiterUser)
	router.POST("/login", handler.LoginUser)
}

func (h *UserHandler) RegiterUser(ctx echo.Context) error {
	user := new(domain.User)

	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.WebResponse{
			Status: "BAD_REQUEST",
			Code:   400,
			Data:   "Json Payload Invalid!",
		})
	}

	err := h.userUseCase.UserRegisterUc(user)
	if err != nil {
		return helpers.ErrorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, helpers.WebResponse{
		Status: "CREATED",
		Code:   201,
		Data:   user,
	})
}

func (h *UserHandler) LoginUser(ctx echo.Context) error {
	user := new(domain.User)

	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.WebResponse{
			Status: "BAD_REQUEST",
			Code:   400,
			Data:   "Json Payload Invalid!",
		})
	}
	password := user.Password
	err := h.userUseCase.UserLoginUc(user)
	if err != nil {
		return helpers.ErrorHandler(ctx, err)
	}

	comparePass := helpers.ComparePassword([]byte(user.Password), []byte(password))
	if !comparePass {
		return ctx.JSON(http.StatusBadRequest, helpers.WebResponse{
			Status: "BAD_REQUEST",
			Code:   400,
			Data:   "Comparing Password Invalid!",
		})
	}

	token := helpers.GenerateToken(user.ID, user.Email)

	return ctx.JSON(http.StatusOK, helpers.WebResponse{
		Status: "OK",
		Code:   200,
		Data:   token,
	})
}
