package main

import (
	"fmt"

	"github.com/bufbuild/protovalidate-go"
	"k8s.io/apiserver/pkg/authentication/cel"
)

func main() {
	fmt.Println(cel.NewDefaultCompiler())
	fmt.Println(protovalidate.Validate(nil))
}
