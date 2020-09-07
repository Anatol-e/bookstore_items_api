package service

import (
	"github.com/Anatol-e/bookstore_items_api/domain/items"
	"github.com/Anatol-e/bookstore_items_api/utils/errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, *errors.RestErr)
	Get(string) (*items.Item, *errors.RestErr)
}

type itemsService struct{}

func (is *itemsService) Create(item items.Item) (*items.Item, *errors.RestErr) {
	return nil, errors.NewInternalServerError("create service not implement")
}

func (is *itemsService) Get(id string) (*items.Item, *errors.RestErr) {
	return nil, errors.NewInternalServerError("get service not implement")
}
