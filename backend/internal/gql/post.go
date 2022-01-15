package gql

import (
	"github.com/ganiszulfa/concise/backend/internal/posts"
	"github.com/graphql-go/graphql"
)

func (g *G) InitializePost() {
	mutationFieldsList = append(mutationFieldsList, postMutationFields)
	queryFieldsList = append(queryFieldsList, postQueryFields)
}

var PostType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Post",
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

var postQueryFields = graphql.Fields{
	"GetPost": &graphql.Field{
		Type:        PostType,
		Description: "Get post by slug",
		Args: graphql.FieldConfigArgument{
			"slug": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return posts.GetBySlug(p.Context, p.Args)
		},
	},
	"ListPost": &graphql.Field{
		Type:        graphql.NewList(PostType),
		Description: "Get list of posts",
		Args: graphql.FieldConfigArgument{
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return posts.GetList(p.Context, p.Args)
		},
	},
}

var postMutationFields = graphql.Fields{
	"CreatePost": &graphql.Field{
		Type:        PostType,
		Description: "Create new post",
		Args: graphql.FieldConfigArgument{
			"title": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"content": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return posts.Create(p.Context, p.Args)
		},
	},

	"UpdatePost": &graphql.Field{
		Type:        PostType,
		Description: "Update post by slug",
		Args: graphql.FieldConfigArgument{
			"slug": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"title": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"content": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return posts.Update(p.Context, p.Args)
		},
	},

	"DeletePost": &graphql.Field{
		Type:        PostType,
		Description: "Delete post by slug",
		Args: graphql.FieldConfigArgument{
			"slug": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return posts.Delete(p.Context, p.Args)
		},
	},
}
