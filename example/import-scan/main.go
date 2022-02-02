// import scan demonstrates the process of uploading a scan report into the DefectDojo platform.
//
// Reports to import are defined by an ImportScanMap structure that defines all parameters in string format.
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

	params := &defectdojo.ImportScanMap{
		"scan_type":           "Trivy Scan",
		"engagement_name":     "New Engagement",
		"product_name":        "New Product",
		"product_type_name":   "New Product Type",
		"auto_create_context": "true",
		"file":                "/tmp/trivy.json",
	}
	resp, err := dj.ImportScan.Create(ctx, params)
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
