package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xopxe23/articles/internal/domain"
)

type ArticlesService interface {
	GetAll() ([]domain.ArticleOutput, error)
}

func (h *Handler) getAllArticles(c *gin.Context) {
	articles, err := h.articlesService.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"articles": articles,
	})
}

func (h *Handler) createArticle(c *gin.Context) {}

func (h *Handler) getArticleById(c *gin.Context) {}

func (h *Handler) updateArticle(c *gin.Context) {}

func (h *Handler) deleteArticle(c *gin.Context) {}
