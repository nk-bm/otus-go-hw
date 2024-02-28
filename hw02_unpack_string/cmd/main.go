package main

import (
	"fmt"

	hw02unpackstring "github.com/bambanik/otus-go-hw/hw02_unpack_string"
)

func main() {
	res, err := hw02unpackstring.Unpack("qw\\ne")
	fmt.Println(res, err)
}
