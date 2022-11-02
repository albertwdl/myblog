package v1

import (
	"myblog/middleware"
	"myblog/model"
	"myblog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user model.User
	c.ShouldBindJSON(&user)

	code := model.CheckLogin(user.Username, user.Password)
	var token string
	if code == errmsg.SUCCESS {
		var tokenCode int
		token, tokenCode = middleware.SetToken(user.Username)
		if tokenCode == errmsg.ERROR {
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
				"token":   token,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})

}
