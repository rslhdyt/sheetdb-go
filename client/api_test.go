package client

import (
	"testing"

	"github.com/rslhdyt/sheetdb-go"
	"github.com/stretchr/testify/assert"
)

func TestApi(t *testing.T) {
  api := API{
    opt: sheetdb.Option{
      Username:   "username",
      Password:   "password",
      DocumentId: "documentId",
      BaseUrl:    "baseUrl",
    },
  }

  api.init()

  assert.Equal(t, *api.Read.Opt, api.opt)
}

func TestNew(t *testing.T) {
  api := New("username", "password", "documentId")

  assert.Equal(t, api.opt.Username, "username")
  assert.Equal(t, api.opt.Password, "password")
  assert.Equal(t, api.opt.DocumentId, "documentId")
  assert.Equal(t, api.opt.BaseUrl, "https://sheetdb.io")
}