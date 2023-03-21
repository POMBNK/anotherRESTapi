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

// @Summary Get All Items
// @Security ApiKeyAuth
// @Tags items
// @Description Get all items
// @ID get-all-items
// @Accept  json
// @Produce  json
// @Success 200 {object} []todo.TodoItem
// @Failure 400,404 {object} e
// @Failure 500 {object} e
// @Failure default {object} e
// @Router /api/lists/:id/items [get]
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

// @Summary Get Item By ID
// @Security ApiKeyAuth
// @Tags items
// @Description Get item by id
// @ID get-item-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} todo.TodoItem
// @Failure 400,404 {object} e
// @Failure 500 {object} e
// @Failure default {object} e
// @Router /api/items/:id [get]
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

// @Summary Create Todo Item
// @Security ApiKeyAuth
// @Tags items
// @Description Create todo item
// @ID create-todo-item
// @Accept  json
// @Produce  json
// @Param input body todo.TodoItem true "item info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} e
// @Failure 500 {object} e
// @Failure default {object} e
// @Router /api/lists/:id/items [post]
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

// @Summary Update Item
// @Security ApiKeyAuth
// @Tags items
// @Description Update item by id
// @ID update-item-by-id
// @Accept  json
// @Produce  json
// @Param input body todo.UpdateItem true "item info"
// @Success 200 {object} updateItemResponse
// @Failure 400,404 {object} e
// @Failure 500 {object} e
// @Failure default {object} e
// @Router /api/items/:id [put]
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

// @Summary Delete Item
// @Security ApiKeyAuth
// @Tags items
// @Description Delete item by id
// @ID Delete-item-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} deleteItemResponse
// @Failure 400,404 {object} e
// @Failure 500 {object} e
// @Failure default {object} e
// @Router /api/items/:id [delete]
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
