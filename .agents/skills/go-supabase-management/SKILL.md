---
name: go-supabase-management
description: Unofficial Go client for the Supabase Management API (`github.com/psyb0t/go-supabase-management`, package `supabase`), generated with oapi-codegen from Supabase's OpenAPI document. Typed methods on `supabase.ClientWithResponses` cover all 115 paths of `https://api.supabase.com`, among them organizations, projects, branches, database and pooler config, auth config, edge functions, secrets, API keys, custom domains, network restrictions, billing add-ons, analytics and OAuth. Requests authenticate with a Bearer personal access token or OAuth access token set in a request editor. Use when Go code needs to list or manage Supabase projects, organizations, branches or project settings, or already imports go-supabase-management. Not for the per-project data API (PostgREST, Auth, Storage).
homepage: https://github.com/psyb0t/go-supabase-management
user-invocable: true
metadata:
  openclaw:
    emoji: "🗄️"
    requires:
      bins:
        - go
      env:
        - SUPABASE_ACCESS_TOKEN
permissions:
  filesystem: "Guides edits to the Go source of the current project. The library reads and writes no files."
  shell: "The Go toolchain (`go get`, `go build`, `go test`) to add and compile the dependency."
  network: "`go get` fetches the module from your Go module proxy. At runtime the client sends HTTPS requests only to the base URL you pass, https://api.supabase.com."
---

# go-supabase-management

Typed Go client for the Supabase Management API. Every endpoint in Supabase's OpenAPI document has a method on `supabase.ClientWithResponses`, and responses decode into `JSON200`, `JSON201` and the other status fields.

For install, auth options, method naming and the operation groups, see [references/setup.md](references/setup.md).

## Security & safety

- Read the token from the environment (`SUPABASE_ACCESS_TOKEN`). Never hardcode it, commit it, or log it, and never print a request's headers. The same goes for anything the API returns as a secret: project API keys, secret values, database passwords.
- Requests go only to `https://api.supabase.com`.
- A personal access token acts with the full rights of its user across every organization. Delete, pause, restore, reset, merge, secret and config write endpoints change live projects: call them only when the user asked for that change, and only on the project ref or branch the user named.

## When to use

- Go code needs to list organizations or projects, read project status, or manage branches.
- Go code automates project settings: auth, database, pooler, network restrictions, custom domains, edge functions, secrets.
- The project already imports `github.com/psyb0t/go-supabase-management`.

## When NOT to use

- You want to read or write rows, call Auth for end users, or use Storage inside one project. That is the project's data API, served from `https://<ref>.supabase.co`, not this client.
- The code is not Go, or a single `supabase` CLI command does the job.

## Usage

```go
import (
	"context"
	"fmt"
	"net/http"
	"os"

	supabase "github.com/psyb0t/go-supabase-management"
)

auth := supabase.WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SUPABASE_ACCESS_TOKEN"))

	return nil
})

client, err := supabase.NewClientWithResponses("https://api.supabase.com", auth)
if err != nil {
	return err
}

projects, err := client.V1ListAllProjectsWithResponse(ctx)
if err != nil {
	return err // transport error
}

if projects.JSON200 == nil {
	return fmt.Errorf("supabase: status %d: %s", projects.StatusCode(), projects.Body)
}

for _, project := range *projects.JSON200 {
	fmt.Println(project.Ref, project.Name, project.Region, project.Status)
}
```

One project and its branches, by the 20 character project ref:

```go
project, err := client.V1GetProjectWithResponse(ctx, ref)
if err != nil {
	return err
}

if project.JSON200 == nil {
	return fmt.Errorf("supabase: status %d: %s", project.StatusCode(), project.Body)
}

fmt.Println(project.JSON200.Database.Host, project.JSON200.Database.Version)

branches, err := client.V1ListAllBranchesWithResponse(ctx, ref)
if err != nil {
	return err
}

if branches.JSON200 != nil {
	for _, branch := range *branches.JSON200 {
		fmt.Println(branch.Name, branch.Status)
	}
}
```

A non-2xx response is not a Go error. Check the status field you expect (`JSON200`, `JSON201`) or `StatusCode()`, and read `Body` for the raw error.
