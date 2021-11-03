package controllers

import (
	"fmt"
	"github.com/aasimsajjad22/bookstore_items-api/domain/items"
	"github.com/aasimsajjad22/bookstore_items-api/services"
	"github.com/aasimsajjad22/bookstore_oauth-go/oauth"
	"net/http"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
}

type itemsController struct {}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request)  {
	if err := oauth.AuthenticateRequest(r); err != nil {
		return
	}
	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}
	result, err := services.ItemsService.Create(item)
	if err != nil {

	}
	fmt.Println(result)
}
