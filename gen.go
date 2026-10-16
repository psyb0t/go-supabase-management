// Package supabase is a Go client for the Supabase Management API, generated
// with oapi-codegen from the OpenAPI document Supabase publishes.
package supabase

//go:generate bash scripts/fetch-spec.sh
//go:generate bash scripts/fix-spec.sh
//go:generate go tool oapi-codegen --config=config.yaml api/openapi.fixed.json
//go:generate go tool oapi-codegen --config=config.client.yaml api/openapi.fixed.json
