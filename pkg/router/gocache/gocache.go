package gocache

import (
	"api/performance_tune/enum"
	"api/performance_tune/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

var cacheItme *cache.Cache

func InitGoCacheRouter(r *gin.Engine) {
	cacheItme = cache.New(5*time.Minute, 10*time.Minute)
	r.GET("/"+enum.GoCache+"/items/:id", getItem)
	r.POST("/"+enum.GoCache+"/items", createItem)
	r.PUT("/"+enum.GoCache+"/items/:id", updateItem)
	r.DELETE("/"+enum.GoCache+"/items/:id", deleteItem)
}

func getItem(c *gin.Context) {
	id := c.Param("id")
	if x, found := c.Get(id); found {
		c.JSON(http.StatusOK, x)
		return
	}

	if item, found := cacheItme.Get(id); found {
		c.JSON(http.StatusOK, item)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
	}

}

func createItem(c *gin.Context) {
	var item *model.Item
	if err := c.BindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// items[item.ID] = item
	// cacheItme.Set(item.ID, item, cache.DefaultExpiration)
	cacheItme.Set(item.ID, item, cache.NoExpiration)
	c.JSON(http.StatusCreated, item)
}

func updateItem(c *gin.Context) {
	id := c.Param("id")
	var item *model.Item
	if err := c.BindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// items[id] = item
	// cacheItme.Set(id, item, cache.DefaultExpiration)
	cacheItme.Set(id, item, cache.NoExpiration)
	c.JSON(http.StatusOK, item)
}

func deleteItem(c *gin.Context) {
	id := c.Param("id")
	cacheItme.Delete(id)
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
