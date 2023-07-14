package http

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"uji/domain"
	"uji/helpers"
	"uji/middlewares"
)

type UserHandler struct {
	userUseCase domain.UserUseCase
}

func NewUserHandler(e *echo.Echo, userUc domain.UserUseCase, db *gorm.DB) {
	handler := UserHandler{
		userUseCase: userUc,
	}
	router := e.Group("/users")
	router.POST("/register", handler.RegiterUser)
	router.POST("/login", handler.LoginUser)
	{
		router.Use(middlewares.Authentication)
		router.GET("/:Id", handler.GetUserById)
		router.GET("/", handler.GetUsers)
		router.Use(middlewares.UserAuthorization(db))
		router.PUT("/:Id", handler.UpdateUser)
		router.DELETE("/:Id", handler.DeleteUser)
	}

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
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.WebResponse{
			Status: "BAD_REQUEST",
			Code:   400,
			Data:   "Validation Error!",
		})
	}

	err = h.userUseCase.UserRegisterUc(user)
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
		fmt.Println(err.Error())
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

func (h *UserHandler) GetUserById(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("Id"))

	res, err := h.userUseCase.GetUserByIdUc(uint32(id))
	if err != nil {
		return helpers.ErrorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusOK, helpers.WebResponse{
		Status: "OK",
		Code:   200,
		Data:   res,
	})
}

func (h *UserHandler) GetUsers(ctx echo.Context) error {
	var user []*domain.User

	res, err := h.userUseCase.GetUsersUc(user)
	if err != nil {
		return helpers.ErrorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusOK, helpers.WebResponse{
		Status: "OK",
		Code:   200,
		Data:   res,
	})
}

func (h *UserHandler) UpdateUser(ctx echo.Context) error {
	newUser := new(domain.User)
	id, _ := strconv.Atoi(ctx.Param("Id"))

	if err := ctx.Bind(newUser); err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.WebResponse{
			Status: "BAD_REQUEST",
			Code:   400,
			Data:   "Json Payload Invalid!",
		})
	}

	_, err := govalidator.ValidateStruct(newUser)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.WebResponse{
			Status: "BAD_REQUEST",
			Code:   400,
			Data:   "Validation Error!",
		})
	}

	res, err := h.userUseCase.UpdateUserUc(uint32(id), newUser)
	if err != nil {
		fmt.Println(err)
		fmt.Println(newUser)
		return helpers.ErrorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusOK, helpers.WebResponse{
		Status: "OK",
		Code:   200,
		Data:   res,
	})
}

func (h *UserHandler) DeleteUser(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("Id"))

	err := h.userUseCase.DeleteUserUc(uint32(id))
	if err != nil {
		return helpers.ErrorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusOK, helpers.WebResponse{
		Status: "OK",
		Code:   200,
		Data:   "Succesfull Deleted!",
	})
}
