// token demonstrate the steps necessary to retrieve an API token programmatically
// from a DefectDojo instance with username and password.
//
// For demonstration purposes, it uses the public demo instance available at:
// https://demo.defectdojo.org/. The test credentials are made available by the
// DefectDojo creators here: https://github.com/DefectDojo/django-DefectDojo#demo.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/truemilk/go-defectdojo/defectdojo"
)

func main() {
	url := "https://demo.defectdojo.org"

	dj, err := defectdojo.NewDojoClient(url, "", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ctx := context.Background()

	resp, err := dj.ApiTokenAuth.Create(ctx, &defectdojo.AuthToken{
		Username: defectdojo.String("admin"),
		Password: defectdojo.String("1Defectdojo@demo#appsec"),
	})
	if err != nil {
		fmt.Println("main.go:", err)
		return
	}

	b, err := json.Marshal(resp.Token)
	if err != nil {
		fmt.Println("main.go:", err)
		return
	}

	fmt.Println(string(b))
}
