package main

import (
	_ "embed"
	"fmt"
)

//go:embed config.json
var ndung string

//go:embed anh.jpg
var data []byte

func main() {
	fmt.Println(ndung)
	// fmt.Println(data)
}
