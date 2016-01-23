package gogen

// Pipeline is a set of stages that must be run
// in order to get the result
type Pipeline struct {
	generators []Generator
}

// Add will add passed generator into the pipeline
func (p *Pipeline) Add(gen Generator) {
	p.generators = append(p.generators, gen)
}
