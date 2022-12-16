package gql

import (
	"github.com/ganiszulfa/concise/backend/internal/controllers"
	"github.com/graphql-go/graphql"
)

var mdQueryFields = graphql.Fields{
	controllers.MetadataQueryList: &graphql.Field{
		Type:        graphql.NewList(MetadataType),
		Description: "Get all metadata",
		Args:        graphql.FieldConfigArgument{},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return metadataCtr.GetAllFromGQL(p.Context, p.Args)
		},
	},
}
