package gogen

import (
	"os"
	"sync"

	"github.com/op/go-logging"
)

var (
	// Models is public static set of models exposed
	// by the Gogen, so generators can use it. This
	// set contains every model that was added either
	// manually or by Define function
	// @Deprecated
	//Models []*Model

	// Resources is set resources that were firstly defined
	Resources ResourceContainer

	// Pipes is set of pipelines that should be run when
	// generate is called
	Pipes []Pipeline
)

var (
	genlog = logging.MustGetLogger("gogen")

	format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
)

// initialize logging
func init() {
	backend := logging.NewLogBackend(os.Stdout, "", 0)
	formatter := logging.NewBackendFormatter(backend, format)

	logging.SetBackend(formatter)
}

// Define will store the defined model for the use in
// the generators.
func Define(resource interface{}) {
	Resources = append(Resources, resource)
}

// Pipe will register new pipe that will be run
// in parallel
func Pipe(gens ...Generator) {
	pipe := Pipeline{}
	for _, gen := range gens {
		pipe.Add(gen)
	}

	Pipes = append(Pipes, pipe)
}

// Generate will startup a
func Generate() error {
	genlog.Info("Starting gogen")

	wg := sync.WaitGroup{}

	for pipeindex, pipe := range Pipes {
		wg.Add(1)
		go func(pipe Pipeline, pipeindex int) {
			for _, gen := range pipe.generators {
				genlog.Info("Starting generator %s in pipe %d", gen.Name(), pipeindex)
				gen.Initialize(&Resources)

				err := gen.Generate()
				genlog.Info("End of generator %s in pipe %d", gen.Name(), pipeindex)
				// TODO: make this not panic, but return the error
				if err != nil {
					panic(err)
				}
			}

			wg.Done()
		}(pipe, pipeindex)
	}

	wg.Wait()

	return nil
}
