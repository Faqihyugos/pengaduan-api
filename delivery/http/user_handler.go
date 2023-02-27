package http

import (
	"net/http"

	entityuser "github.com/faqihyugos/pengaduan-api/entities/entityUser"
	"github.com/faqihyugos/pengaduan-api/helper"
	"github.com/faqihyugos/pengaduan-api/services/serviceuser"
	"github.com/gin-gonic/gin"
)

type httpUser struct {
	srv serviceuser.ServiceUser
}

func NewUser(srv serviceuser.ServiceUser) DeliveryUser {
	return &httpUser{srv}
}

type DeliveryUser interface {
	Create(ctx *gin.Context)
	Login(ctx *gin.Context)
	Update(ctx *gin.Context)
	DeleteByID(ctx *gin.Context)
}

// create user
func (h *httpUser) Create(ctx *gin.Context) {

	contentType := helper.GetContentType(ctx)
	data := new(entityuser.Request)

	// check content type
	if contentType == "multipart/form-data" {
		// get form data
		if err := ctx.ShouldBind(data); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
			return
		}
	} else {
		// get json data
		if err := ctx.ShouldBindJSON(data); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
			return
		}
	}

	// // get form data or json data

	// if err := ctx.ShouldBindJSON(data); err != nil {
	// 	ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
	// 	return
	// }

	response, err := h.srv.Create(*data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err))
		return
	}

	ctx.JSON(http.StatusCreated, helper.NewResponse(http.StatusCreated, response, nil))

}

// login user
func (h *httpUser) Login(ctx *gin.Context) {
	data := new(entityuser.RequestLogin)

	if err := ctx.ShouldBindJSON(data); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	response, err := h.srv.Login(*data)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.NewResponse(http.StatusInternalServerError, nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, nil))
}

// update user
func (h *httpUser) Update(ctx *gin.Context) {
	data := new(entityuser.Request)

	if err := ctx.ShouldBindJSON(data); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	response, err := h.srv.Update(*data)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.NewResponse(http.StatusInternalServerError, nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, nil))
}

// delete user
func (h *httpUser) DeleteByID(ctx *gin.Context) {
	data := new(entityuser.Request)

	if err := ctx.ShouldBindJSON(data); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	err := h.srv.DeleteByID(data.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.NewResponse(http.StatusInternalServerError, nil, err))
		return
	}
}
