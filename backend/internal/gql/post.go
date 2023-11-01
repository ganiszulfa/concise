package gql

import (
	"github.com/ganiszulfa/concise/backend/config/app"
	"github.com/ganiszulfa/concise/backend/internal/controllers"
	"github.com/ganiszulfa/concise/backend/internal/repos"
	"github.com/ganiszulfa/concise/backend/internal/usecases"
	"github.com/graphql-go/graphql"
)

var postCtr controllers.PostCtrInterface

func (g *G) InitializePost() {
	pr := repos.NewPostRepo(app.DB)
	sr := repos.NewSessionRepo(app.DB)
	mr := repos.NewMetadataRepo(app.DB)

	uu := usecases.NewUserUc(mr, sr)
	pu := usecases.NewPostUc(pr, uu)

	postCtr = controllers.NewPostCtr(pu)

	mutationFieldsList = append(mutationFieldsList, postMutationFields)
	queryFieldsList = append(queryFieldsList, postQueryFields)
}

var PostType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: controllers.ObjectNamePost,
		Fields: graphql.Fields{
			controllers.ArgsId: &graphql.Field{
				Type: graphql.Int,
			},
			controllers.ArgsCreatedAt: &graphql.Field{
				Type: graphql.DateTime,
			},
			controllers.ArgsUpdatedAt: &graphql.Field{
				Type: graphql.DateTime,
			},
			controllers.ArgsPublishedAt: &graphql.Field{
				Type: graphql.DateTime,
			},
			controllers.ArgsPostsIsPage: &graphql.Field{
				Type: graphql.Boolean,
			},
			controllers.ArgsPostsIsPublished: &graphql.Field{
				Type: graphql.Boolean,
			},
			controllers.ArgsPostsSlug: &graphql.Field{
				Type: graphql.String,
			},
			controllers.ArgsPostsTitle: &graphql.Field{
				Type: graphql.String,
			},
			controllers.ArgsPostsContent: &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
