package routers

import (
	v1 "myblog/api/v1"
	"myblog/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	r.Use(middleware.Cors())

	authapiv1 := r.Group("api/v1")
	authapiv1.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		authapiv1.GET("users", v1.GetUsers)
		authapiv1.PUT("user/:id", v1.EditUser)
		authapiv1.DELETE("user/:id", v1.DeleteUser)

		// 标签模块的路由接口
		authapiv1.POST("tag/add", v1.AddTag)
		authapiv1.PUT("tag/:id", v1.EditTag)
		authapiv1.DELETE("tag/:id", v1.DeleteTag)

		// 文章模块的路由接口
		authapiv1.POST("article/add", v1.AddArticle)
		authapiv1.PUT("article/:id", v1.EditArticle)
		authapiv1.DELETE("article/:id", v1.DeleteArticle)

		// 上传文件
		authapiv1.POST("upload", v1.UpLoad)

	}

	apiv1 := r.Group("api/v1")
	{
		apiv1.POST("user/add", v1.AddUser)
		apiv1.GET("tags", v1.GetTags)
		apiv1.GET("articles", v1.GetArticles)
		apiv1.GET("articlesbytag/:id", v1.GetArticlesByTag)
		apiv1.GET("article/info/:id", v1.GetArticleInfo)
		apiv1.POST("login", v1.Login)
	}

	return r
}
