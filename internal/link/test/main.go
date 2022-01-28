package main

import (
	"fmt"

	"github.com/mcuv3/demo/internal/link"
)

func main() {
	test := link.Encode(1234)
	fmt.Println(test)

	fmt.Println(link.Decode("4t"))
}
