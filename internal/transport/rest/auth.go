package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xopxe23/articles/internal/domain"
)

type AuthService interface {
	SignUp(input domain.User) error
	SignIn(input domain.SignInInput) (string, string, error)
	RefreshTokens(token string) (string, string, error)
	ParseToken(token string) (int, error)
}

// @Summary Sign Up
// @Tags Users auth
// @ID sign-up
// @Accept json
// @Produce json
// @Param input body domain.User true "Sign up input"
// @Success 200
// @Failure 400
// @Router /auth/sign-up [post]
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

// @Summary Sign In
// @Tags Users auth
// @ID sign-in
// @Accept json
// @Produce json
// @Param input body domain.SignInInput true "Sign in input"
// @Success 200 {string} string
// @Failure 400
// @Failure 500
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input domain.SignInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, refreshToken, err := h.authService.SignIn(input)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Header("Set-Cookie", fmt.Sprintf("refresh-token=%s; HttpOnly", refreshToken))
	c.JSON(http.StatusOK, map[string]string{
		"token": accessToken,
	})
}

// @Summary Refresh
// @Tags Users auth
// @ID refresh
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Failure 500
// @Router /auth/refresh [get]
func (h *Handler) refresh(c *gin.Context) {
	cookie, err := c.Cookie("refresh-token")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	accessToken, refreshToken, err := h.authService.RefreshTokens(cookie)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Set-Cookie", fmt.Sprintf("refresh-token=%s; HttpOnly", refreshToken))
	c.JSON(http.StatusOK, map[string]string{
		"token": accessToken,
	})
}
