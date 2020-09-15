package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huynhdev/go-gin-structure/models"
	"github.com/huynhdev/go-gin-structure/pkg/app"
	"github.com/huynhdev/go-gin-structure/pkg/e"
	"github.com/huynhdev/go-gin-structure/pkg/setting"
	"github.com/huynhdev/go-gin-structure/serializers"
	"github.com/huynhdev/go-gin-structure/util"
)

func GetArticles(c *gin.Context) {
	appG := app.Gin{C: c}
	pageNum := util.GetPage(c)
	pageSize := setting.AppSetting.PageSize
	articles, err := models.GetArticles(pageNum, pageSize)
	if err != nil {
		appG.Response(http.StatusOK, e.RECORD_NOT_FOUND, nil)
		return
	}

	serializer := serializers.ArticlesSerializer{c, articles}
	appG.Response(http.StatusOK, e.SUCCESS, serializer.Response())
}

func GetArticle(c *gin.Context) {
	appG := app.Gin{C: c}
	id := c.Param("id")
	article, err := models.GetArticle(id)
	if err != nil {
		appG.Response(http.StatusOK, e.RECORD_NOT_FOUND, nil)
		return
	}

	serializer := serializers.ArticleSerializer{c, article}
	appG.Response(http.StatusOK, e.SUCCESS, serializer.Response())
}

type ArticleParams struct {
	TagID         int    `json:"tag_id" binding:"required"`
	Title         string `json:"title" binding:"required"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	State         int    `json:"state" binding:"gte=0,lte=1"`
	CoverImageUrl string `json:"cover_image_url"`
}

func AddArticle(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form ArticleParams
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	articleParams := map[string]interface{}{
		"tag_id":          form.TagID,
		"title":           form.Title,
		"desc":            form.Desc,
		"content":         form.Content,
		"cover_image_url": form.CoverImageUrl,
		"state":           form.State,
	}
	article, err := models.AddArticle(articleParams)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_ARTICLE_FAIL, nil)
		return
	}

	serializer := serializers.ArticleSerializer{c, article}
	appG.Response(http.StatusCreated, e.SUCCESS, serializer.Response())
}

func EditArticle(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form ArticleParams
	)

	id := c.Param("id")

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	articleParams := map[string]interface{}{
		"tag_id":          form.TagID,
		"title":           form.Title,
		"desc":            form.Desc,
		"content":         form.Content,
		"cover_image_url": form.CoverImageUrl,
		"state":           form.State,
	}
	article, err := models.EditArticle(id, articleParams)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_ARTICLE_FAIL, nil)
		return
	}

	serializer := serializers.ArticleSerializer{c, article}
	appG.Response(http.StatusCreated, e.SUCCESS, serializer.Response())
}

func DeleteArticle(c *gin.Context) {
	var appG = app.Gin{C: c}
	id := c.Param("id")

	err := models.DeleteArticle(id)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_ARTICLE_FAIL, nil)
		return
	}
	appG.Response(http.StatusNoContent, e.SUCCESS, nil)

}
