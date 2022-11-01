package routers

import (
	v1 "myblog/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("api/v1")
	{
		apiv1.GET("hello", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		})

		// 用户模块的路由接口
		apiv1.POST("user/add", v1.AddUser)
		apiv1.GET("users", v1.GetUsers)
		apiv1.PUT("user/:id", v1.EditUser)
		apiv1.DELETE("user/:id", v1.DeleteUser)

		// 标签模块的路由接口
		apiv1.POST("tag/add", v1.AddTag)
		apiv1.GET("tags", v1.GetTags)
		apiv1.PUT("tag/:id", v1.EditTag)
		apiv1.DELETE("tag/:id", v1.DeleteTag)

		// 文章模块的路由接口
		apiv1.POST("article/add", v1.AddArticle)
		apiv1.GET("articles", v1.GetArticles)
		apiv1.PUT("article/:id", v1.EditArticle)
		apiv1.DELETE("article/:id", v1.DeleteArticle)

	}

	return r
}
