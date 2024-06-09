package fast

import (
	"api/performance_tune/enum"
	"api/performance_tune/model"
	"encoding/json"
	"fmt"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

var items = map[string]*model.Item{}

func InitFastHttpRouter(fr *fasthttprouter.Router) {
	items = make(map[string]*model.Item)
	fr.GET("/"+enum.FastHttp+"/items/:id", getItem)
	fr.POST("/"+enum.FastHttp+"/items", createItem)
	fr.PUT("/"+enum.FastHttp+"/items/:id", updateItem)
	fr.DELETE("/"+enum.FastHttp+"/items/:id", deleteItem)
}

func getItem(c *fasthttp.RequestCtx) {
	id := c.UserValue("id")
	c.Response.Header.Set("Content-Type", "application/json")

	if item, exists := items[id.(string)]; exists && item != nil {
		c.Response.Header.SetStatusCode(fasthttp.StatusOK)
		c.Response.SetBodyString(fmt.Sprintf("%v", item))
	} else {
		c.Response.Header.SetStatusCode(fasthttp.StatusNotFound)
		c.Response.SetBodyString(`{"error": "Item not found"}`)
	}
}

func createItem(c *fasthttp.RequestCtx) {
	var item *model.Item
	c.Response.Header.Set("Content-Type", "application/json")

	err := json.Unmarshal(c.PostBody(), &item)
	if err != nil {
		c.Response.SetStatusCode(fasthttp.StatusNotImplemented)
		c.Error("Failed to parse request", fasthttp.StatusBadRequest)
		return
	}

	items[item.ID] = item
	c.Response.SetStatusCode(fasthttp.StatusOK)
	c.Response.SetBodyString(fmt.Sprintf("Item %s saved successfully", item.ID))
}

func updateItem(c *fasthttp.RequestCtx) {
	id := c.UserValue("id")
	c.Response.Header.Set("Content-Type", "application/json")

	if item, exists := items[id.(string)]; exists && item != nil {
		var item *model.Item
		if err := json.Unmarshal(c.PostBody(), &item); err != nil {
			c.Response.Header.SetStatusCode(fasthttp.StatusNotFound)
			c.Response.SetBodyString(`{"error": "Post Body not found"}`)
		}

		items[id.(string)] = item
		c.Response.SetStatusCode(fasthttp.StatusOK)
		c.Response.SetBodyString(fmt.Sprintf("Item %s saved successfully", item.ID))
	} else {
		c.Response.Header.SetStatusCode(fasthttp.StatusOK)
		c.Response.SetBodyString("id can't find")
	}
}

func deleteItem(c *fasthttp.RequestCtx) {
	id := c.UserValue("id")
	c.Response.Header.Set("Content-Type", "application/json")

	if item, exists := items[id.(string)]; exists && item != nil {
		delete(items, id.(string))
		c.Response.SetStatusCode(fasthttp.StatusOK)
		c.Response.SetBodyString(fmt.Sprintf("Item %s saved successfully", item.ID))
	} else {
		c.Response.SetStatusCode(fasthttp.StatusOK)
		c.Response.SetBodyString("id can't find")
	}
}
