package controllers

import (
	"fmt"
	"golang/initializers"
	"golang/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struct {
		Name     string
		Email    string
		Password string
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid input",
		})

		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	if err := initializers.DB.Create(&models.Author{
		Name:      body.Name,
		Email:     body.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to create author",
		})

		return
	}

	c.JSON(200, gin.H{
		"message": "Author created successfully",
	})
}

func Login(c *gin.Context) {
	author := models.Author{}

	var body struct {
		Email    string
		Password string
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid input",
		})

		return
	}

	if err := initializers.DB.First(&author, "email = ?", body.Email).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Error fetching author",
		})

		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(author.Password), []byte(body.Password)); err != nil {
		c.JSON(500, gin.H{
			"error": "Invalid credentials",
		})

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": author.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		fmt.Println("Error creating token:", err)
		c.JSON(500, gin.H{
			"error": "Failed to create token",
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("jwt", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(200, gin.H{
		"message": "Authenticated",
	})
}
