package http

import (
	"github.com/asaskevich/govalidator"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"uji/domain"
	"uji/helpers"
	"uji/middlewares"
)

type SosmedHandler struct {
	sosmedUseCase domain.SosmedUseCase
}

func NewSosmedHandler(e *echo.Echo, sosmedUc domain.SosmedUseCase, db *gorm.DB) {
	handler := SosmedHandler{
		sosmedUseCase: sosmedUc,
	}
	router := e.Group("/sosmed")
	{
		router.Use(middlewares.Authentication)
		router.POST("/", handler.CreateSosmed)
		router.GET("/", handler.GetSosmeds)
		router.Use(middlewares.SosmedAuthorization(db))
		router.PUT("/:Id", handler.UpdateSosmed)
		router.DELETE("/:Id", handler.DeleteUser)

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

func (h *SosmedHandler) GetSosmeds(ctx echo.Context) error {
	var sosmed []*domain.SocialMedia

	res, err := h.sosmedUseCase.GetSosmedsUC(sosmed)
	if err != nil {
		return helpers.ErrorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusOK, helpers.WebResponse{
		Status: "OK",
		Code:   200,
		Data:   res,
	})
}

func (h *SosmedHandler) UpdateSosmed(ctx echo.Context) error {
	newSosmed := new(domain.SocialMedia)
	id, _ := strconv.Atoi(ctx.Param("Id"))

	if err := ctx.Bind(newSosmed); err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.WebResponse{
			Status: "BAD_REQUEST",
			Code:   400,
			Data:   "Json Payload Invalid!",
		})
	}

	//_, err := govalidator.ValidateStruct(newSosmed)
	//if err != nil {
	//	return ctx.JSON(http.StatusBadRequest, helpers.WebResponse{
	//		Status: "BAD_REQUEST",
	//		Code:   400,
	//		Data:   "Validation Error!",
	//	})
	//}

	res, err := h.sosmedUseCase.UpdateSosmedUC(uint(id), newSosmed)
	if err != nil {
		return helpers.ErrorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusOK, helpers.WebResponse{
		Status: "OK",
		Code:   200,
		Data:   res,
	})
}

func (h *SosmedHandler) DeleteUser(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("Id"))

	err := h.sosmedUseCase.DeleteSosmedUC(uint(id))
	if err != nil {
		return helpers.ErrorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusOK, helpers.WebResponse{
		Status: "OK",
		Code:   200,
		Data:   "Succesfull Deleted!",
	})
}
