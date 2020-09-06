package query

import (
	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.ObjectConfig{
	Name: "RootQueryType",
	Fields: graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "World", nil
			},
		},
	},
}
