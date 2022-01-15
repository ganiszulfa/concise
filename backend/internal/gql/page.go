package gql

import (
	"github.com/ganiszulfa/concise/backend/internal/pages"
	"github.com/graphql-go/graphql"
)

func (g *G) InitializePage() {
	mutationFieldsList = append(mutationFieldsList, pageMutationFields)
	queryFieldsList = append(queryFieldsList, pageQueryFields)
}

var PageType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Page",
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
			"order": &graphql.Field{
				Type: graphql.Int,
			},
			"slug": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"content": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: UserType,
			},
		},
	},
)

var pageQueryFields = graphql.Fields{
	"GetPage": &graphql.Field{
		Type:        PageType,
		Description: "Get page by slug",
		Args: graphql.FieldConfigArgument{
			"slug": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return pages.GetBySlug(p.Context, p.Args)
		},
	},
	"ListPage": &graphql.Field{
		Type:        graphql.NewList(PageType),
		Description: "Get list of pages",
		Args: graphql.FieldConfigArgument{
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return pages.GetList(p.Context, p.Args)
		},
	},
}

var pageMutationFields = graphql.Fields{
	"CreatePage": &graphql.Field{
		Type:        PageType,
		Description: "Create new page",
		Args: graphql.FieldConfigArgument{
			"order": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"title": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"content": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return pages.Create(p.Context, p.Args)
		},
	},

	"UpdatePage": &graphql.Field{
		Type:        PageType,
		Description: "Update page by slug",
		Args: graphql.FieldConfigArgument{
			"title": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"content": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return pages.Update(p.Context, p.Args)
		},
	},
}
