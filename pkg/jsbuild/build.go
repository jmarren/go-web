package jsbuild

import (
	"fmt"
	"os"

	"github.com/evanw/esbuild/pkg/api"
)

// builds javascript files into a single bundle
// and writes it to <root>/web/public/index.js
func Build() {
	fmt.Println("js-build")

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", wd)

	result := api.Build(api.BuildOptions{
		EntryPoints: []string{"web/js/dist/app.js"},
		Bundle:      true,
		Write:       true,
		Outfile:     "web/public/index.js",
	})
	fmt.Printf("result: %v\n", result.OutputFiles[0].Path)
	if len(result.Errors) != 0 {
		os.Exit(1)
	}
}
