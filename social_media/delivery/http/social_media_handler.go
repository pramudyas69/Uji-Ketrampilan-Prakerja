package http

import (
	"github.com/asaskevich/govalidator"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"uji/domain"
	"uji/helpers"
	"uji/middlewares"
)

type SosmedHandler struct {
	sosmedUseCase domain.SosmedUseCase
}

func NewSosmedHandler(e *echo.Echo, sosmedUc domain.SosmedUseCase) {
	handler := SosmedHandler{
		sosmedUseCase: sosmedUc,
	}
	router := e.Group("/sosmed")
	//router.POST("/login", handler.LoginUser)
	{
		router.Use(middlewares.Authentication)
		router.POST("", handler.CreateSosmed)
		//router.GET("/:Id", handler.GetUserById)
		//router.GET("/", handler.GetUsers)
		//router.Use(middlewares.UserAuthorization(db))
		//router.PUT("/:Id", handler.UpdateUser)
		//router.DELETE("/:Id", handler.DeleteUser)
	}
}

func (h *SosmedHandler) CreateSosmed(ctx echo.Context) error {
	sosmed := new(domain.SocialMedia)

	if err := ctx.Bind(&sosmed); err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.WebResponse{
			Status: "BAD_REQUEST",
			Code:   400,
			Data:   "Json Payload Invalid!",
		})
	}
	_, err := govalidator.ValidateStruct(sosmed)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.WebResponse{
			Status: "BAD_REQUEST",
			Code:   400,
			Data:   "Validation Error!",
		})
	}

	userData := ctx.Get("userData").(jwt.MapClaims)
	userID := uint32(userData["id"].(float64))

	sosmed.UserID = userID
	err = h.sosmedUseCase.CreateSosmedUC(sosmed)
	if err != nil {
		return helpers.ErrorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, helpers.WebResponse{
		Status: "CREATED",
		Code:   201,
		Data:   sosmed,
	})
}
