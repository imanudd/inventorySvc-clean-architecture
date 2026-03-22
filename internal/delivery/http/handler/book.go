package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imanudd/inventorySvc-clean-architecture/constant"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/delivery/http/handler/dto"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/delivery/http/helper"
)

// AddBook handler
// @Summary add book
// @Description add book
// @Tags book
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param input body dto.CreateBookRequest true "data"
// @Success 200 {object} helper.JSONResponse
// @Failure 400 {object} helper.JSONResponse
// @Failure 500 {object} helper.JSONResponse
// @Router /inventorysvc/managements/book [POST]
func (h *Handler) AddBook(c *gin.Context) {
	var reqDTO dto.CreateBookRequest

	req, err := reqDTO.TransformIn(c)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, constant.ErrBadRequest)
		return
	}

	err = h.usecase.GetBookUseCase().AddBook(c, req)
	if err != nil {
		helper.InternalError(c, err)
		return
	}

	helper.Success(c, http.StatusCreated)
}

// UpdateBook handler
// @Summary update book
// @Description update book
// @Tags book
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param id path string true "book id"
// @Param input body dto.UpdateBookRequest true "data"
// @Success 200 {object} helper.JSONResponse
// @Failure 400 {object} helper.JSONResponse
// @Failure 500 {object} helper.JSONResponse
// @Router /inventorysvc/managements/book/{id} [POST]
func (h *Handler) UpdateBook(c *gin.Context) {
	var reqDTO dto.UpdateBookRequest

	req, err := reqDTO.TransformIn(c)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, constant.ErrBadRequest)
		return
	}

	err = h.usecase.GetBookUseCase().UpdateBook(c, req)
	if err != nil {
		helper.InternalError(c, err)
		return
	}

	helper.Success(c, http.StatusOK)
}

// DeleteBook handler
// @Summary delete book
// @Description delete book
// @Tags book
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param id path string true "book id"
// @Success 200 {object} helper.JSONResponse
// @Failure 400 {object} helper.JSONResponse
// @Failure 500 {object} helper.JSONResponse
// @Router /inventorysvc/managements/book/{id} [DELETE]
func (h *Handler) DeleteBook(c *gin.Context) {

	id := c.GetInt64("id")
	if id == 0 {
		helper.Error(c, http.StatusBadRequest, constant.ErrBadRequest)
		return
	}

	err := h.usecase.GetBookUseCase().DeleteBook(c, id)
	if err != nil {
		helper.InternalError(c, err)
		return
	}

	helper.Success(c, http.StatusOK)
}

// GetDetailBook handler
// @Summary get detail book
// @Description get detail book
// @Tags book
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param id path string true "book id"
// @Success 200 {object} helper.JSONResponse{data=dto.DetailBookResponse}
// @Failure 400 {object} helper.JSONResponse
// @Failure 500 {object} helper.JSONResponse
// @Router /inventorysvc/managements/book/{id} [GET]
func (h *Handler) GetDetailBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.Error(c, http.StatusBadRequest, constant.ErrBadRequest)
		return
	}

	detail, err := h.usecase.GetBookUseCase().GetDetailBook(c, int64(id))
	if err != nil {
		helper.InternalError(c, err)
		return
	}

	var respDTO dto.DetailBookResponse
	respDTO.TransformOut(detail)

	helper.Success(c, http.StatusOK, respDTO)
}
