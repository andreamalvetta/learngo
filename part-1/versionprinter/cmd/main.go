package main

import (
	"fmt"

	"github.com/andreamalvetta/learngo/first/versionprinter"
)

func main() {
	fmt.Println("Go version: " + versionprinter.Version())
}
