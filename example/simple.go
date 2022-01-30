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

	ctx := context.Background()

	params := &defectdojo.ImportScanMap{
		"scan_type":           "Trivy Scan",
		"engagement_name":     "New Engagement",
		"product_name":        "New Product",
		"product_type_name":   "New Product Type",
		"auto_create_context": "true",
		"file":                "/tmp/trivy-out.json",
	}
	_, err = dj.ImportScan.Create(ctx, params)
	if err != nil {
		fmt.Println("main.go:", err)
		return
	}

	opts := &defectdojo.FindingsOptions{
		Limit:    5,
		Offset:   1,
		Prefetch: "duplicate_finding",
	}
	resp, err := dj.Findings.List(ctx, opts)
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
