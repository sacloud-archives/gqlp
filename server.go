// Copyright 2021 The gqlp Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gqlp

import (
	"log"
	"net/http"
	"strings"

	"github.com/sacloud/gqlp/trace"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/sacloud/gqlp/graph"
	"github.com/sacloud/gqlp/graph/generated"
)

func Serve(port string) error {
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	http.Handle("/", PlaygroundHandler("/query"))
	http.Handle("/query", GraphQLQueryWithTraceHandler())

	log.Printf("connect to http://localhost%s/ for GraphQL playground", port)
	return http.ListenAndServe(port, nil)
}

func PlaygroundHandler(gqlQueryEndpoint string) http.Handler {
	return playground.Handler("GraphQL playground", gqlQueryEndpoint)
}

func GraphQLQueryHandler() http.Handler {
	return handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
}

func GraphQLQueryWithTraceHandler() http.Handler {
	return traceHandler(GraphQLQueryHandler())
}

func traceHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.Tracer.Start(r.Context(), "sacloud/gqlp")
		defer span.End()

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
