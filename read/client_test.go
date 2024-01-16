package read

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/rslhdyt/sheetdb-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type APICallerMockGet struct {
  mock.Mock
}

// func initTest(ApiCaller sheetdb.APICaller) {
//   sheetdb.SetAPICaller(ApiCaller)
// }

func (m *APICallerMockGet) Call(ctx context.Context, method string, url string, opt sheetdb.Option, body interface{}, response interface{}) error {
  m.Called(ctx, method, url, opt, body, response)

  jsonResponse := `[
    {
      "id": 1,
      "name": "John",
      "age": 20
    },
    {
      "id": 2,
      "name": "Jane",
      "age": 19
    }
  ]`

  _ = json.Unmarshal([]byte(jsonResponse), &response)

  return nil
}

func TestGetContent(t *testing.T) {
  APICallerMockGetObj := new(APICallerMockGet)
  
  opt := sheetdb.Option{
    Username: "username",
    Password: "password",
    DocumentId: "documentId",
    BaseUrl: "https://sheetdb.io",
  }

  sheetDbClient := Client{
    Opt: &opt,
    APICaller: APICallerMockGetObj,
  }
  
  path := fmt.Sprintf("%s/api/v1/%s", "https://sheetdb.io", "documentId")

  body := sheetdb.GetContentParams{
    Sheet: "sheet",
  }

  APICallerMockGetObj.On(
    "Call",
    context.Background(),
    "GET",
    path,
    opt,
    &body,
    &sheetdb.Contents{},
  ).Return(nil)

  contents, err := sheetDbClient.GetContent(&body)

  fmt.Println(contents)

  // assert correct type data
  assert.Equal(t, reflect.TypeOf(contents), reflect.TypeOf(&sheetdb.Contents{}))

  // assert correct data
  // assert.Equal(t, contents[0]["id"], float64(1))
  // assert.Equal(t, contents[0]["name"], "John")
  // assert.Equal(t, contents[0]["age"], float64(20))

  assert.Equal(t, err, nil)
}