# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.3.3, 2026-10-27

- The development image is based on Debian (`golang:1.26.10-trixie`) instead of Alpine. Alpine keeps only the latest build of each package, so the exact pins broke the image build whenever a package was updated. System packages now come from the snapshot.debian.org date the base image was built from, which keeps them fixed for a given base image digest. jq 1.8.2, shellcheck 0.11.0 and shfmt 3.13.1 keep their versions, and Semgrep runs in its own virtual environment.

## v0.3.2, 2026-10-26

- The Supabase Management API now documents a 404 response for the project API key endpoints. The generated parsers `ParseV1GetProjectApiKeyResponse`, `ParseV1UpdateProjectApiKeyResponse` and `ParseV1DeleteProjectApiKeyResponse` handle status 404 as a response with no body.
- No exported names changed.
- Generated from the Supabase Management API spec, still version 1.0.0.

## v0.3.1, 2026-10-23

- Added two advisor lint names to `V1ProjectAdvisorsResponseOutputLintsName`: `V1ProjectAdvisorsResponseOutputLintsNameRealtimeAuthorizationTimeouts` (`realtime_authorization_timeouts`) and `V1ProjectAdvisorsResponseOutputLintsNameRealtimePostgresChangesSpike` (`realtime_postgres_changes_spike`). The `Valid` method accepts both.
- Raised the Go requirement from 1.26.9 to 1.26.10 in `go.mod`, the development image and the setup notes of the agent skill. Go 1.26.10 fixes GO-2026-6809 in `crypto/x509`, which `govulncheck` flagged.
- Generated from the Supabase Management API spec, still version 1.0.0.

## v0.3.0, 2026-10-20

- The Supabase Management API now describes the update-function request body for two content types. The generated `V1UpdateFunctionBody` type is renamed to `V1UpdateFunctionBodyJson`, with the same fields `Body`, `Name` and `VerifyJwt`.
- Added `V1UpdateFunctionBodyEszip`, an alias of `openapi_types.File`, for the raw Brotli-compressed ESZIP upload. Its doc comment lists the `EZBR` magic bytes and the `ezbr_sha256` query parameter.
- `V1UpdateAFunctionJSONRequestBody` now aliases `V1UpdateFunctionBodyJson`. Code that names `V1UpdateFunctionBody` must switch to the new name.
- Generated from the Supabase Management API spec, still version 1.0.0.

## v0.2.0, 2026-10-19

- The Supabase Management API renamed the `RegionsInfo_Output` schema to `RegionsInfoV1_Output`. The generated `RegionsInfoOutput` type is renamed to `RegionsInfoV1Output`, along with its twelve nested enum types: `RegionsInfoOutputAllSmartGroupCode`, `RegionsInfoOutputAllSmartGroupType`, `RegionsInfoOutputAllSpecificCode`, `RegionsInfoOutputAllSpecificProvider`, `RegionsInfoOutputAllSpecificStatus`, `RegionsInfoOutputAllSpecificType`, `RegionsInfoOutputRecommendationsSmartGroupCode`, `RegionsInfoOutputRecommendationsSmartGroupType`, `RegionsInfoOutputRecommendationsSpecificCode`, `RegionsInfoOutputRecommendationsSpecificProvider`, `RegionsInfoOutputRecommendationsSpecificStatus` and `RegionsInfoOutputRecommendationsSpecificType` become the matching `RegionsInfoV1Output...` names.
- `V1GetAvailableRegionsResponse.JSON200` and its `GetJSON200` method now return `*RegionsInfoV1Output` instead of `*RegionsInfoOutput`.
- The fields and enum values are the same. Code that names the old types must switch to the new names.
- Generated from the Supabase Management API spec, still version 1.0.0.

## v0.1.2, 2026-10-17

- Documentation only. The agent skill explains how to look up every endpoint, parameter and model with `go doc` or on pkg.go.dev, and declares its license.
- The skill no longer lists the API credentials under OpenClaw's `requires.env`, which kept OpenClaw from loading it until they were set. It requires only `go`; the credentials are still documented as what your program reads.
- The README points to the full API reference, and the examples in the README and the skill were checked against the current code.

## v0.1.1, 2026-10-17

- Documentation only. The README and `AGENTS.md` describe the daily regeneration workflow: what it checks, how it fixes failing checks, how it picks the next version and which paths a fix may not touch.

## v0.1.0, 2026-10-16

- Added `supabase`, a REST client generated with oapi-codegen from Supabase's OpenAPI document (Supabase API (v1), spec version 1.0.0, OpenAPI 3.0.0, from api.supabase.com/api/v1-json). It covers all 115 paths of the Management API: organizations, projects, branches, database and pooler config, auth config, edge functions, secrets, custom domains, network restrictions, billing add-ons, analytics and OAuth.
- Number fields are `float64` instead of oapi-codegen's default `float32`.
- `make generate` downloads the current spec, keeps the array form of the health endpoint's string-or-array `services` parameter, renames the deprecated `connectionString` field to `DeprecatedConnectionString`, and regenerates the client into `types.gen.go` and `client.gen.go`. The spec is not committed.
- Added the `go-supabase-management` agent skill. Install it with `claude plugin install go-supabase-management@psyb0t` (Claude Code), `codex plugin add go-supabase-management@psyb0t` (Codex) or `openclaw skills install @psyb0t/go-supabase-management` (ClawHub).
