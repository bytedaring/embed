// version_prod.go
// +build !dev

package main

import (
	_ "embed"
)

var (
	//go:embed version.txt
	version string
)
