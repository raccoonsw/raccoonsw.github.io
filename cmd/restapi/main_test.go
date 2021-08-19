package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"restApiProject/database"
	"restApiProject/models"
	"testing"
)

var router *gin.Engine
var sqlDB database.DBModel

func init() {
	var s Specification
	err := envconfig.Process("restapiproject", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	config := models.Config{DBUser: s.DBUser, DBPassword: s.DBPassword, DBHost: s.DBHost, DBPort: s.DBPort, DBName: s.DBName}
	sqlDB = database.DBModel{DB: database.Connect(config)}
	//defer sqlDB.Close()

	router = setupRouter(&sqlDB)
}

func TestPingRoute(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"status":"ok"}`, w.Body.String())
}

func TestCreateItem(t *testing.T) {
	sqlDB.ClearTable()
	w := httptest.NewRecorder()
	var jsonData = []byte(`{"sku": "big_rocket", "name": "Big Rocket", "type": "virtual_currency", "cost": 0.5}`)
	req, _ := http.NewRequest("POST", "/api/items", bytes.NewBuffer(jsonData))
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
	req, _ := http.NewRequest("PUT", "/api/items/1", bytes.NewBuffer(jsonData))
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
	req, _ := http.NewRequest("PUT", "/api/items/sku/big_rocket", bytes.NewBuffer(jsonData))
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
	req, _ := http.NewRequest("DELETE", "/api/items/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestDeleteItemBySku(t *testing.T) {
	sqlDB.ClearTable()
	item := models.Item{Sku: "big_rocket", Name: "Big Rocket", Type: "virtual_currency", Cost: 0.5}
	_, _ = sqlDB.CreateItem(item)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/items/sku/big_rocket", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestGetItemById(t *testing.T) {
	sqlDB.ClearTable()
	item := models.Item{Sku: "big_rocket", Name: "Big Rocket", Type: "virtual_currency", Cost: 0.5}
	newItem, _ := sqlDB.CreateItem(item)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/items/1", nil)
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
	req, _ := http.NewRequest("GET", "/api/items/sku/big_rocket", nil)
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
	req, _ := http.NewRequest("GET", "/api/items?limit=1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var responsePaging1 Paging
	_ = json.Unmarshal([]byte(w.Body.String()), &responsePaging1)
	assert.Equal(t, true, responsePaging1.HasMore)
	assert.Equal(t, []models.Item{newItem1}, responsePaging1.Items)

	//get second record
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/items?limit=1&offset=1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var responsePaging2 Paging
	_ = json.Unmarshal([]byte(w.Body.String()), &responsePaging2)
	assert.Equal(t, false, responsePaging2.HasMore)
	assert.Equal(t, []models.Item{newItem2}, responsePaging2.Items)
}

func TestGetAllItemsFilter(t *testing.T) {
	sqlDB.ClearTable()
	_, _ = sqlDB.CreateItem(models.Item{Sku: "big_rocket", Name: "Big Rocket", Type: "virtual_currency", Cost: 0.5})
	newItem2, _ := sqlDB.CreateItem(models.Item{Sku: "bla_rocket", Name: "Bla Rocket", Type: "virtual_good", Cost: 1.5})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/items?type=virtual_good", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var responsePaging1 Paging
	_ = json.Unmarshal([]byte(w.Body.String()), &responsePaging1)
	assert.Len(t, responsePaging1.Items, 1)
	assert.Equal(t, []models.Item{newItem2}, responsePaging1.Items)
}
