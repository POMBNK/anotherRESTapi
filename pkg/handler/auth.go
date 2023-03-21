package handler

import (
	"github.com/POMBNK/restAPI"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Sign-Up
// @tags auth
// @Description Create new account
// @ID create-account
// @Accept json
// @Produce  json
// @Param input body todo.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} e
// @Failure 500 {object} e
// @Failure default {object} e
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary Sign-In
// @tags auth
// @Description Login in account
// @ID Login-in-account
// @Accept json
// @Produce  json
// @Param input body signInInput true "credentials"
// @Success 200 {integer} string "token"
// @Failure 400,404 {object} e
// @Failure 500 {object} e
// @Failure default {object} e
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
