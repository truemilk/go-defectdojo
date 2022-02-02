// custom demonstrates how to specify a custom HTTP transport and a custom Context with the requests.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/truemilk/go-defectdojo/defectdojo"
)

func main() {
	url := os.Getenv("DOJO_URI")
	token := os.Getenv("DOJO_APIKEY")

	client := &http.Client{
		Timeout: time.Minute,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}

	dj, err := defectdojo.NewDojoClient(url, token, client)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := dj.Users.List(ctx, nil)
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
