package gql

import (
	"github.com/ganiszulfa/concise/backend/internal/metadata"
	"github.com/graphql-go/graphql"
)

var mdMutationFields = graphql.Fields{
	"CreateMetadata": &graphql.Field{
		Type:        MetadataType,
		Description: "Create new metadata",
		Args: graphql.FieldConfigArgument{
			"key": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"value": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return metadata.Create(p.Context, p.Args)
		},
	},

	"UpdateMetadata": &graphql.Field{
		Type:        MetadataType,
		Description: "Update metadata by key",
		Args: graphql.FieldConfigArgument{
			"key": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"value": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return metadata.Update(p.Context, p.Args)
		},
	},
}
