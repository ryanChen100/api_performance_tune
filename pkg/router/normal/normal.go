package normal

import (
	"api/performance_tune/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var items = map[string]model.Item{}

func InitRouter(r *gin.Engine) {
	items = make(map[string]model.Item)
	r.GET("/items/:id", getItem)
	r.POST("/items", createItem)
	r.PUT("/items/:id", updateItem)
	r.DELETE("/items/:id", deleteItem)
}

func getItem(c *gin.Context) {
	id := c.Param("id")
	if item, exists := items[id]; exists {
		c.JSON(http.StatusOK, item)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
	}
}

func createItem(c *gin.Context) {
	var item model.Item
	if err := c.BindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	items[item.ID] = item
	c.JSON(http.StatusCreated, item)
}

func updateItem(c *gin.Context) {
	id := c.Param("id")
	var item model.Item
	if err := c.BindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	items[id] = item
	c.JSON(http.StatusOK, item)
}

func deleteItem(c *gin.Context) {
	id := c.Param("id")
	if _, exists := items[id]; exists {
		delete(items, id)
		c.JSON(http.StatusOK, gin.H{"status": "deleted"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
	}
}
