package http

import (
	"github.com/labstack/echo/v4"
	"uji/domain"
)

type CommentHandler struct {
	commentUC domain.CommentUseCase
}

func NewCommentUseCase(e *echo.Echo, commentUC domain.CommentUseCase) {
	handler := CommentHandler{
		commentUC,
	}
	router := e.Group("/comment")
	{
		router.POST("/", handler.CreateComment)
	}
}

func (h *CommentHandler) CreateComment(ctx echo.Context) error {

}