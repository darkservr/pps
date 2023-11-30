package main

import (
	_ "embed"

	"github.com/Pawswap/posa/command/root"
	"github.com/Pawswap/posa/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
