package main

import (
	"fmt"
	"github.com/noChaos1012/tour/zero_study/cmd/parse/parser"
	"log"
)

var apiFile = "demo.api"

func main() {

	parse, err := parser.NewParser(apiFile)
	if err != nil {
		log.Fatal(err)
	}

	api , err := parse.Parse()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("api.Service is :%v\n", api.Service)
}
