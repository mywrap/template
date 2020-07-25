package main

import (
	"fmt"

	"github.com/mywrap/template"
)

func main() {
	s := template.DemoFunc0(1, 2)
	fmt.Println("sum:", s)
}
