package mongodb

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/anup-gupta1/go-graphql-mongodb/schema"

	"github.com/graphql-go/graphql"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

func main() {
	fmt.Println("hello")
	schema, err := graphql.NewSchema(schema.SchemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
