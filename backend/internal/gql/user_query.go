package gql

import (
	"github.com/ganiszulfa/concise/backend/internal/controllers"
	"github.com/graphql-go/graphql"
)

var userQueryFields = graphql.Fields{
	controllers.UserQueryLogin: &graphql.Field{
		Type:        UserLoginType,
		Description: "Login and get session id",
		Args: graphql.FieldConfigArgument{
			controllers.ArgsUserId: &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			controllers.ArgsUserPassword: &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return userCtr.UserLoginFromGQL(p.Context, p.Args)
		},
	},
}
