package controllers

import (
	"golang/initializers"
	"golang/models"

	"github.com/gin-gonic/gin"
)

func AuthorRead(c *gin.Context) {
	authors := []models.Author{}

	if err := initializers.DB.Find(&authors).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to fetch authors",
		})

		return
	}

	c.JSON(200, gin.H{
		"authors": authors,
	})
}

func AuthorReadOne(c *gin.Context) {
	author := models.Author{}

	if err := initializers.DB.First(&author, c.Param("id")).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to fetch author",
		})

		return
	}

	c.JSON(200, gin.H{
		"author": author,
	})
}

func AuthorUpdate(c *gin.Context) {
	author := models.Author{}

	if err := initializers.DB.First(&author, c.Param("id")).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to fetch author",
		})

		return
	}

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to bind author",
		})

		return
	}

	if err := initializers.DB.Save(&author).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to update author",
		})

		return
	}

	c.JSON(200, gin.H{
		"author": author,
	})
}

func AuthorDelete(c *gin.Context) {
	author := models.Author{}

	if err := initializers.DB.First(&author, c.Param("id")).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to fetch author",
		})
	}

	if err := initializers.DB.Delete(&author).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to delete author",
		})
	}

	c.JSON(200, gin.H{
		"author": author,
	})
}

func AuthorMyProfile(c *gin.Context) {
	author, _ := c.Get("author")

	c.JSON(200, gin.H{
		"author": author,
	})
}
