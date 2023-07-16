package http

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"uji/domain"
	"uji/helpers"
	"uji/helpers/other_helpers"
	"uji/middlewares"
)

type PhotoHandler struct {
	photoUseCase domain.PhotoUseCase
}

func NewPhotoHandler(e *echo.Echo, photoUseCase domain.PhotoUseCase, db *gorm.DB) {
	handler := PhotoHandler{
		photoUseCase: photoUseCase,
	}
	router := e.Group("/photo")
	{
		router.Use(middlewares.Authentication)
		router.POST("/", handler.CreatePhoto)
		router.GET("/", handler.GetPhotos)
		router.PUT("/:Id", handler.UpdatePhoto, middlewares.PhotoAuthorization(db))
		router.DELETE("/:Id", handler.DeletePhoto, middlewares.PhotoAuthorization(db))
	}
}

func (h *PhotoHandler) CreatePhoto(ctx echo.Context) error {
	photo := new(domain.Photo)

	if err := ctx.Bind(&photo); err != nil {
		fmt.Println(photo)
		return ctx.JSON(http.StatusBadRequest, helpers.WebResponse{
			Status: "BAD_REQUEST",
			Code:   400,
			Data:   "Json Payload Invalid!",
		})
	}

	_, err := govalidator.ValidateStruct(photo)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.WebResponse{
			Status: "BAD_REQUEST",
			Code:   400,
			Data:   "Validation Error!",
		})
	}

	isUrl := govalidator.IsURL(photo.PhotoURL)
	if !isUrl {
		return ctx.JSON(http.StatusBadRequest, helpers.WebResponse{
			Status: "BAD_REQUEST",
			Code:   400,
			Data:   "photo_url is not url!",
		})
	}

	userData := ctx.Get("userData").(jwt.MapClaims)
	userID := uint32(userData["id"].(float64))

	photo.UserID = userID
	err = h.photoUseCase.CreatePhotoUC(photo)

	if err != nil {
		return helpers.ErrorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, helpers.WebResponse{
		Status: "CREATED",
		Code:   201,
		Data:   photo,
	})

}

func (h *PhotoHandler) GetPhotos(ctx echo.Context) error {
	var photo *[]domain.Photo

	res, err := h.photoUseCase.GetPhotosUC(photo)
	if err != nil {
		return helpers.ErrorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusOK, helpers.WebResponse{
		Status: "OK",
		Code:   200,
		Data:   res,
	})
}

func (h *PhotoHandler) UpdatePhoto(ctx echo.Context) error {
	newPhoto := new(domain.PhotoUpdateInput)
	id, _ := strconv.Atoi(ctx.Param("Id"))

	if err := ctx.Bind(newPhoto); err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.WebResponse{
			Status: "BAD_REQUEST",
			Code:   400,
			Data:   "Json Payload Invalid!",
		})
	}

	_, err := govalidator.ValidateStruct(newPhoto)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.WebResponse{
			Status: "BAD_REQUEST",
			Code:   400,
			Data:   "Validation Error!",
		})
	}

	photo := other_helpers.CopyStructPhoto(newPhoto)

	res, err := h.photoUseCase.UpdatePhotoUC(uint(id), photo)
	if err != nil {
		return helpers.ErrorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusOK, helpers.WebResponse{
		Status: "OK",
		Code:   200,
		Data:   res,
	})
}

func (h *PhotoHandler) DeletePhoto(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("Id"))

	err := h.photoUseCase.DeletePhotoUC(uint(id))
	if err != nil {
		return helpers.ErrorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusOK, helpers.WebResponse{
		Status: "OK",
		Code:   200,
		Data:   "Succesfull Deleted!",
	})
}
