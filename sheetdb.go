package sheetdb

import (
	"net/http"
	"sync"
)

type Option struct {
	Username   string
	Password   string
	DocumentId string

	BaseUrl string
}

var apiRequesterWrapper APIRequesterWrapper = APIRequesterWrapper{
	mu:           new(sync.RWMutex),
	apiRequester: &APIRequesterImpl{HTTPClient: &http.Client{}},
}

// APIRequesterWrapper is the APIRequester with locker for setting the APIRequester
type APIRequesterWrapper struct {
	apiRequester APIRequester
	mu           *sync.RWMutex
}

func GetAPIRequester() APIRequester {
	apiRequesterWrapper.mu.RLock()
	defer apiRequesterWrapper.mu.RUnlock()

	return apiRequesterWrapper.apiRequester
}
