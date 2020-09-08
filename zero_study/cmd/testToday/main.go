package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", "package main; var a = 3", parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	var v visitor
	ast.Walk(v, f)
}

type visitor struct{}

func (v visitor) Visit(n ast.Node) ast.Visitor {
	fmt.Printf("%T\n", n)
	return v
}

//{
//	fs := token.NewFileSet()
//	//f, err := parser.ParseFile(fs, "", "package main; var a = 3", parser.AllErrors)
//	f, err := parser.ParseFile(fs, "", "package main; var a = 3", parser.ImportsOnly)
//	//f, err := parser.ParseFile(fs, "", "package main; var a = 3", parser.Trace)
//	if err != nil {
//		log.Fatal(err)
//	}
//	printer.Fprint(os.Stdout, fs, f)
//}
