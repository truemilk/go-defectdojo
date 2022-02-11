// importscan demonstrates the process of uploading a scan report into DefectDojo.
//
// Details of the import are defined by an ImportScan struct.
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

	scan := &defectdojo.ImportScan{
		ProductTypeName:   defectdojo.Str("Hello1"),
		ProductName:       defectdojo.Str("Hello1"),
		EngagementName:    defectdojo.Str("Hello1"),
		AutoCreateContext: defectdojo.Bool(true),
		File:              defectdojo.Str("/tmp/trivy.json"),
		ScanType:          defectdojo.Str("Trivy Scan"),
		Tags:              defectdojo.Slice([]string{"AAAA", "BBBB"}),
	}

	resp, err := dj.ImportScan.Create(ctx, scan)
	if err != nil {
		fmt.Println("main:", err)
		return
	}

	b, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("main:", err)
		return
	}

	fmt.Println(string(b))
}
