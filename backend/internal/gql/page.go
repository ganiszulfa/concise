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
			"isPublished": &graphql.Field{
				Type: graphql.Boolean,
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
		Description: "Get page by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return pages.GetById(p.Context, p.Args)
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
			"isPublished": &graphql.ArgumentConfig{
				Type: graphql.Boolean,
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
			"isPublished": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Boolean),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return pages.Create(p.Context, p.Args)
		},
	},

	"UpdatePage": &graphql.Field{
		Type:        PageType,
		Description: "Update page by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"title": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"content": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"order": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"isPublished": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Boolean),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return pages.Update(p.Context, p.Args)
		},
	},

	"DeletePage": &graphql.Field{
		Type:        PageType,
		Description: "Delete page by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return pages.Delete(p.Context, p.Args)
		},
	},
}
