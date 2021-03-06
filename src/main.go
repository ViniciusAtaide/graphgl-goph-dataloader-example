package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"runtime"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/viniciusataide/graphql-go-example/src/resolvers"
)

func main() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Join(filepath.Dir(b), "..")

	bstr, err := ioutil.ReadFile(basepath + "/schema.graphql")

	if err != nil {
		panic(err)
	}

	db, err := sqlx.Connect("postgres", "user=postgres dbname=postgres sslmode=disable password=pass")

	if err != nil {
		panic(err)
	}

	schemaString := string(bstr)

	schema := graphql.MustParseSchema(schemaString, &resolvers.RootResolver{Db: *db})

	type JSON = map[string]interface{}

	http.Handle("/query", &relay.Handler{Schema: schema})
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	}))

	log.Println("Listening on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))

}

var page = []byte(`
<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.css" />
		<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.1.0/fetch.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.js"></script>
	</head>
	<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
		<div id="graphiql" style="height: 100vh;">Loading...</div>
		<script>
			function graphQLFetcher(graphQLParams) {
				return fetch("/query", {
					method: "post",
					body: JSON.stringify(graphQLParams),
					credentials: "include",
				}).then(function (response) {
					return response.text();
				}).then(function (responseBody) {
					try {
						return JSON.parse(responseBody);
					} catch (error) {
						return responseBody;
					}
				});
			}
			ReactDOM.render(
				React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
				document.getElementById("graphiql")
			);
		</script>
	</body>
</html>
`)
