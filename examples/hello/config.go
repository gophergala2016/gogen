package main

import (
	"github.com/gophergala2016/gogen"
	"github.com/gophergala2016/gogen/generator"
	"github.com/gophergala2016/gogen/resource"
)

func main() {
	// Define resources that should be available to
	// generators.
	gogen.Define(resource.User)

	// set output for the generator
	generator.Model.SetOutputDir("./model")

	// Define pipes that should be run
	gogen.Pipe(
		generator.Model,
	)

	// start the generator
	if err := gogen.Generate(); err != nil {
		panic(err)
	}
}
