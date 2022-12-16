package gql

import (
	"github.com/ganiszulfa/concise/backend/internal/controllers"
	"github.com/graphql-go/graphql"
)

var postMutationFields = graphql.Fields{
	controllers.PostMutationCreate: &graphql.Field{
		Type:        PostType,
		Description: "Create a new post",
		Args: graphql.FieldConfigArgument{
			controllers.ArgsPostsTitle: &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			controllers.ArgsPostsContent: &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			controllers.ArgsPostsIsPublished: &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Boolean),
			},
			controllers.ArgsPostsIsPage: &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Boolean),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return postCtr.CreateFromGQL(p.Context, p.Args)
		},
	},

	controllers.PostMutationUpdate: &graphql.Field{
		Type:        PostType,
		Description: "Update a post by slug",
		Args: graphql.FieldConfigArgument{
			controllers.ArgsPostsSlug: &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			controllers.ArgsPostsTitle: &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			controllers.ArgsPostsContent: &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			controllers.ArgsPostsIsPublished: &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Boolean),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return postCtr.UpdateFromGQL(p.Context, p.Args)
		},
	},

	controllers.PostMutationDelete: &graphql.Field{
		Type:        PostType,
		Description: "Delete post by slug",
		Args: graphql.FieldConfigArgument{
			controllers.ArgsPostsSlug: &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return nil, postCtr.DeleteFromGQL(p.Context, p.Args)
		},
	},
}
