package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huynhdev/go-gin-structure/models"
	"github.com/huynhdev/go-gin-structure/pkg/app"
	"github.com/huynhdev/go-gin-structure/pkg/e"
	"github.com/huynhdev/go-gin-structure/util"
)

type UserParams struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

func UserRegistration(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form UserParams
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	userExists := models.CheckUserExists(form.Username)
	if userExists {
		appG.Response(http.StatusForbidden, e.ERROR_USERNAME_EXISTS, nil)
		return
	}

	userParams := map[string]interface{}{
		"username": form.Username,
		"password": util.GetHash([]byte(form.Password)),
	}

	_, err := models.CreateUser(userParams)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_REGISTRATION_FAIL, nil)
		return
	}

	appG.Response(http.StatusCreated, e.SUCCESS, nil)

}
