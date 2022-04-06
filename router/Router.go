package router

import (
	"example/sample/service"
	"github.com/gin-gonic/gin"
)

func AdminRoute(group *gin.RouterGroup) {
	group.GET("home", service.HomePage)
	group.GET("sample", service.ReturnAllArticles)
	group.GET("articles", service.ArticleList)
	group.POST("articles", service.CreateArticle)
	group.PUT("articles", service.UpdateArticle)
	group.DELETE("articles", service.DeleteArticle)
}
