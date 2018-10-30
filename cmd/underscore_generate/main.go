package main

import (
	"flag"
	"fmt"
	"log"
)

var schema string
var pkg string
var output string
var typ string

func init() {
	flag.StringVar(&schema, "schema", "", "which type code to generate")
	flag.StringVar(&pkg, "pkg", "", "which package")
	flag.StringVar(&output, "o", "", "filepath to generate")
	flag.StringVar(&typ, "type", "", "whcih type to generate in set schema")
	flag.Parse()
}

func main() {
	switch schema {
	case "set":
		if err := GenerateSet(pkg, typ, output); err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Println(schema, pkg, output, typ)
	}
}
