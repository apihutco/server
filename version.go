package main

import (
	"fmt"
)

var (
	version   string
	buildTime string
)

func init() {
	fmt.Printf("Verson: %s\nBuilt: %s\n", version, buildTime)
}
