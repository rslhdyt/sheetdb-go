package sheetdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Option struct {
	Username 	string
	Password 	string
	DocumentId 	string

	BaseUrl 	string
}