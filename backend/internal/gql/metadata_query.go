package gql

import (
	"github.com/ganiszulfa/concise/backend/internal/metadata"
	"github.com/graphql-go/graphql"
)

var mdQueryFields = graphql.Fields{
	"GetMetadata": &graphql.Field{
		Type:        MetadataType,
		Description: "Get metadata by key",
		Args: graphql.FieldConfigArgument{
			"key": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return metadata.GetByKey(p.Context, p.Args)
		},
	},
	"ListMetadata": &graphql.Field{
		Type:        graphql.NewList(MetadataType),
		Description: "Get list of metadata",
		Args: graphql.FieldConfigArgument{
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return metadata.GetList(p.Context, p.Args)
		},
	},
}
