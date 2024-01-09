package handlers

import (
	"fmt"
	"frankie060392/rest-api-clean-arch/internal/book/model"
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
	fmt.Println(id)
	result, err := bh.bookService.GetById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Cant not get")
		return
	}
	ctx.JSON(http.StatusAccepted, result)
}

func (bh *bookHandler) Create(ctx *gin.Context) {
	var payload *model.BookCreate

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, "Error")
		return
	}

	now := time.Now()
	newBook := model.Book{
		Name:      payload.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := bh.bookService.Create(ctx, &newBook)
	fmt.Println(newBook)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Can not create")
		return
	}

	bookResponse := &model.BookResponse{
		Name:      newBook.Name,
		ID:        newBook.ID,
		CreatedAt: newBook.CreatedAt,
		UpdatedAt: newBook.UpdatedAt,
	}

	ctx.JSON(http.StatusCreated, bookResponse)
}
