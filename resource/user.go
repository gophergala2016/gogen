package resource

import "github.com/gophergala2016/gogen"

// User is entity definition for the user
var User = &gogen.Model{
	Name: "User",
	Fields: []gogen.Field{
		{
			Name: "Username",
			Type: gogen.String,
		},
		{
			Name: "Password",
			Type: gogen.String,
		},
		{
			Name: "Email",
			Type: gogen.String,
		},
	},
}
