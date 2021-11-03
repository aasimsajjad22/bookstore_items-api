package services

import (
	"github.com/aasimsajjad22/bookstore_items-api/domain/items"
	"github.com/aasimsajjad22/bookstore_utils-go/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, *rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

type itemsService struct {}

func (s *itemsService) Get(id string) (*items.Item, rest_errors.RestErr) {
	panic("implement me")
}

func (s *itemsService) Create(item items.Item) (*items.Item, *rest_errors.RestErr) {
	return &item, nil
}

