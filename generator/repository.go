package generator

import "github.com/gophergala2016/gogen"

var (
	// Repository is global registration of the generator
	Repository = &RepositoryGenerator{}
)

// RepositoryGenerator encapsulates the logic behind
// generating of models
type RepositoryGenerator struct {
	gogen.GeneratorContext
}

// Generate will call the generator to generate
// results
func (g *RepositoryGenerator) Generate() error {
	return nil
}
