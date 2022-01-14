package gql

import (
	"github.com/ganiszulfa/concise/backend/internal/metadata"
	"github.com/graphql-go/graphql"
)

func (g *G) InitializeMetadata() {
	mutationFieldsList = append(mutationFieldsList, mdMutationFields)
	queryFieldsList = append(queryFieldsList, mdQueryFields)
}

var MetadataType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Metadata",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"createdAt": &graphql.Field{
				Type: graphql.DateTime,
			},
			"updatedAt": &graphql.Field{
				Type: graphql.DateTime,
			},
			"key": &graphql.Field{
				Type: graphql.String,
			},
			"value": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
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
