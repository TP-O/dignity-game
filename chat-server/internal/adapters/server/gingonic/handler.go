package gingonic

import (
	"github.com/gin-gonic/gin"
	"chat-server/internal/apps/models"
)

func (a *Adapter) addItem(c *gin.Context) {
	var item models.Item

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H {
			"message": "Invalid request body",
		})
		return
	}

	if err := a.api.AddItem(item.Name); err != nil {
		c.JSON(500, gin.H {
			"message": "Error adding item",
		})
		return
	}

	c.JSON(200, gin.H {
		"message": "Item added successfully",
	})
}

func (a *Adapter) getItems(c *gin.Context) {
	var items []*models.Item

	items, err := a.api.GetItems()
	if err != nil {
		c.JSON(500, gin.H {
			"message": "Error getting items",
		})
		return
	}

	c.JSON(200, gin.H {
		"items": items,
	})
}