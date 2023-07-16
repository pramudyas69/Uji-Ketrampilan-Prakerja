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

type CommentHandler struct {
	commentUC domain.CommentUseCase
}

func NewCommentUseCase(e *echo.Echo, commentUC domain.CommentUseCase, db *gorm.DB) {
	handler := CommentHandler{
		commentUC,
	}
	router := e.Group("/comment")
	{
		router.Use(middlewares.Authentication)
		router.POST("/", handler.CreateComment)
		router.GET("/", handler.GetComments)
		router.PUT("/:Id", handler.UpdateComment, middlewares.CommentAuthorization(db))
		router.DELETE("/:Id", handler.DeleteComment, middlewares.CommentAuthorization(db))
	}
}

func (h *CommentHandler) CreateComment(ctx echo.Context) error {
	comment := new(domain.Comment)

	if err := ctx.Bind(&comment); err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.WebResponse{
			Status: "BAD_REQUEST",
			Code:   400,
			Data:   "Json Payload Invalid!",
		})
	}
	_, err := govalidator.ValidateStruct(comment)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.WebResponse{
			Status: "BAD_REQUEST",
			Code:   400,
			Data:   "Validation Error!",
		})
	}

	userData := ctx.Get("userData").(jwt.MapClaims)
	userID := uint32(userData["id"].(float64))

	comment.UserID = userID
	err = h.commentUC.CreateCommentUC(comment)
	if err != nil {
		return helpers.ErrorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, helpers.WebResponse{
		Status: "CREATED",
		Code:   201,
		Data:   comment,
	})
}

func (h *CommentHandler) GetComments(ctx echo.Context) error {
	var comment *[]domain.Comment

	res, err := h.commentUC.GetCommentsUC(comment)
	if err != nil {
		return helpers.ErrorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusOK, helpers.WebResponse{
		Status: "OK",
		Code:   200,
		Data:   res,
	})
}

func (h *CommentHandler) UpdateComment(ctx echo.Context) error {
	newComment := new(domain.Comment)
	id, _ := strconv.Atoi(ctx.Param("Id"))

	if err := ctx.Bind(newComment); err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.WebResponse{
			Status: "BAD_REQUEST",
			Code:   400,
			Data:   "Json Payload Invalid!",
		})
	}

	_, err := govalidator.ValidateStruct(newComment)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.WebResponse{
			Status: "BAD_REQUEST",
			Code:   400,
			Data:   "Validation Error!",
		})
	}

	res, err := h.commentUC.UpdateCommentUC(uint(id), newComment)
	if err != nil {
		return helpers.ErrorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusOK, helpers.WebResponse{
		Status: "OK",
		Code:   200,
		Data:   res,
	})
}

func (h *CommentHandler) DeleteComment(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("Id"))

	err := h.commentUC.DeleteCommentUC(uint(id))
	if err != nil {
		return helpers.ErrorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusOK, helpers.WebResponse{
		Status: "OK",
		Code:   200,
		Data:   "Succesfull Deleted!",
	})
}
