package gql

import (
	"github.com/ganiszulfa/concise/backend/config/app"
	"github.com/ganiszulfa/concise/backend/internal/controllers"
	"github.com/ganiszulfa/concise/backend/internal/repos"
	"github.com/ganiszulfa/concise/backend/internal/usecases"
	"github.com/graphql-go/graphql"
)

var metadataCtr controllers.MetadataCtrInterface

func (g *G) InitializeMetadata() {
	r := repos.NewMetadataRepo(app.DB)
	u := usecases.NewMetadataUc(r)
	metadataCtr = controllers.NewMetadataCtr(u)

	queryFieldsList = append(queryFieldsList, mdQueryFields)
}

var MetadataType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: controllers.ObjectNameMetadata,
		Fields: graphql.Fields{
			controllers.ArgsMetadataKey: &graphql.Field{
				Type: graphql.String,
			},
			controllers.ArgsMetadataValue: &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
