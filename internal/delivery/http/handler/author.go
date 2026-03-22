package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imanudd/inventorySvc-clean-architecture/constant"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/delivery/http/handler/dto"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/delivery/http/helper"
)

func (h *Handler) CreateAuthorAndBook(c *gin.Context) {
	var reqDTO dto.CreateAuthorAndBookRequest

	req, err := reqDTO.TransformIn(c)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, constant.ErrBadRequest)
		return
	}

	err = h.usecase.GetAuthorUseCase().CreateAuthorAndBook(c, req)
	if err != nil {
		helper.InternalError(c, err)
		return
	}

	helper.Success(c, http.StatusCreated)

}

// CreateAuthor handler
// @Summary create new author
// @Description create new author
// @Tags author
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param input body dto.CreateAuthorRequest true "data"
// @Success 200 {object} helper.JSONResponse
// @Failure 400 {object} helper.JSONResponse
// @Failure 500 {object} helper.JSONResponse
// @Router /inventorysvc/managements/author [POST]
func (h *Handler) CreateAuthor(c *gin.Context) {
	var reqDTO dto.CreateAuthorRequest

	req, err := reqDTO.TransformIn(c)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, constant.ErrBadRequest)
		return
	}

	err = h.usecase.GetAuthorUseCase().CreateAuthor(c, req)
	if err != nil {
		helper.InternalError(c, err)
		return
	}

	helper.Success(c, http.StatusCreated)

}

// AddAuthorBook handler
// @Summary add author book
// @Description add author book
// @Tags author
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "author id"
// @Param input body dto.AddAuthorBookRequest true "data"
// @Success 200 {object} helper.JSONResponse
// @Failure 400 {object} helper.JSONResponse
// @Failure 500 {object} helper.JSONResponse
// @Router /inventorysvc/managements/author/{id} [POST]
func (h *Handler) AddAuthorBook(c *gin.Context) {

	var reqDTO dto.AddAuthorBookRequest

	req, err := reqDTO.TransformIn(c)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, constant.ErrBadRequest)
		return
	}

	err = h.usecase.GetAuthorUseCase().AddAuthorBook(c, req)
	if err != nil {
		helper.InternalError(c, err)
		return
	}

	helper.Success(c, http.StatusCreated)
}

// GetListBookByAuthor handler
// @Summary get list book by author
// @Description get list book by author
// @Tags author
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "author id"
// @Success 200 {object} helper.JSONResponse{data=[]dto.BookResponse}
// @Failure 400 {object} helper.JSONResponse
// @Failure 500 {object} helper.JSONResponse
// @Router /inventorysvc/managements/author/{id}/list [GET]
func (h *Handler) GetListBookByAuthor(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.Error(c, http.StatusBadRequest, constant.ErrBadRequest)
		return
	}

	books, err := h.usecase.GetAuthorUseCase().GetListBookByAuthor(c, int64(id))
	if err != nil {
		helper.InternalError(c, err)
		return
	}

	var respDTOs []dto.BookResponse
	for _, book := range books {
		var bookDTO dto.BookResponse
		bookDTO.TransformOut(book)
		respDTOs = append(respDTOs, bookDTO)
	}

	helper.Success(c, http.StatusOK, respDTOs)
}

// GetListBookByAuthor handler
// @Summary get list book by author
// @Description get list book by author
// @Tags author
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "author id"
// @Param bookid path string true "book id"
// @Success 200 {object} helper.JSONResponse
// @Failure 400 {object} helper.JSONResponse
// @Failure 500 {object} helper.JSONResponse
// @Router /inventorysvc/managements/author/{id}/books/{bookid} [DELETE]
func (h *Handler) DeleteBookByAuthor(c *gin.Context) {
	bookId := c.GetInt64("bookid")
	id := c.GetInt64("id")

	if bookId == 0 || id == 0 {
		helper.Error(c, http.StatusBadRequest, constant.ErrBadRequest)
		return
	}

	err := h.usecase.GetAuthorUseCase().DeleteBookByAuthor(c, id, bookId)
	if err != nil {
		helper.InternalError(c, err)
		return
	}

	helper.Success(c, http.StatusOK)
}
