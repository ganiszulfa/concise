package gql

import (
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
