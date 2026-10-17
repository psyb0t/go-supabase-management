# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.1.1, 2026-10-17

- Documentation only. The README and `AGENTS.md` describe the daily regeneration workflow: what it checks, how it fixes failing checks, how it picks the next version and which paths a fix may not touch.

## v0.1.0, 2026-10-16

- Added `supabase`, a REST client generated with oapi-codegen from Supabase's OpenAPI document (Supabase API (v1), spec version 1.0.0, OpenAPI 3.0.0, from api.supabase.com/api/v1-json). It covers all 115 paths of the Management API: organizations, projects, branches, database and pooler config, auth config, edge functions, secrets, custom domains, network restrictions, billing add-ons, analytics and OAuth.
- Number fields are `float64` instead of oapi-codegen's default `float32`.
- `make generate` downloads the current spec, keeps the array form of the health endpoint's string-or-array `services` parameter, renames the deprecated `connectionString` field to `DeprecatedConnectionString`, and regenerates the client into `types.gen.go` and `client.gen.go`. The spec is not committed.
- Added the `go-supabase-management` agent skill. Install it with `claude plugin install go-supabase-management@psyb0t` (Claude Code), `codex plugin add go-supabase-management@psyb0t` (Codex) or `openclaw skills install @psyb0t/go-supabase-management` (ClawHub).
