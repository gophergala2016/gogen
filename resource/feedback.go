package resource

import "github.com/gophergala2016/gogen/generator-model"

// Feedback is model for the feedback entity
var Feedback = &model.Model{
	Name: "Feedback",
	Fields: []model.Field{
		{
			Name: "Description",
			Type: model.String,
		},
	},
}
