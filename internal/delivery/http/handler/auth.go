package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imanudd/inventorySvc-clean-architecture/constant"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/delivery/http/handler/dto"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/delivery/http/helper"
)

// Login handler
// @Summary login user
// @Description login user
// @Tags auth
// @Accept json
// @Produce json
// @Param input body dto.LoginRequest true "login data"
// @Success 200 {object} helper.JSONResponse
// @Failure 400 {object} helper.JSONResponse
// @Failure 500 {object} helper.JSONResponse
// @Router /inventorysvc/auth/login [POST]
func (h *Handler) Login(c *gin.Context) {
	var reqDTO dto.LoginRequest

	req, err := reqDTO.TransformIn(c)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, constant.ErrBadRequest)
		return
	}

	resp, err := h.usecase.GetAuthUseCase().Login(c, req)
	if err != nil {
		helper.InternalError(c, err)
		return
	}

	var respDTO dto.LoginResponse
	respDTO.TransformOut(resp)

	helper.Success(c, http.StatusOK, respDTO)
}

// Register handler
// @Summary register user
// @Description register user
// @Tags auth
// @Accept json
// @Produce json
// @Param input body dto.RegisterRequest true "register data"
// @Success 200 {object} helper.JSONResponse
// @Failure 400 {object} helper.JSONResponse
// @Failure 500 {object} helper.JSONResponse
// @Router /inventorysvc/auth/register [POST]
func (h *Handler) Register(c *gin.Context) {
	var reqDTO dto.RegisterRequest

	req, err := reqDTO.TransformIn(c)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, constant.ErrBadRequest)
		return
	}

	err = h.usecase.GetAuthUseCase().Register(c, req)
	if err != nil {
		helper.InternalError(c, err)
		return
	}

	helper.Success(c, http.StatusOK)
}
