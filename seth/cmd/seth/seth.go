package main

import (
	"os"

	seth "github.com/goplugin/plugin-testing-framework/seth/cmd"
)

func main() {
	if err := seth.RunCLI(os.Args); err != nil {
		panic(err)
	}
}
