package generator

import (
	"bytes"
	"text/template"

	"github.com/gophergala2016/gogen"
	"github.com/gophergala2016/gogen/generator/repositorytmpl"
)

var (
	// Repository is global registration of the generator
	Repository = &RepositoryGenerator{}
)

// Repository types that can be used
const (
	MongoRepository = iota
	PostgresRepository
	RedisRepository
)

// RepositoryGenerator encapsulates the logic behind
// generating of models
type RepositoryGenerator struct {
	gogen.GeneratorContext

	repositoryType int
}

// SetRepositoryType will set the type of the generated
// repository. Defaults to Mongo
func (g *RepositoryGenerator) SetRepositoryType(t int) {
	g.repositoryType = t
}

// Generate will call the generator to generate
// results
func (g *RepositoryGenerator) Generate() error {
	err := g.Prepare()
	if err != nil {
		return err
	}

	// temporary template variable
	var templ string
	switch g.repositoryType {
	case MongoRepository:
		templ = repositorytmpl.MongoRepositoryTemplate
	}

	// compile mongo repository
	repoTmpl, err := template.New("repository").Parse(templ)
	if err != nil {
		return err
	}

	for _, model := range gogen.Models {
		content := bytes.Buffer{}
		repoTmpl.Execute(&content,
			struct {
				*gogen.Model
				PackageName string
			}{
				Model:       model,
				PackageName: g.PackageName(),
			},
		)
		g.SaveFile(model.Name+"Repository", content)
	}

	return nil
}
