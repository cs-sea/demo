package controller

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {string} http.code
// @Header 200 {string} Token "qwerty"
// @Failure 500 {string} error
// @Failure default {string} error
// @Router /accounts/{id} [get]
func (c *Controller) ShowAccount(ctx *gin.Context) {

}

// ShowAccount1 godoc
// @Summary 获取列表
// @Description get string by ID
// @ID get-string-by-list
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {string} http.code
// @Header 200 {string} Token "qwerty"
// @Failure 500 {string} error
// @Failure default {string} error
// @Router /account/list/s [get]
func (c *Controller) ShowAccount1(ctx *gin.Context) {

}
