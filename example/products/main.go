// findings demonstrate how to retrieve the findings from DefectDojo.
// Options are defined to specify the size of the page retrieved and the offset.
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

	prod := &defectdojo.Product{
		Name:        defectdojo.Str("fancyOne"),
		Description: defectdojo.Str("fancyOne"),
		Tags:        defectdojo.Slice([]string{"fancy1", "fancy2"}),
		ProdType:    defectdojo.Int(14),
	}
	resp, err := dj.Products.Create(ctx, prod)
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
