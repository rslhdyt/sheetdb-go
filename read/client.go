package read

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rslhdyt/sheetdb-go"
)

type Client struct {
	Opt          *sheetdb.Option
	APIRequester sheetdb.APIRequester
}

func (c *Client) GetContent(body *sheetdb.GetContentParams) (*sheetdb.Content, error) {
	ctx := context.Background()

	response := &sheetdb.Content{}
	headers := http.Header{}

	path := fmt.Sprintf("%s/api/v1/%s", c.Opt.BaseUrl, c.Opt.DocumentId)

	err := c.APIRequester.Call(
		ctx,
		http.MethodGet,
		path,
		*c.Opt,
		body,
		headers,
		response,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetKeys() (*sheetdb.TableKeys, error) {
	ctx := context.Background()

	response := &sheetdb.TableKeys{}
	headers := http.Header{}

	path := fmt.Sprintf("%s/api/v1/%s/keys", c.Opt.BaseUrl, c.Opt.DocumentId)

	err := c.APIRequester.Call(
		ctx,
		http.MethodGet,
		path,
		*c.Opt,
		nil,
		headers,
		response,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) DocumentName() (*sheetdb.DocumentName, error) {
	ctx := context.Background()

	response := &sheetdb.DocumentName{}
	headers := http.Header{}

	path := fmt.Sprintf("%s/api/v1/%s/name", c.Opt.BaseUrl, c.Opt.DocumentId)

	err := c.APIRequester.Call(
		ctx,
		http.MethodGet,
		path,
		*c.Opt,
		nil,
		headers,
		response,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) Count() (*sheetdb.Count, error) {
	ctx := context.Background()

	response := &sheetdb.Count{}
	headers := http.Header{}

	path := fmt.Sprintf("%s/api/v1/%s/count", c.Opt.BaseUrl, c.Opt.DocumentId)

	err := c.APIRequester.Call(
		ctx,
		http.MethodGet,
		path,
		*c.Opt,
		nil,
		headers,
		response,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}
