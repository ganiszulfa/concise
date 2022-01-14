package gql

import (
	"github.com/ganiszulfa/concise/internal/users"
	"github.com/graphql-go/graphql"
)

func (g *G) InitializeUser() {
	mutationFieldsList = append(mutationFieldsList, userMutationFields)
	queryFieldsList = append(queryFieldsList, userQueryFields)
}

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"firstName": &graphql.Field{
				Type: graphql.String,
			},
			"lastName": &graphql.Field{
				Type: graphql.String,
			},
			"isOwner": &graphql.Field{
				Type: graphql.Boolean,
			},
			"token": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var userQueryFields = graphql.Fields{
	"GetUser": &graphql.Field{
		Type:        UserType,
		Description: "Get user by email",
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return users.GetByEmail(p.Context, p.Args)
		},
	},
	"ListUsers": &graphql.Field{
		Type:        graphql.NewList(UserType),
		Description: "Get list of users",
		Args: graphql.FieldConfigArgument{
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return users.GetList(p.Context, p.Args)
		},
	},
	"UserLogin": &graphql.Field{
		Type:        UserType,
		Description: "login with email and password (will get token)",
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return users.Login(p.Context, p.Args)
		},
	},
}

var userMutationFields = graphql.Fields{
	"CreateUser": &graphql.Field{
		Type:        UserType,
		Description: "Create new user",
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"firstName": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"lastName": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"isOwner": &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
			"ownerPassword": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return users.Create(p.Context, p.Args)
		},
	},

	"UpdateUser": &graphql.Field{
		Type:        UserType,
		Description: "Update user by email",
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"firstName": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"lastName": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return users.Update(p.Context, p.Args)
		},
	},
}
