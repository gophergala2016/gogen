package generator

import "github.com/gophergala2016/gogen"

var (
	// API is global registration of the generator
	API = &APIGenerator{}
)

// APIGenerator encapsulates the logic behind
// generating of models
type APIGenerator struct {
	gogen.GeneratorContext
}

// Generate will call the generator to generate
// results
func (g *APIGenerator) Generate() error {
	return nil
}
