package gql

import (
	"github.com/ganiszulfa/concise/backend/config/app"
	"github.com/ganiszulfa/concise/backend/internal/controllers"
	"github.com/ganiszulfa/concise/backend/internal/repos"
	"github.com/ganiszulfa/concise/backend/internal/usecases"
	"github.com/graphql-go/graphql"
)

var userCtr controllers.UserCtrInterface

func (g *G) InitializeUser() {

	mr := repos.NewMetadataRepo(app.DB)
	sr := repos.NewSessionRepo(app.DB)
	uu := usecases.NewUserUc(mr, sr)

	userCtr = controllers.NewUserCtr(uu)

	queryFieldsList = append(queryFieldsList, userQueryFields)
}

var UserLoginType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: controllers.ObjectNameUserLogin,
		Fields: graphql.Fields{
			controllers.ArgsUserSession: &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
