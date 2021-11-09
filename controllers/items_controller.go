package controllers

import (
	"encoding/json"
	"github.com/aasimsajjad22/bookstore_items-api/domain/items"
	"github.com/aasimsajjad22/bookstore_items-api/services"
	"github.com/aasimsajjad22/bookstore_items-api/utils/http_utils"
	"github.com/aasimsajjad22/bookstore_oauth-go/oauth"
	"github.com/aasimsajjad22/bookstore_utils-go/rest_errors"
	"io/ioutil"
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
		//http_utils.RespondError(w, err)
		return
	}
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	itemRequest.Seller = oauth.GetCallerId(r)
	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondError(w, createErr)
		return
	}
	http_utils.RespondJson(w, http.StatusCreated, result)
}
