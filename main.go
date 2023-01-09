package main

import (
	_ "embed"

	"github.com/gitshock-labs/go-client/command/root"
	"github.com/gitshock-labs/go-client/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
