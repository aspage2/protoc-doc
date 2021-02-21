package main

import (
	protodoc "doc"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(protodoc.RealMain)
}
