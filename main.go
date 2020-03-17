package main

import (
	"flag"
	"fmt"

	"github.com/dave/jennifer/jen"
)

const (
	Name int = 0
)

func main() {
	flag.Parse()
	name := flag.Arg(Name)
	switch name {
	case "main":
		Main()
	default:
	}
}

func Main() {
	f := jen.NewFile("main")
	f.Func().Id("main").Params().Block()
	fmt.Printf("%#v", f)
}
