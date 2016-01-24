package repository

import (
	"bytes"
	"text/template"

	"github.com/gophergala2016/gogen"
	"github.com/gophergala2016/gogen/generator-model"
	"github.com/gophergala2016/gogen/generator-repository/tmpl"
	"github.com/op/go-logging"
)

var (
	// Generator is global registration of the generator
	Generator = &generator{}

	genlog = logging.MustGetLogger("gogen")
)

// Repository types that can be used
const (
	Mongo = iota
	Postgres
	Redis
)

// generator encapsulates the logic behind
// generating of models
type generator struct {
	gogen.GeneratorContext

	repositoryType int
}

// Name returns name of the generator
func (g *generator) Name() string {
	return "RepositoryGenerator"
}

// SetRepositoryType will set the type of the generated
// repository. Defaults to Mongo
func (g *generator) SetRepositoryType(t int) {
	g.repositoryType = t
}

// Generate will call the generator to generate
// results
func (g *generator) Generate() error {
	err := g.Prepare()
	if err != nil {
		return err
	}

	// temporary template variable
	var templ string
	switch g.repositoryType {
	case Mongo:
		templ = repositorytmpl.MongoRepositoryTemplate
	}

	// compile mongo repository
	repoTmpl, err := template.New("repository").Parse(templ)
	if err != nil {
		return err
	}

	for _, resource := range *g.Resources {
		if entity, ok := resource.(*model.Model); ok {
			genlog.Info("Generating repository for model %s", entity.Name)
			content := bytes.Buffer{}
			err = repoTmpl.Execute(&content,
				struct {
					Model       *model.Model
					PackageName string
				}{
					Model:       entity,
					PackageName: g.PackageName(),
				},
			)
			if err != nil {
				return err
			}
			g.SaveFile(entity.Name+"Repository", content)
		}
	}

	return nil
}
