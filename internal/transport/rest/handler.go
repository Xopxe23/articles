package rest

import (
	_ "github.com/xopxe23/articles/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Handler struct {
	authService     AuthService
	articlesService ArticlesService
}

func NewHandler(authService AuthService, articlesService ArticlesService) *Handler {
	return &Handler{authService: authService, articlesService: articlesService}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swag/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/refresh", h.refresh)
	}
	articles := router.Group("/articles")
	articles.Use(h.userIdentity)
	{
		articles.GET("", h.getAllArticles)
		articles.POST("", h.createArticle)
		articles.GET("/:id", h.getArticleById)
		articles.PUT("/:id", h.updateArticle)
		articles.DELETE("/:id", h.deleteArticle)
	}
	return router
}
