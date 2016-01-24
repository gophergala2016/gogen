package model

import (
	"bytes"
	"text/template"

	"github.com/gophergala2016/gogen"
	"github.com/op/go-logging"
)

var (
	// Generator is global registration of the generator
	Generator = &generator{}

	genlog = logging.MustGetLogger("gogen")
)

// ModelGenerator encapsulates the logic behind
// generating of models
type generator struct {
	gogen.GeneratorContext
}

// Name returns name of the generator
func (g *generator) Name() string {
	return "ModelGenerator"
}

// Generate will call the generator to generate
// results
func (g *generator) Generate() error {
	err := g.Prepare()
	if err != nil {
		return err
	}

	// compile package template
	packTmpl, err := template.New("package").Parse(packageTemplate)
	if err != nil {
		return err
	}

	// compile model template
	tmpl, err := template.New("model").Parse(modelTemplate)
	if err != nil {
		return err
	}

	for _, resource := range *g.Resources {
		if model, ok := resource.(*Model); ok {
			genlog.Info("Generating model for %s", model.Name)
			content := bytes.Buffer{}
			packTmpl.Execute(&content, g)
			tmpl.Execute(&content, model)
			g.SaveFile(model.Name, content)
		}
	}

	return nil
}

// Templates
var (
	packageTemplate = `package {{.PackageName}}`

	modelTemplate = `
		//  {{.Name}} is model representing the entity
		type {{.Name}} struct {
		  {{range .Fields}}{{.Name}} {{.Type.Name}}
		  {{end}}
		}`
)
