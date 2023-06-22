package main

import (
	_ "embed"

	"github.com/mauricio-b-r/polygon-edge/command/root"
	"github.com/mauricio-b-r/polygon-edge/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
