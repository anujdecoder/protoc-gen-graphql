package main

import (
	"github.com/lyft/protoc-gen-star"
	"github.com/srikrsna/protoc-gen-graphql/module"
)

func main() {
	pgs.
		Init().
		RegisterModule(module.New()).
		RegisterPostProcessor(pgs.GoFmt()).
		Render()
}
