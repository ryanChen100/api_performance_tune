package syncpool

import (
	"api/performance_tune/model"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var pool = sync.Pool{
	New: func() interface{} {
		return &model.Item{}
	},
}

func InitSyncPoolRouter(r *gin.Engine) {
	r.GET("/syncpool/items/:id", getItem)
	r.POST("/syncpool/items", createItem)
	r.PUT("/syncpool/items/:id", updateItem)
	r.DELETE("/syncpool/items/:id", deleteItem)

}

func getItem(c *gin.Context) {
	item := pool.Get().(*model.Item)
	defer pool.Put(item)
	if item == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func createItem(c *gin.Context) {
	item := pool.Get().(*model.Item)
	defer pool.Put(item)
	if err := c.BindJSON(item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, item)
}

func updateItem(c *gin.Context) {
	item := pool.Get().(*model.Item)
	defer pool.Put(item)
	if err := c.BindJSON(item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

func deleteItem(c *gin.Context) {
	item := pool.Get().(*model.Item)
	if item == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	} else {
		c.JSON(http.StatusOK, item)
	}
}
