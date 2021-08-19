package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"restApiProject/database"
	"restApiProject/models"
	"restApiProject/validation"
	"strconv"
)

// CreateItem POST /item
func CreateItem(db database.ItemsInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		receivedItem, err := validation.ValidateJsonItem(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		item, err := db.CreateItem(receivedItem)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"id": item.Id})
	}
}

// UpdateItemById PUT /item/:id
func UpdateItemById(db database.ItemsInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		receivedItem, err := validation.ValidatePathIdJsonItem(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = db.UpdateItemById(receivedItem)
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
}

// UpdateItemBySku PUT /item/sku/:sku
func UpdateItemBySku(db database.ItemsInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		receivedItem, err := validation.ValidatePathSkuJsonItem(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = db.UpdateItemBySku(receivedItem)
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
}

// DeleteItemById DELETE /item/:id
func DeleteItemById(db database.ItemsInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := validation.ValidatePathId(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = db.DeleteItemById(id)
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
}

// DeleteItemBySku DELETE /item/sku/:sku
func DeleteItemBySku(db database.ItemsInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		sku, err := validation.ValidatePathSku(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = db.DeleteItemBySku(sku)
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
}

// GetItemById GET /item/:id
func GetItemById(db database.ItemsInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := validation.ValidatePathId(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		item, err := db.GetItemById(id)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		c.Header("cache-control", "no-cache, no-store, must-revalidate")
		c.JSON(http.StatusOK, item)
	}
}

// GetItemBySku GET /item/sku/:sku
func GetItemBySku(db database.ItemsInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		sku, err := validation.ValidatePathSku(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		item, err := db.GetItemBySku(sku)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		c.Header("cache-control", "no-cache, no-store, must-revalidate")
		c.JSON(http.StatusOK, item)
	}
}

// GetAllItems GET /items
func GetAllItems(db database.ItemsInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter models.Paging
		filter, err := validation.ValidateQueryPaging(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		items, hasMore, err := db.GetAllItems(filter)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		c.Header("cache-control", "no-cache, no-store, must-revalidate")
		c.Header("has-more", strconv.FormatBool(hasMore))
		c.JSON(http.StatusOK, gin.H{"has_more": hasMore, "items": items})
	}
}
