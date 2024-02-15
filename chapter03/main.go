package main

import (
	"flag"
	"log"

	"github.com/ibiscum/Test-Driven-Development-in-Go/chapter03/calculator"
	"github.com/ibiscum/Test-Driven-Development-in-Go/chapter03/input"
)

func main() {
	expr := flag.String("expression", "", "mathematical expression to parse")
	flag.Parse()

	engine := calculator.NewEngine()
	validator := input.NewValidator(engine.GetNumOperands(), engine.GetValidOperators())
	parser := input.NewParser(engine, validator)
	result, err := parser.ProcessExpression(*expr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(*result)
}
