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

func GetTags(c *gin.Context) {
	appG := app.Gin{C: c}

	pageNum := util.GetPage(c)
	pageSize := setting.AppSetting.PageSize

	tags, err := models.GetTags(pageNum, pageSize)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_TAGS_FAIL, nil)
		return
	}

	serializer := serializers.TagsSerializer{c, tags}
	appG.Response(http.StatusOK, e.SUCCESS, serializer.Response())
}

func GetTag(c *gin.Context) {
	appG := app.Gin{C: c}
	id := c.Param("id")
	tag, err := models.GetTag(id)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_TAG_FAIL, nil)
		return
	}

	serializer := serializers.TagSerializer{c, tag}
	appG.Response(http.StatusOK, e.SUCCESS, serializer.Response())
}

type TagPrams struct {
	Name  string `form:"name" json:"name" binding:"required"`
	State int    `form:"name" json:"state" binding:"gte=0,lte=1"`
}

func AddTag(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form TagPrams
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	tag, err := models.AddTag(form.Name, form.State)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	serializer := serializers.TagSerializer{c, tag}
	appG.Response(http.StatusOK, e.SUCCESS, serializer.Response())
}

func EditTag(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form TagPrams
	)
	id := c.Param("id")

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	tagParams := map[string]interface{}{
		"name":  form.Name,
		"state": form.State,
	}

	tag, err := models.EditTag(id, tagParams)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_TAG_FAIL, nil)
		return
	}

	serializer := serializers.TagSerializer{c, tag}
	appG.Response(http.StatusOK, e.SUCCESS, serializer.Response())
}

func DeleteTag(c *gin.Context) {
	appG := app.Gin{C: c}
	id := c.Param("id")
	err := models.DeleteTag(id)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_TAG_FAIL, nil)
		return
	}
	appG.Response(http.StatusNoContent, e.SUCCESS, nil)
}
