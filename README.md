# go-defectdojo

go-defectdojo is a Go client library for accessing the [DefectDojo API](https://defectdojo.github.io/django-DefectDojo/integrations/api-v2-docs/)

## Requirements ##

- Go version 1.17 has been used for the initial development
- Go mod is used for dependency management
- The latest version of DefectDojo the APIs are have been tested with is `v2.6.2`

## Usage ##

```go
import "github.com/truemilk/go-defectdojo/defectdojo"
```

Construct a new Defectdojo client, then use the various methods to access the DefectDojo API.

```go
client := &http.Client{
    Timeout: time.Minute,
    Transport: &http.Transport{
        Proxy: http.ProxyFromEnvironment,
    },
}

dj, err := defectdojo.NewDojoClient(url, token, client)
```

### Authentication ###

The go-defectdojo library handles authentication via Token. You can retrieve a valid `API v2 Key` from within your DefectDojo instance.

```go
package main

import (
  "context"
  "encoding/json"
  "fmt"
  "os"

  "github.com/truemilk/go-defectdojo/defectdojo"
)

func main() {
  url := os.Getenv("DOJO_URI")
  token := os.Getenv("DOJO_APIKEY")

  dj, err := defectdojo.NewDojoClient(url, token, nil)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  ctx := context.Background()

  resp, err := dj.UserProfileList(ctx)
  if err != nil {
    fmt.Println("main.go:", err)
    return
  }

  b, err := json.Marshal(resp)
  if err != nil {
    fmt.Println("main.go:", err)
    return
  }

  fmt.Println(string(b))
}
```

## Roadmap ##

This library is being initially developed for personal use, so API methods will likely be implemented in the order that they are needed. Eventually, it would be ideal to cover the entire DefectDojo API, so contributions are of course always welcome. The calling pattern is pretty well established, so adding new methods is relatively straightforward.

## License ##

MIT licensed, see [LICENSE][LICENSE] file.

[LICENSE]: ./LICENSE
