package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/anup-gupta1/go-graphql-mongodb/query"
)

var SchemaConfig = graphql.SchemaConfig{
	Query: graphql.NewObject(query.RootQuery)
}
