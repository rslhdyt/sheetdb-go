# SheetDB GO Client
The Unoficial [SheetDB](https://sheetdb.io) client written in go 

## Documentation
For the API documentation, check [SheetDB API Docs](https://docs.sheetdb.io/).

## Installation
Install sheetDB go with

```bash
go get -u github.com/rslhdyt/sheetdb-go
```

Import using
    
```go
import (
    "github.com/rslhdyt/sheetdb-go"
    "github.com/rslhdyt/sheetdb-go/client"
)
```

## Usage

### Create a client
```go
package main

import (
    "github.com/rslhdyt/sheetdb-go"
    "github.com/rslhdyt/sheetdb-go/client"
)

func main() {
    ...

    // Basic setup
    sheetDB := client.New(sheetDBUsername, sheetDBPassword, sheetDBDocumentId)

    // Set params
    getContentParams := sheetdb.GetContentParams{}

    contents, err := sheetDB.Read.GetContent(&getContentParams)

    ...
}
```

### Examples
For more examples, please check example directory.

### Todo
- [x] Read
- [x] Search
- [ ] Create
- [ ] Update
- [ ] Delete
- [ ] Unit Test