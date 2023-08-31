package read

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	opt 			sheetdb.Option
	apiRequester 	sheetdb.APIRequester
}

func (c *Client) GetContent(data *GetContentParams) (*sheetdb.Content, *sheetdb.Error) {
	ctx := context.Background()

	response := &sheetdb.Content{}
	headers := http.Header{}

	path = fmt.Sprintf("%s/api/v1/%s", c.opt.BaseUrl, c.opt.DocumentId)

	err := APIRequester.Call(
		ctx, 
		http.MethodGet, 
		c.opt, 
		path, 
		data, 
		headers, 
		response
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetKeys() (*sheetdb.TableKeys, *sheetdb.Error) {
	ctx := context.Background()

	response := &sheetdb.TableKeys{}
	headers := http.Header{}

	path = fmt.Sprintf("%s/api/v1/%s/keys", c.opt.BaseUrl, c.opt.DocumentId)

	err := APIRequester.Call(
		ctx, 
		http.MethodGet, 
		c.opt, 
		path, 
		nil, 
		headers, 
		response
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) DocumentName() (*sheetdb.DocumentName, *sheetdb.Error) {
	ctx := context.Background()

	response := &sheetdb.DocumentName{}
	headers := http.Header{}

	path = fmt.Sprintf("%s/api/v1/%s/name", c.opt.BaseUrl, c.opt.DocumentId)

	err := APIRequester.Call(
		ctx, 
		http.MethodGet, 
		c.opt, 
		path, 
		nil, 
		headers, 
		response
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) Count() (*sheetdb.Count, *sheetdb.Error) {
	ctx := context.Background()

	response := &sheetdb.Count{}
	headers := http.Header{}

	path = fmt.Sprintf("%s/api/v1/%s/count", c.opt.BaseUrl, c.opt.DocumentId)

	err := APIRequester.Call(
		ctx, 
		http.MethodGet, 
		c.opt, 
		path, 
		nil, 
		headers, 
		response
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}