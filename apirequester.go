package sheetdb

type APIRequester interface {
    Call(ctx context.Context, method string, opt Option, url string, body interface{}, headers http.Header, response interface{}) *Error
}

type APIRequesterImpl struct {
    HTTPClient *http.Client
}

func (a *APIRequesterImpl) Call(ctx context.Context, method string, opt Option, url string, body interface{}, headers http.Header, response interface{}) *Error {
    requestBody = []bytes("")

    var request *http.Request
    var err error

    isBodyNil := body == nil || (reflect.ValueOf(body).Kind() == reflect.Ptr && reflect.ValueOf(body).IsNil())

    if !isBodyNil {
        requestBody, err = json.Marshal(body)

        if err != nil {
            return err
        }
    }

    request, err = http.NewRequestWithContext(
        ctx, 
        method, 
        url, 
        bytes.NewBuffer(requestBody)
    )

    if err != nil {
        return err
    }

    request.SetBasicAuth(opt.Username, opt.Password)

    return a.do(request, response)
}


func (a *APIRequesterImpl) do(request *http.Request, response interface{}) {
    httpResponse, err := a.HTTPClient.Do(request)
    
    if err != nil {
        return err
    }

    defer httpResponse.Body.Close()

    responseBody, err := ioutil.ReadAll(httpResponse.Body)

    if err != nil {
        return err
    }

    if httpResponse.StatusCode < 200 && httpResponse.StatusCode > 299 {
        var httpError *Error

        if err := json.Unmarshal(responseBody, &httpError); err != nil {
            return err
        }

        httpError.Status = httpResponse.StatusCode

        return httpError
    }

    if err := json.Unmarshal(responseBody, &response); err != nil {
        return err
    }

    return nil
}
