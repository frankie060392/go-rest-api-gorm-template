package handlers

import (
	"frankie060392/rest-api-clean-arch/http/messages"
	"frankie060392/rest-api-clean-arch/internal/book/model"
	"frankie060392/rest-api-clean-arch/internal/common"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	bookService model.BookServiceInterface
}

func NewBookHandler(bookService model.BookRepositoryInterface) bookHandler {
	return bookHandler{bookService: bookService}
}

func (bh *bookHandler) GetByID(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	result, err := bh.bookService.GetById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseData{Status: false, Message: messages.ErrorCreate})
		return
	}
	ctx.JSON(http.StatusAccepted, result)
}

func (bh *bookHandler) Create(ctx *gin.Context) {
	var payload *model.BookCreate

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseData{Status: false, Message: messages.ErrorCreate})
		return
	}

	now := time.Now()
	newBook := model.Book{
		Name:      payload.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := bh.bookService.Create(ctx, &newBook)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseData{Status: false, Message: messages.ErrorCreate})
		return
	}

	bookResponse := &model.BookResponse{
		Name:      newBook.Name,
		ID:        newBook.ID,
		CreatedAt: newBook.CreatedAt,
		UpdatedAt: newBook.UpdatedAt,
	}

	ctx.JSON(http.StatusCreated, common.ResponseData{Status: true, Data: bookResponse, Message: messages.CreateSuccess})
}
