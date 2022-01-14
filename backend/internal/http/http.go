package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ganiszulfa/concise/backend/config"
	"github.com/ganiszulfa/concise/backend/internal/gql"
	"github.com/ganiszulfa/concise/backend/internal/middleware"
	"github.com/ganiszulfa/concise/backend/pkg/trace"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var graphQL *handler.Handler

func Serve() {
	trace.Func()

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    gql.QueryType,
			Mutation: gql.MutationType,
		},
	)

	if err != nil {
		panic(err)
	}

	graphQL = handler.New(&handler.Config{
		Schema:           &schema,
		Pretty:           true,
		GraphiQL:         true,
		ResultCallbackFn: resultCallback,
	})

	http.HandleFunc("/health", healthHandler)

	http.HandleFunc("/gql", middleware.Authorize(gqlHandler))

	if config.Config.HttpServer.Port == "" {
		log.Warn("http port is not defined, not running the http server...")
		return
	}

	address := fmt.Sprintf(":%s", config.Config.HttpServer.Port)
	log.Infof("running HTTP server in %s", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Error(err)
	}
}

func resultCallback(
	ctx context.Context,
	params *graphql.Params,
	result *graphql.Result,
	responseBody []byte,
) {
	trace.Func()

	if result.Errors != nil {
		for _, e := range result.Errors {
			if e.Message != gorm.ErrRecordNotFound.Error() {
				log.Error(e.Message, " > ", e.Path)
			} else {
				log.Debug(e.Message, " > ", e.Path)
			}
		}
	}
}

func healthHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("ok"))
	if err != nil {
		log.Error(err)
	}
}

func gqlHandler(w http.ResponseWriter, req *http.Request) {
	trace.Func()

	graphQL.ServeHTTP(w, req)
}
