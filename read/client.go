package read

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

func (c *Client) GetContent(body *sheetdb.GetContentParams) (*sheetdb.Contents, error) {
	ctx := context.Background()

	response := &sheetdb.Contents{}

	path := fmt.Sprintf("%s/api/v1/%s", c.Opt.BaseUrl, c.Opt.DocumentId)

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

func (c *Client) GetKeys() (*sheetdb.TableKeys, error) {
	ctx := context.Background()

	response := &sheetdb.TableKeys{}

	path := fmt.Sprintf("%s/api/v1/%s/keys", c.Opt.BaseUrl, c.Opt.DocumentId)

	err := c.APICaller.Call(
		ctx,
		http.MethodGet,
		path,
		*c.Opt,
		nil,
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

	path := fmt.Sprintf("%s/api/v1/%s/name", c.Opt.BaseUrl, c.Opt.DocumentId)

	err := c.APICaller.Call(
		ctx,
		http.MethodGet,
		path,
		*c.Opt,
		nil,
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

	path := fmt.Sprintf("%s/api/v1/%s/count", c.Opt.BaseUrl, c.Opt.DocumentId)

	err := c.APICaller.Call(
		ctx,
		http.MethodGet,
		path,
		*c.Opt,
		nil,
		response,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}
