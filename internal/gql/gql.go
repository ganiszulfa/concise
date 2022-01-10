package gql

import (
	"reflect"

	"github.com/ganiszulfa/concise/pkg/trace"
	"github.com/graphql-go/graphql"
)

var QueryType *graphql.Object
var MutationType *graphql.Object
var mutationFieldsList []graphql.Fields
var queryFieldsList []graphql.Fields

type G struct{} // used for adding fields

func Initialize() {
	trace.Func()

	addAllFields()

	allMutationFields := make(graphql.Fields)
	allQueryFields := make(graphql.Fields)
	mergeAllGqlFields(&allMutationFields, mutationFieldsList)
	mergeAllGqlFields(&allQueryFields, queryFieldsList)

	MutationType = graphql.NewObject(
		graphql.ObjectConfig{
			Name:   "Mutation",
			Fields: allMutationFields,
		})

	QueryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name:   "Query",
			Fields: allQueryFields,
		})
}

func mergeAllGqlFields(fields *graphql.Fields, list []graphql.Fields) {
	trace.Func()

	for _, i := range list {
		for k, v := range i {
			_, alreadyExist := (*fields)[k]
			if alreadyExist {
				panic("duplicate field: " + k)
			}
			(*fields)[k] = v
		}
	}
}

func addAllFields() {
	trace.Func()

	g := G{}
	methodFinder := reflect.TypeOf(&g)
	for i := 0; i < methodFinder.NumMethod(); i++ {
		method := methodFinder.Method(i)
		method.Func.Call([]reflect.Value{reflect.ValueOf(&g)})
	}
}
