package rest

import "github.com/gin-gonic/gin"

type Handler struct {
	authService AuthService
}

func NewHandler(authService AuthService) *Handler {
	return &Handler{authService: authService}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/refresh", h.refresh)
	}
	articles := router.Group("/articles")
	{
		articles.GET("", h.getAllArticles)
		articles.POST("", h.createArticle)
		articles.GET("/:id", h.getArticleById)
		articles.PUT("/:id", h.updateArticle)
		articles.DELETE("/:id", h.deleteArticle)
	}
	return router
}
