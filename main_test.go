package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"restApiProject/models"
	"testing"
)

var router *gin.Engine
var sqlDB models.DBModel

func init() {
	config := models.Config{User: "root", Password: "g7y48UPH", DBName: "restApi"}
	sqlDB = models.DBModel{DB: models.Connect(config)}
	//defer sqlDB.Close()

	router = setupRouter(sqlDB)
}

func TestPingRoute(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"status":"ok"}`, w.Body.String())
}

func TestCreateItem(t *testing.T) {
	sqlDB.ClearTable()
	w := httptest.NewRecorder()
	var jsonData = []byte(`{"sku": "big_rocket", "name": "Big Rocket", "type": "virtual_currency", "cost": 0.5}`)
	req, _ := http.NewRequest("POST", "/item", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"id":1}`, w.Body.String())
}

func TestUpdateItemById(t *testing.T) {
	sqlDB.ClearTable()
	item := models.Item{Sku: "big_rocket", Name: "Big Rocket", Type: "virtual_currency", Cost: 0.5}
	newItem, _ := sqlDB.CreateItem(item)
	w := httptest.NewRecorder()
	var jsonData = []byte(`{"sku": "bla_rocket", "name": "Bla Rocket", "type": "virtual_good", "cost": 1.5}`)
	req, _ := http.NewRequest("PUT", "/item/1", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

	item, _ = sqlDB.GetItemById(newItem.Id)
	_ = json.Unmarshal(jsonData, &newItem)
	assert.Equal(t, newItem, item)
}

func TestUpdateItemBySku(t *testing.T) {
	sqlDB.ClearTable()
	item := models.Item{Sku: "big_rocket", Name: "Big Rocket", Type: "virtual_currency", Cost: 0.5}
	newItem, _ := sqlDB.CreateItem(item)
	w := httptest.NewRecorder()
	var jsonData = []byte(`{"name": "Bla Rocket", "type": "virtual_good", "cost": 1.5}`)
	req, _ := http.NewRequest("PUT", "/item/sku/big_rocket", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

	item, _ = sqlDB.GetItemBySku(newItem.Sku)
	_ = json.Unmarshal(jsonData, &newItem)
	assert.Equal(t, newItem, item)
}

func TestDeleteItemById(t *testing.T) {
	sqlDB.ClearTable()
	item := models.Item{Sku: "big_rocket", Name: "Big Rocket", Type: "virtual_currency", Cost: 0.5}
	_, _ = sqlDB.CreateItem(item)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/item/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestDeleteItemBySku(t *testing.T) {
	sqlDB.ClearTable()
	item := models.Item{Sku: "big_rocket", Name: "Big Rocket", Type: "virtual_currency", Cost: 0.5}
	_, _ = sqlDB.CreateItem(item)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/item/sku/big_rocket", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestGetItemById(t *testing.T) {
	sqlDB.ClearTable()
	item := models.Item{Sku: "big_rocket", Name: "Big Rocket", Type: "virtual_currency", Cost: 0.5}
	newItem, _ := sqlDB.CreateItem(item)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/item/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	_ = json.Unmarshal([]byte(w.Body.String()), &item)
	assert.Equal(t, newItem, item)
}

func TestGetItemBySku(t *testing.T) {
	sqlDB.ClearTable()
	item := models.Item{Sku: "big_rocket", Name: "Big Rocket", Type: "virtual_currency", Cost: 0.5}
	newItem, _ := sqlDB.CreateItem(item)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/item/sku/big_rocket", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	_ = json.Unmarshal([]byte(w.Body.String()), &item)
	assert.Equal(t, newItem, item)
}

type Paging struct {
	HasMore bool `json:"has_more"`
	Items   []models.Item
}

func TestGetAllItems(t *testing.T) {
	sqlDB.ClearTable()
	newItem1, _ := sqlDB.CreateItem(models.Item{Sku: "big_rocket", Name: "Big Rocket", Type: "virtual_currency", Cost: 0.5})
	newItem2, _ := sqlDB.CreateItem(models.Item{Sku: "bla_rocket", Name: "Bla Rocket", Type: "virtual_good", Cost: 1.5})
	//get first record
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items?limit=1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var responsePaging1 Paging
	_ = json.Unmarshal([]byte(w.Body.String()), &responsePaging1)
	assert.Equal(t, true, responsePaging1.HasMore)
	assert.Equal(t, []models.Item{newItem1}, responsePaging1.Items)

	//get second record
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/items?limit=1&offset=1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var responsePaging2 Paging
	_ = json.Unmarshal([]byte(w.Body.String()), &responsePaging2)
	assert.Equal(t, false, responsePaging2.HasMore)
	assert.Equal(t, []models.Item{newItem2}, responsePaging2.Items)
}