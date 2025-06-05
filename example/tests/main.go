// tests demonstrate how to list tests with different filters to DefectDojo.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/truemilk/go-defectdojo/defectdojo"
)

func main() {
	uri := os.Getenv("DOJO_URI")
	token := os.Getenv("DOJO_APIKEY")

	dj, err := defectdojo.NewDojoClient(uri, token, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ctx := context.Background()

	opts := &defectdojo.TestsOptions{
		Limit: 10,
		Engagement: 1,
	}
	resp, err := dj.Tests.List(ctx, opts)
	fmt.Println("Test filter Engagement")
	if err != nil {
		fmt.Println("main Engagement:", err)
	} else {
		displayTests(resp)
	}
	
	opts = &defectdojo.TestsOptions{
		Limit: 10,
		TestType: 1,
	}
	resp, err = dj.Tests.List(ctx, opts)
	fmt.Println("Test filter TestType")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		displayTests(resp)
	}
	
	opts = &defectdojo.TestsOptions{
		Limit: 10,
		Title: url.QueryEscape("Title Test"),
	}
	resp, err = dj.Tests.List(ctx, opts)
	fmt.Println("Test filter Title")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		displayTests(resp)
	}
}


func displayTests(resp *defectdojo.Tests) {
	b, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("main:", err)
		return
	}

	fmt.Println(string(b))
}