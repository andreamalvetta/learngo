package main

import (
	"fmt"
	"path"
)

func main() {

	// var dir, file string
	// var file string

	// dir, file = path.Split("css/main.css")

	// if you don't want to use one var put _
	// _, file = path.Split("css/main.css")

	// You can actually use a short declaration statement
	_, file := path.Split("css/main.css")

	// fmt.Println("dir:", dir)
	fmt.Println("file:", file)
}
