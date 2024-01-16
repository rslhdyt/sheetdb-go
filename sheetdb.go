package sheetdb

import (
	"net/http"
	"sync"
)

var Opt Option = Option{
  BaseUrl: "https://sheetdb.io",
}

// Option represents a set of options that can be used to configure the behavior of a client.
type Option struct {
	Username   string
	Password   string
	DocumentId string

	BaseUrl string
}

// APICallerWrapper wraps an APICaller and guards all methods with a mutex.
// It provides thread-safe access to the APICaller.
var apiCallerWrapper APICallerWrapper = APICallerWrapper{
	mutex:     new(sync.RWMutex),
  apiCaller: &APICallerImpl{HTTPClient: &http.Client{}},
}

// APICallerWrapper is the APICaller with locker for setting the APICaller
type APICallerWrapper struct {
	mutex       *sync.RWMutex
	apiCaller   APICaller
}

// GetAPICaller returns the current API caller.
func GetAPICaller() APICaller {
	apiCallerWrapper.mutex.RLock()
	defer apiCallerWrapper.mutex.RUnlock()

	return apiCallerWrapper.apiCaller
}

func SetAPICaller(apiCaller APICaller) {
  apiCallerWrapper.mutex.Lock()
  defer apiCallerWrapper.mutex.Unlock()

  apiCallerWrapper.apiCaller = apiCaller
}