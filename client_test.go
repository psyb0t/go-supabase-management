package supabase_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	supabase "github.com/psyb0t/go-supabase-management"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testToken      = "test-token"
	testProjectRef = "abcdefghijklmnopqrst"
)

func newTestClient(
	t *testing.T,
	handler http.HandlerFunc,
) *supabase.ClientWithResponses {
	t.Helper()

	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	auth := supabase.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set("Authorization", "Bearer "+testToken)

			return nil
		},
	)

	client, err := supabase.NewClientWithResponses(srv.URL, auth)
	require.NoError(t, err)

	return client
}

func assertBearer(t *testing.T, r *http.Request) {
	t.Helper()

	assert.Equal(t, "Bearer "+testToken, r.Header.Get("Authorization"))
}

func TestListOrganizations(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/organizations", r.URL.Path)
		assertBearer(t, r)

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[{"id":"org-1","name":"Org One"}]`))
	})

	res, err := client.V1ListAllOrganizationsWithResponse(
		context.Background(),
	)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.NotNil(t, res.JSON200)
	require.Len(t, *res.JSON200, 1)

	org := (*res.JSON200)[0]
	assert.Equal(t, "org-1", org.Id)
	assert.Equal(t, "Org One", org.Name)
}

func TestGetProject(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/projects/"+testProjectRef, r.URL.Path)
		assertBearer(t, r)

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(
			`{"id":"` + testProjectRef + `","ref":"` + testProjectRef + `",` +
				`"name":"demo","region":"us-east-1",` +
				`"status":"ACTIVE_HEALTHY","organization_id":"org-1",` +
				`"organization_slug":"org-1","created_at":"2026-01-02",` +
				`"database":{"host":"db.example.test","version":"17",` +
				`"postgres_engine":"17","release_channel":"ga"}}`,
		))
	})

	res, err := client.V1GetProjectWithResponse(
		context.Background(),
		testProjectRef,
	)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.NotNil(t, res.JSON200)

	assert.Equal(t, testProjectRef, res.JSON200.Ref)
	assert.Equal(t, "us-east-1", res.JSON200.Region)
	assert.Equal(t, "ACTIVE_HEALTHY", string(res.JSON200.Status))
	assert.Equal(t, "db.example.test", res.JSON200.Database.Host)
}
