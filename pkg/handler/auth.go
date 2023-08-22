package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/v1lezz/todo"
	"net/http"
)

func (h *Handler) singUp(c *gin.Context) {
	inp := todo.UserEntity{}
	if err := c.BindJSON(&inp); err != nil {
		newErrResp(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(inp)
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) singIn(c *gin.Context) {

}
