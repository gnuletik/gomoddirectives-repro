package main

import (
	"fmt"

	"example.com/pkg"
)

func main() {
	fmt.Println(pkg.Get("Hello"))
}
