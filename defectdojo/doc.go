/*
Package defectdojo provides a client for using the DefectDojo API.

Installation

Import the module in your source code files:

	import "github.com/truemilk/go-defectdojo/defectdojo"

Usage

Define a new Defectdojo client:

	url := os.Getenv("DOJO_URI")
	token := os.Getenv("DOJO_APIKEY")

	dj, err := defectdojo.NewDojoClient(url, token, nil)

It is possible to specify a custom HTTP Transport when creating the client:

	url := os.Getenv("DOJO_URI")
	token := os.Getenv("DOJO_APIKEY")

	client := &http.Client{
		Timeout: time.Minute,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}

	dj, err := defectdojo.NewDojoClient(url, token, client)

Then use the various methods to access the DefectDojo API.

The services of a client divide the API into logical chunks and correspond to the structure of the DefectDojo API documentation.

Auxiliary methods have are provided to facilitate the definition of parameters of Bool/String/Int/Int64 type.
These should be used when passing parameters to methods that create/update/delete objects.

	ctx := context.Background()

	resp, err := dj.Technologies.Create(ctx, &defectdojo.Technology{
		Name:         defectdojo.String("A new technology"),
		Product:      defectdojo.Int(1),
		User:         defectdojo.Int(1),
	})

One exception to the pattern above is made for "import-scan", where a map of strings is used for all parameters:

	params := &defectdojo.ImportScanMap{
		"scan_type":           "Trivy Scan",
		"engagement_name":     "Today's Engagement",
		"product_name":        "A Secure Product",
		"product_type_name":   "Core Infrastructure",
		"auto_create_context": "true",
		"file":                "/tmp/trivy-output.json",
	}

	resp, err := dj.ImportScan.Create(ctx, params)

NOTE: Using the context package, one can easily pass cancellation signals and deadlines to various services of the client for handling a request.
In case there is no context available, then context.Background() can be used as a starting point.

Authentication

The go-defectdojo library handles authentication via Token. You can retrieve a valid API v2 Key from within your DefectDojo instance.

It is also possible to retrieve the API key from an un-authenticated call to the "/api-token-auth/" endpoint, specifying valid username and password.
For the purpose of this API call, the client can be instantiated with an empty string as the `token` parameter.

	dj, _ := defectdojo.NewDojoClient(url, "", nil)

	resp, err := dj.ApiTokenAuth.Create(ctx, &defectdojo.AuthToken{
		Username: defectdojo.String("username"),
		Password: defectdojo.String("password"),
	})

	fmt.Println(string(*resp.Token))

The token can be later used to instantiate the client again for further authenticated API calls.
*/
package defectdojo
