package handler

import (
	todo "github.com/POMBNK/restAPI"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type deleteItemResponse struct {
	Status string `json:"status"`
}

type updateItemResponse struct {
	Status string `json:"status"`
}

func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	//TODO: Should check too look at list.go
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Bad id")
	}

	items, err := h.services.TodoItem.GetAll(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	//TODO: Should check too look at list.go
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Bad id")
	}

	item, err := h.services.TodoItem.GetByID(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	//TODO: Should check too look at list.go
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Bad id")
	}

	var item todo.TodoItem
	if err = c.BindJSON(&item); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoItem.Create(userId, listId, item)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	//TODO: Should check if id exists in db already. If it not - abort.
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Bad id")
	}

	var inputUpdate todo.UpdateItem

	if err = c.BindJSON(&inputUpdate); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.TodoItem.Update(userId, itemId, inputUpdate)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, updateItemResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	//TODO: Should check too look at list.go
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Bad id")
	}

	err = h.services.TodoItem.Delete(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, deleteItemResponse{"ok"})
}
