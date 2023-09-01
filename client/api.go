package client

import (
	"github.com/rslhdyt/sheetdb-go"
	"github.com/rslhdyt/sheetdb-go/read"
	"github.com/rslhdyt/sheetdb-go/search"
)

type API struct {
	opt          sheetdb.Option
	apiCaller sheetdb.APICaller

	Read   *read.Client
	Search *search.Client
	// Create		 *create.Client
	// Update	 	 *update.Client
	// Delete		 *delete.Client
}

func (a *API) init() {
	a.Read   = &read.Client{Opt: &a.opt, APICaller: a.apiCaller}
	a.Search = &search.Client{Opt: &a.opt, APICaller: a.apiCaller}
}

func New(username string, password string, documentId string) *API {
	api := API{
		opt: sheetdb.Option{
			Username:   username,
			Password:   password,
			DocumentId: documentId,

			BaseUrl: "https://sheetdb.io",
		},
		apiCaller: sheetdb.GetAPICaller(),
	}

	api.init()

	return &api
}
