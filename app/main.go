package main

import (
	"fmt"
	"github.com/jakegee/toolkit"
)

func main() {
	var tools toolkit.Tools

	s := tools.RandomString(10)
	fmt.Println("Random String:", s)

}
