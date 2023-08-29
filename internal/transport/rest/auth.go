package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xopxe23/articles/internal/domain"
)

type AuthService interface {
	SignUp(domain.User) error
}

func (h *Handler) signUp(c *gin.Context) {
	var input domain.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.authService.SignUp(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, map[string]string{
		"status": "registration completed",
	})
}

func (h *Handler) signIn(c *gin.Context) {}

func (h *Handler) refresh(c *gin.Context) {}
