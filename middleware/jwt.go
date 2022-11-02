package middleware

import (
	"myblog/global"
	"myblog/model"
	"myblog/utils/errmsg"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 生成token
func SetToken(username string) (string, int) {
	var JwtKey = []byte(global.ServerSetting.JwtKey)
	expireTime := time.Now().Add(24 * time.Hour)
	claims := MyClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "albertsblog",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return ss, errmsg.SUCCESS
}

// 验证token
func CheckToken(tokenString string) (*MyClaims, int) {
	token, _ := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.ServerSetting.JwtKey), nil
	})

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, errmsg.SUCCESS
	} else {
		return nil, errmsg.ERROR
	}
}

// Jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		code := errmsg.SUCCESS
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_NOT_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		// 先判断是否符合"Bearer ******..."的格式
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_FORMAT_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		//再判断Token是否符合"Header.Claims.Signature"的格式
		tokenPart := strings.SplitN(checkToken[1], ".", 3)
		if len(tokenPart) != 3 {
			code = errmsg.ERROR_TOKEN_FORMAT_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		key, tokenCode := CheckToken(checkToken[1])
		if tokenCode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		if time.Now().Unix() > key.ExpiresAt.Unix() {
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		// 检查该用户是否存在
		if err := global.DBEngine.Where("username = ?", key.Username).Find(&model.User{}); err != nil {
			code = errmsg.ERROR_USER_NOT_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		c.Set("username", key.Username)
		c.Next()
	}
}
