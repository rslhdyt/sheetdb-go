package client

type API struct {
	opt 			sheetdb.Option
	apiRequester 	sheetdb.APIRequester

	Read 			*read.Client
	// Search 			*search.Client
	// Create 			*create.Client
	// Update 			*update.Client
	// Delete			*delete.Client
}

func (a *API) init() {
	a.Read = &read.Client{ opt: a.opt, apiRequester: a.apiRequester }
}

func New(username string, password string, documentId string) *API {
	api := API{
		opt: sheetdb.Option{
			Username: sheetdb.Option.Username,
			Password: password,
			DocumentId: documentId,

			BaseUrl: "https://sheetdb.io",
		},
		apiRequester: sheetdb.GetAPIRequester(),
	}

	api.init()

	return &api
}

