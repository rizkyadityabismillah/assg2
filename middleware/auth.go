package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)
type Claims struct {
    UserID int `json:"user_id"`
    jwt.StandardClaims
  }
  func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie("session_token")
		if err != nil || cookie == nil {
			if ctx.GetHeader("Content-Type") == "application/json" {
				ctx.AbortWithStatus(http.StatusUnauthorized)
			} else {
				ctx.Redirect(http.StatusFound, "/login")
			}
			return
		}
		
		tokenString := cookie.Value
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Set("id", claims.UserID)
		
		
	})
}



