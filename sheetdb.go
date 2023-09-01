package sheetdb

import (
	"net/http"
	"sync"
)

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
	mu:        new(sync.RWMutex),
	apiCaller: &APICallerImpl{HTTPClient: &http.Client{}},
}

// APICallerWrapper is the APICaller with locker for setting the APICaller
type APICallerWrapper struct {
	apiCaller   APICaller
	mu          *sync.RWMutex
}

// GetAPICaller returns the current API caller.
func GetAPICaller() APICaller {
	apiCallerWrapper.mu.RLock()
	defer apiCallerWrapper.mu.RUnlock()

	return apiCallerWrapper.apiCaller
}
