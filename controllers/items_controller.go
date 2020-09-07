package controllers

import (
	"fmt"
	"github.com/Anatol-e/bookstore_items_api/domain/items"
	"github.com/Anatol-e/bookstore_items_api/service"
	"net/http"
)

var (
	ItemControllers itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	result, err := service.ItemsService.Create(items.Item{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	result, err := service.ItemsService.Get("response_id")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
