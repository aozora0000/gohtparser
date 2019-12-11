package main

import (
	"github.com/aozora0000/gohtparser"
	"os"
)

func main() {
	fd, err := os.Open(os.Args[1])
	if err != nil {
		panic(err.Error())
	}

	parser := gohtparser.NewParser(fd)
	ast, err := parser.GenerateAst()
	if err != nil {
		panic(err.Error())
	}
	ast.Dump()
	//for str := range ast.ToHtAccess() {
	//	fmt.Println(str)
	//}
}
