package supabase_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	supabase "github.com/psyb0t/go-supabase-management"
)

// This is the README usage example. A local server stands in for
// https://api.supabase.com so the example runs without a token.
func Example() {
	srv := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`[{"ref":"abcdefghijklmnopqrst"}]`))
		},
	))
	defer srv.Close()

	auth := supabase.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set(
				"Authorization",
				"Bearer "+os.Getenv("SUPABASE_ACCESS_TOKEN"),
			)

			return nil
		},
	)

	client, err := supabase.NewClientWithResponses(srv.URL, auth)
	if err != nil {
		panic(err)
	}

	res, err := client.V1ListAllProjectsWithResponse(context.Background())
	if err != nil {
		panic(err)
	}

	for _, project := range *res.JSON200 {
		fmt.Println(project.Ref)
	}
	// Output: abcdefghijklmnopqrst
}
