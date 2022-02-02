# go-defectdojo

go-defectdojo is a Go client library for accessing the [DefectDojo API](https://defectdojo.github.io/django-DefectDojo/integrations/api-v2-docs/)

## Requirements ##

- Minimum tested Go version is `1.17`
- The latest version of DefectDojo the APIs are have been tested with is `v2.6.2`

## Basic Usage ##

Import the module in your source code:

```go
import "github.com/truemilk/go-defectdojo/defectdojo"
```

Create a client:

```go
url := os.Getenv("DOJO_URI")
token := os.Getenv("DOJO_APIKEY")

client := &http.Client{
        Timeout: time.Minute,
        Transport: &http.Transport{
          Proxy: http.ProxyFromEnvironment,
        },
    }

dj, err := defectdojo.NewDojoClient(url, token, client)
```

Reference the provided methods to call the API:

```go
ctx := context.Background()

opts := &defectdojo.FindingsOptions{
    Limit:    20,
    Offset:   5,
    Prefetch: "duplicate_finding",
}

resp, err := dj.Findings.List(ctx, opts)
```

More detailed documentation is available at: https://pkg.go.dev/github.com/truemilk/go-defectdojo/defectdojo

For additional usage examples, brouse the [example](example) folder.

## Roadmap ##

This library is being initially developed for personal use, so API methods will likely be implemented in the order that they are needed.
Eventually, it would be ideal to cover the entire DefectDojo API, so contributions are of course always welcome.
The calling pattern is pretty well established, so adding new methods is relatively straightforward.

## License ##

MIT licensed, see [LICENSE][LICENSE] file.

[LICENSE]: ./LICENSE
