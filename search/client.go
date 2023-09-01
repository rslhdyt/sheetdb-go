package search

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rslhdyt/sheetdb-go"
)

type Client struct {
	Opt       *sheetdb.Option
	APICaller sheetdb.APICaller
}

func (c *Client) Search(body *sheetdb.SearchParams) (*sheetdb.Content, error) {
	ctx := context.Background()

	response := &sheetdb.Content{}

	path := fmt.Sprintf("%s/api/v1/%s/search", c.Opt.BaseUrl, c.Opt.DocumentId)

	err := c.APICaller.Call(
		ctx,
		http.MethodGet,
		path,
		*c.Opt,
		body,
		response,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Return single object of Content
func (c *Client) Find(body *sheetdb.SearchParams) (*sheetdb.Content, error) {
  ctx := context.Background()

  response := &sheetdb.Content{}

  (*body)["single_object"] = "true"

	path := fmt.Sprintf("%s/api/v1/%s/search", c.Opt.BaseUrl, c.Opt.DocumentId)

	err := c.APICaller.Call(
		ctx,
		http.MethodGet,
		path,
		*c.Opt,
		body,
		response,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) SearchOr(body *sheetdb.SearchParams) (*sheetdb.Content, error) {
	ctx := context.Background()

	response := &sheetdb.Content{}

	path := fmt.Sprintf("%s/api/v1/%s/search_or", c.Opt.BaseUrl, c.Opt.DocumentId)

	err := c.APICaller.Call(
		ctx,
		http.MethodGet,
		path,
		*c.Opt,
		body,
		response,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}
