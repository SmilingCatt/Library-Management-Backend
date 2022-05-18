package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"lms/util"
	"net/http"
)

type Claims struct {
	UserID int
	jwt.StandardClaims
}

func UserAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		if userID, ok := auth(context, util.UserKey); ok {
			context.Set("userId", userID)
			context.Next()
		} else {
			if userID, ok = auth(context, util.AdminKey); ok {
				context.Set("userId", userID)
				context.Next()
			} else {
				context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorized"})
			}
		}
	}
}

func MemberAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if userID, ok := auth(c, util.UserKey); ok {
			c.Set("userId", userID)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorized"})
		}
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if userID, ok := auth(c, util.AdminKey); ok {
			c.Set("userId", userID)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorized"})
		}
	}
}

func auth(c *gin.Context, key []byte) (userID int, ok bool) {
	tokenString := c.PostForm("token")
	userID, ok = util.AuthToken(tokenString, key)
	return
}
