package handler

import (
	todo "github.com/POMBNK/restAPI"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type allListsResponse struct {
	Data []todo.TodoList `json:"data"`
}
type deleteListResponse struct {
	Status string `json:"status"`
}

type updateListResponse struct {
	Status string `json:"status"`
}

// @Summary Get All Lists
// @Security ApiKeyAuth
// @Tags lists
// @Description Get all lists
// @ID get-all-lists
// @Accept  json
// @Produce  json
// @Success 200 {object} allListsResponse
// @Failure 400,404 {object} e
// @Failure 500 {object} e
// @Failure default {object} e
// @Router /api/lists [get]
func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, allListsResponse{
		Data: lists,
	})
}

// @Summary Get List By ID
// @Security ApiKeyAuth
// @Tags lists
// @Description Get list by ID
// @ID get-list-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} allListsResponse
// @Failure 400,404 {object} e
// @Failure 500 {object} e
// @Failure default {object} e
// @Router /api/lists/:id [get]
func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Bad id")
	}
	list, err := h.services.TodoList.GetByID(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

// @Summary Create Todo List
// @Security ApiKeyAuth
// @Tags lists
// @Description Create todo list
// @ID create-todo-list
// @Accept  json
// @Produce  json
// @Param input body todo.TodoList true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} e
// @Failure 500 {object} e
// @Failure default {object} e
// @Router /api/lists [post]
func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var list todo.TodoList
	if err = c.BindJSON(&list); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userId, list)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Update Todo List
// @Security ApiKeyAuth
// @Tags lists
// @Description Update todo list
// @ID update-todo-list
// @Accept  json
// @Produce  json
// @Param input body todo.UpdateList true "updated list info"
// @Success 200 {object} updateListResponse
// @Failure 400,404 {object} e
// @Failure 500 {object} e
// @Failure default {object} e
// @Router /api/lists/:id [put]
func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	//TODO: Should check if id exists in db already. If it not - abort.
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Bad id")
	}

	var inputUpdate todo.UpdateList
	if err = c.BindJSON(&inputUpdate); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.TodoList.Update(userId, listId, inputUpdate)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, updateListResponse{
		Status: "ok",
	})
}

// @Summary Delete Todo List
// @Security ApiKeyAuth
// @Tags lists
// @Description Delete todo list
// @ID delete-todo-list
// @Accept  json
// @Produce  json
// @Success 200 {object} deleteListResponse
// @Failure 400,404 {object} e
// @Failure 500 {object} e
// @Failure default {object} e
// @Router /api/lists/:id [delete]
func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Bad id")
	}

	err = h.services.TodoList.Delete(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusNoContent, err.Error())
		return
	}
	c.JSON(http.StatusOK, deleteListResponse{
		Status: "ok",
	})
}
