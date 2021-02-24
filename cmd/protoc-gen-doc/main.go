package main

import (
	protodoc "github.com/aspage2/protoc-doc"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(protodoc.RealMain)
}
