// Package main is the entry point for gitx.
package main

import (
	"os"

	"github.com/oxodx/x/gitx/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
