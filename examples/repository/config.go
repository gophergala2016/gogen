package main

import (
	"github.com/gophergala2016/gogen"
	"github.com/gophergala2016/gogen/generator-model"
	"github.com/gophergala2016/gogen/generator-repository"
	"github.com/gophergala2016/gogen/resource"
)

func main() {
	// Define resources that should be available to
	// generators.
	gogen.Define(resource.User)

	// set output for the generator
	model.Generator.SetOutputDir("./model")
	repository.Generator.SetOutputDir("./repository")

	// Define pipes that should be run
	gogen.Pipe(
		model.Generator,
		repository.Generator,
	)

	// start the generator
	if err := gogen.Generate(); err != nil {
		panic(err)
	}
}
