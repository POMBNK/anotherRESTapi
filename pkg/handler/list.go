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
