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

Define a new Defectdojo client:

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

Then use the various methods to access the DefectDojo API.

```go
ctx := context.Background()

resp, err := dj.Technologies.Create(ctx, &defectdojo.Technology{
    Name:         defectdojo.String("A new technology"),
    Product:      defectdojo.Int(1),
    User:         defectdojo.Int(1),
})
```

One exception to the pattern above is made for `import-scan`:

```go
params := &defectdojo.ImportScanMap{
        "scan_type":           "Trivy Scan",
        "engagement_name":     "Today's Engagement",
        "product_name":        "A Secure Product",
        "product_type_name":   "Core Infrastructure",
        "auto_create_context": "true",
        "file":                "/tmp/trivy-output.json",
    }

resp, err := dj.ImportScan.Create(ctx, params)
```

More detailed documentation is available at: https://pkg.go.dev/github.com/truemilk/go-defectdojo/defectdojo

### Authentication ###

The go-defectdojo library handles authentication via Token. You can retrieve a valid _API v2 Key_ from within your DefectDojo instance.

It is also possible to retrieve the API key from an un-authenticated call to the `/api-token-auth/` endpoint, specifying valid username and password.
For the purpose of this API call, the client can be instantiated with an empty string as the `token` parameter.

```go
dj, _ := defectdojo.NewDojoClient(url, "", nil)

resp, err := dj.ApiTokenAuth.Create(ctx, &defectdojo.AuthToken{
    Username: defectdojo.String("username"),
    Password: defectdojo.String("password"),
})

fmt.Println(string(*resp.Token))
```

The token can be later used to instantiate the client again for further authenticated API calls.


## Roadmap ##

This library is being initially developed for personal use, so API methods will likely be implemented in the order that they are needed.
Eventually, it would be ideal to cover the entire DefectDojo API, so contributions are of course always welcome.
The calling pattern is pretty well established, so adding new methods is relatively straightforward.

## License ##

MIT licensed, see [LICENSE][LICENSE] file.

[LICENSE]: ./LICENSE
