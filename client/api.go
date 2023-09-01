package client

import (
	"github.com/rslhdyt/sheetdb-go"
	"github.com/rslhdyt/sheetdb-go/read"
)

type API struct {
	opt          sheetdb.Option
	apiRequester sheetdb.APIRequester

	Read *read.Client
	// Search 			*search.Client
	// Create 			*create.Client
	// Update 			*update.Client
	// Delete			*delete.Client
}

func (a *API) init() {
	a.Read = &read.Client{Opt: &a.opt, APIRequester: a.apiRequester}
}

func New(username string, password string, documentId string) *API {
	api := API{
		opt: sheetdb.Option{
			Username:   username,
			Password:   password,
			DocumentId: documentId,

			BaseUrl: "https://sheetdb.io",
			// BaseUrl: "https://webhook.site/08b7687d-e470-4d40-8ee6-2072204685f8",
		},
		apiRequester: sheetdb.GetAPIRequester(),
	}

	api.init()

	return &api
}
