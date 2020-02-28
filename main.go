package main

import (
	"fmt"
	"strconv"
)

func main() {
	parsed, err := strconv.Atoi("123")
	if err != nil {
		panic(err)
	}

	converted := int32(parsed)
	fmt.Println(converted)
}
