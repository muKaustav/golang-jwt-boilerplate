package middlewares

import (
	"fmt"
	"golang/initializers"
	"golang/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func IsAuth(c *gin.Context) {
	tokenString, err := c.Cookie("jwt")

	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			c.JSON(500, gin.H{"error": "Internal server error"})
			c.Abort()
			return nil, fmt.Errorf("internal server error")
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()

		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if time.Now().Unix() > int64(claims["exp"].(float64)) {
			c.JSON(401, gin.H{"error": "Unauthorized, token expired"})
			c.Abort()

			return
		}

		author := models.Author{}
		initializers.DB.First(&author, claims["sub"].(float64))

		if author.ID == 0 {
			c.JSON(401, gin.H{"error": "Unauthorized, author not found"})
			c.Abort()

			return
		}

		c.Set("author", author)
	} else {
		c.JSON(401, gin.H{"error": "Unauthorized, invalid token"})
		c.Abort()

		return
	}

	c.Next()
}
