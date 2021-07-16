package validation

import (
	"errors"
	"github.com/gin-gonic/gin"
	"regexp"
	"restApiProject/models"
)

func isSku(str string) error {
	var validID = regexp.MustCompile(`^[a-z0-9_.-]*$`)
	if isSku := validID.MatchString(str); !isSku {
		return errors.New("parameter sku should contain lowercase Latin alphanumeric characters, periods, dashes, and underscores")
	}
	return nil
}

func ValidatePathId(c *gin.Context) (uint, error) {
	var id struct {
		Id uint `uri:"id" binding:"required"`
	}
	err := c.BindUri(&id)
	return id.Id, err
}

func ValidatePathSku(c *gin.Context) (string, error) {
	var sku struct {
		Sku string `uri:"sku" binding:"required,ascii,lowercase,lte=100"`
	}
	err := c.BindUri(&sku)
	if err := isSku(sku.Sku); err != nil {
		return sku.Sku, err
	}
	return sku.Sku, err
}

func ValidatePathIdJsonItem(c *gin.Context) (models.Item, error) {
	id, errId := ValidatePathId(c)
	if errId != nil {
		return models.Item{}, errId
	}
	receivedItem, errItem := ValidateJsonItem(c)
	if errItem == nil {
		receivedItem.Id = id
	}
	return receivedItem, errItem
}

func ValidatePathSkuJsonItem(c *gin.Context) (models.Item, error) {
	sku, errId := ValidatePathSku(c)
	if errId != nil {
		return models.Item{}, errId
	}
	var receivedItem struct {
		Name string  `json:"name" binding:"required,lte=100,ascii"`
		Type string  `gorm:"not null;type:varchar(100)" json:"type" binding:"required,oneof=virtual_good virtual_currency bundle"`
		Cost float32 `json:"cost" binding:"required"`
		//UserId	uint	//`gorm:"not null"`
	}
	errItem := c.Bind(&receivedItem)
	item := models.Item{Name: receivedItem.Name, Type: receivedItem.Type, Cost: receivedItem.Cost}
	if errItem == nil {
		item.Sku = sku
	}
	return item, errItem
}

func ValidateJsonItem(c *gin.Context) (models.Item, error) {
	var receivedItem models.Item
	err := c.Bind(&receivedItem)
	if err := isSku(receivedItem.Sku); err != nil {
		return receivedItem, err
	}
	return receivedItem, err
}

func ValidateQueryPaging(c *gin.Context) (models.Filter, error) {
	var filter models.Filter
	err := c.BindQuery(&filter)
	if filter.Limit == 0 {
		filter.Limit = 30
	}
	return filter, err
}
