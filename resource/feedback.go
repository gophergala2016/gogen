package resource

import "github.com/gophergala2016/gogen"

// Feedback is model for the feedback entity
var Feedback = &gogen.Model{
	Name: "Feedback",
	Fields: []gogen.Field{
		{
			Name: "Description",
			Type: gogen.String,
		},
	},
}
