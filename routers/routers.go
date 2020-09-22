package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/huynhdev/go-gin-structure/middleware/jwt"
	"github.com/huynhdev/go-gin-structure/routers/api"
	v1 "github.com/huynhdev/go-gin-structure/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	r.POST("/login", api.Login)
	r.POST("/logout", api.Logout)
	r.POST("/token/refresh", api.Refresh)
	apiv1 := r.Group("/api/v1")
	apiv1.GET("/articles", v1.GetArticles)
	apiv1.GET("/articles/:id", v1.GetArticle)
	apiv1.GET("/tags", v1.GetTags)
	apiv1.GET("/tags/:id", v1.GetTag)
	apiv1.POST("/registrations", v1.UserRegistration)
	apiv1.Use(jwt.JWT())
	{
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}
