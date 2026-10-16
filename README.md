# go-supabase-management

[![Go Reference](https://pkg.go.dev/badge/github.com/psyb0t/go-supabase-management.svg)](https://pkg.go.dev/github.com/psyb0t/go-supabase-management)
[![CI](https://github.com/psyb0t/go-supabase-management/actions/workflows/pipeline.yml/badge.svg?branch=main)](https://github.com/psyb0t/go-supabase-management/actions/workflows/pipeline.yml)
[![version](https://raw.githubusercontent.com/psyb0t/go-supabase-management/badges/version.svg)](https://github.com/psyb0t/go-supabase-management/tags)
[![license](https://raw.githubusercontent.com/psyb0t/go-supabase-management/badges/license.svg)](LICENSE)
[![imported by](https://raw.githubusercontent.com/psyb0t/go-supabase-management/badges/importers.svg)](https://github.com/psyb0t/go-supabase-management/blob/badges/importers.md)

Unofficial Go client for the [Supabase Management API](https://supabase.com/docs/reference/api/introduction), generated with [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) from the OpenAPI document Supabase publishes. I am not affiliated with Supabase.

This is the API for managing projects, organizations, branches, database settings, edge functions, secrets and the like. It is not the per-project data API (PostgREST, Auth, Storage).

The package is `supabase`, in the repository root. Everything in `types.gen.go` and `client.gen.go` is generated. Don't edit those files by hand, regenerate them.

## Install

```sh
go get github.com/psyb0t/go-supabase-management
```

## Usage

Every endpoint in the spec has a typed method on `ClientWithResponses`. Responses decode into `JSON200` and the other status fields.

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	supabase "github.com/psyb0t/go-supabase-management"
)

func main() {
	auth := supabase.WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+os.Getenv("SUPABASE_ACCESS_TOKEN"))

		return nil
	})

	client, err := supabase.NewClientWithResponses("https://api.supabase.com", auth)
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
}
```

The same code is in `example_test.go`, run against a local test server, so `make test` checks that it works.

Method names come from the spec's operation IDs, so they all start with `V1`: `V1ListAllOrganizations`, `V1GetProject`, `V1CreateABranch` and so on. Most project endpoints take the 20 character project ref as their first argument after the context.

## Auth

The Management API takes a Bearer token in the `Authorization` header. It is either a personal access token, created under Access Tokens in the Supabase dashboard, or an OAuth access token issued to a Supabase OAuth app. The request editor is the same for both, as in the example above. An OAuth token only reaches the endpoints its approved scopes allow.

## Regenerating

The spec is not committed. `make generate` downloads the current OpenAPI document into the gitignored `api/` directory and regenerates the client:

```sh
make generate
```

`gen.go` declares the steps, so `make generate` is just `go generate ./...`:

- `scripts/fetch-spec.sh` downloads `https://api.supabase.com/api/v1-json` to `api/openapi.json`.
- `scripts/fix-spec.sh` writes `api/openapi.fixed.json` with two fixes. The `services` query parameter of the project health endpoint is `anyOf` a comma-separated string or an array of enum values, which makes oapi-codegen emit a type that refers to itself; the script keeps the array form. The pooler config response has both `connectionString` (deprecated) and `connection_string`, which would both become the Go field `ConnectionString`; the script names the deprecated one `DeprecatedConnectionString`.
- oapi-codegen runs twice: `config.yaml` writes the models to `types.gen.go` and `config.client.yaml` writes the client to `client.gen.go`. Both map `number` to `float64`, because oapi-codegen would otherwise use `float32` for numbers with no format.

## Agent integrations

The [skill](.agents/skills/go-supabase-management) teaches a coding agent how to use go-supabase-management when it writes Go. It works in any agent that reads `.agents/skills/`, and installs natively in the clients below.

### Claude Code

```bash
claude plugin marketplace add psyb0t/agents
claude plugin install go-supabase-management@psyb0t
```

### Codex

```bash
codex plugin marketplace add psyb0t/agents
codex plugin add go-supabase-management@psyb0t
```

Codex also picks the skill up with no install in any checkout of this repo.

### OpenClaw

The skill is published to ClawHub on every release:

```bash
openclaw skills install @psyb0t/go-supabase-management
```

## Development

Every command runs in the `Dockerfile.dev` image. The host needs only Docker and Make.

```sh
make help      # list targets
make dep       # go mod tidy + vendor
make lint      # shellcheck, shfmt, go fix, golangci-lint
make test      # go test -race
make sec       # gitleaks, govulncheck, semgrep
make generate  # download the spec and regenerate
```

## License

[MIT](LICENSE). The generated code comes from Supabase's published API specification. The specification itself is not redistributed here.
