package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"restApiProject/models"
	"restApiProject/validation"
	"strconv"
)

type Env struct {
	DBModel models.DBModel
}

// GET /healthcheck
func (env *Env) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// POST /item
func (env *Env) CreateItem(c *gin.Context) {
	receivedItem, err := validation.ValidateJsonItem(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item, err := env.DBModel.CreateItem(receivedItem)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": item.Id})
}

// PUT /item/:id
func (env *Env) UpdateItemById(c *gin.Context) {
	receivedItem, err := validation.ValidatePathIdJsonItem(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = env.DBModel.UpdateItemById(receivedItem)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// PUT /item/sku/:sku
func (env *Env) UpdateItemBySku(c *gin.Context) {
	receivedItem, err := validation.ValidatePathSkuJsonItem(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = env.DBModel.UpdateItemBySku(receivedItem)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// DELETE /item/:id
func (env *Env) DeleteItemById(c *gin.Context) {
	id, err := validation.ValidatePathId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = env.DBModel.DeleteItemById(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// DELETE /item/sku/:sku
func (env *Env) DeleteItemBySku(c *gin.Context) {
	sku, err := validation.ValidatePathSku(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = env.DBModel.DeleteItemBySku(sku)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// GET /item/:id
func (env *Env) GetItemById(c *gin.Context) {
	id, err := validation.ValidatePathId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item, err := env.DBModel.GetItemById(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

// GET /item/sku/:sku
func (env *Env) GetItemBySku(c *gin.Context) {
	sku, err := validation.ValidatePathSku(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item, err := env.DBModel.GetItemBySku(sku)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

// GET /items
func (env *Env) GetAllItems(c *gin.Context) {
	var filter models.Filter
	filter, err := validation.ValidateQueryPaging(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	items, hasMore, err := env.DBModel.GetAllItems(filter)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.Header("has-more", strconv.FormatBool(hasMore))
	c.JSON(http.StatusOK, gin.H{"has_more": hasMore, "items": items})
}
