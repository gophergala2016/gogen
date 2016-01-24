package resource

import "github.com/gophergala2016/gogen/generator-model"

// User is entity definition for the user
var User = &model.Model{
	Name: "User",
	Fields: []model.Field{
		{
			Name: "Username",
			Type: model.String,
		},
		{
			Name: "Password",
			Type: model.String,
		},
		{
			Name: "Email",
			Type: model.String,
		},
	},
}
