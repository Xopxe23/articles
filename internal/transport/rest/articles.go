package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xopxe23/articles/internal/domain"
)

type ArticlesService interface {
	GetAll() ([]domain.ArticleOutput, error)
	Create(input domain.ArticleInput, userId int) error
	GetById(id int) (domain.ArticleOutput, error)
	Delete(id, userId int) error
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

func (h *Handler) createArticle(c *gin.Context) {
	userId := c.GetInt(userCtx)
	var input domain.ArticleInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.articlesService.Create(input, userId); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"status": "article created",
	})
}

func (h *Handler) getArticleById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	article, err := h.articlesService.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"article": article,
	})
}

func (h *Handler) updateArticle(c *gin.Context) {}

func (h *Handler) deleteArticle(c *gin.Context) {
	userId := c.GetInt(userCtx)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.articlesService.Delete(id, userId); err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"status": "artcile deleted",
	})
}
