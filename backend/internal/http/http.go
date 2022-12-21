package http

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	address := ":8080"
	srv := &http.Server{
		Addr:    address,
		Handler: http.DefaultServeMux,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Infof("running HTTP server in localhost%s", address)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error(err)
		}
	}()

	<-done
	log.Info("HTTP server is shutting down..")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %+v", err)
	}
	log.Info("Server shutdown properly.")
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
