package gql

import (
	"github.com/ganiszulfa/concise/backend/internal/controllers"
	"github.com/graphql-go/graphql"
)

var postQueryFields = graphql.Fields{
	controllers.PostQueryGet: &graphql.Field{
		Type:        PostType,
		Description: "Get post by slug",
		Args: graphql.FieldConfigArgument{
			controllers.ArgsPostsSlug: &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return postCtr.GetBySlugFromGQL(p.Context, p.Args)
		},
	},
	controllers.PostQueryList: &graphql.Field{
		Type:        graphql.NewList(PostType),
		Description: "Get list of posts",
		Args: graphql.FieldConfigArgument{
			controllers.ArgsLimit: &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			controllers.ArgsPage: &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			controllers.ArgsPostsIsPage: &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
			controllers.ArgsPostsIsPublished: &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return postCtr.GetListFromGQL(p.Context, p.Args)
		},
	},
}
