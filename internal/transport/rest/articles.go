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
	Update(id, userId int, input domain.ArticleInput) error
	Delete(id, userId int) error
}

// @Summary Get All Articles
// @Security BearerAuth
// @Tags Articles
// @ID get-all-articles
// @Accept json
// @Produce json
// @Success 200 {array} domain.ArticleOutput
// @Failure 400
// @Failure 500
// @Router /articles [get]
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

// @Summary Create Article
// @Security BearerAuth
// @Tags Articles
// @ID create-articles
// @Accept json
// @Produce json
// @Param input body domain.ArticleInput true "Article input"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /articles [post]
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

// @Summary Get Article By Id
// @Security BearerAuth
// @Tags Articles
// @ID get-article-by-id
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Success 200 {object} domain.ArticleOutput
// @Failure 400
// @Failure 500
// @Router /articles/{id} [get]
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

// @Summary Update Article
// @Security BearerAuth
// @Tags Articles
// @ID update-article
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Param input body domain.ArticleInput true "Update Article input"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /articles/{id} [put]
func (h *Handler) updateArticle(c *gin.Context) {
	userId := c.GetInt(userCtx)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var input domain.ArticleInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.articlesService.Update(id, userId, input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"status": "article updated",
	})
}

// @Summary Delete Article
// @Security BearerAuth
// @Tags Articles
// @ID delete-article
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /articles/{id} [delete]
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
