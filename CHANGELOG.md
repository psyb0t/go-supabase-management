# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.5.0, 2026-11-10

- Added `ProjectRef` (`*string`, query parameter `project_ref`) to `V1AuthorizeUserParams`. `NewV1AuthorizeUserRequest` sends it when set.
- Breaking: the `Timestamp` field of the items in `V1GetUsageApiCountResponseOutput.Result` changed from `time.Time` to `string`. Code that reads it as a time must parse the string.

## v0.4.1, 2026-11-06

- Added the partial Postgres config update endpoint (`PATCH /v1/projects/{ref}/config/database/postgres`): `ClientInterface.V1PatchPostgresConfig` and `V1PatchPostgresConfigWithBody`, with matching methods on `Client`, `ClientWithResponses` (`V1PatchPostgresConfigWithResponse`, `V1PatchPostgresConfigWithBodyWithResponse`), and the helpers `NewV1PatchPostgresConfigRequest`, `NewV1PatchPostgresConfigRequestWithBody`, `ParseV1PatchPostgresConfigResponse`.
- Added the response type `V1PatchPostgresConfigResponse` and the request body alias `V1PatchPostgresConfigJSONRequestBody` for `UpdatePostgresConfigBody`.
- Added `ProjectRefs` (`*[]string`, `null` means all projects) and `Scope` (`*string`) to `OAuthTokenResponseOutput`.
- Custom implementations of `ClientInterface` and `ClientWithResponsesInterface` must add the new `V1PatchPostgresConfig` methods.

## v0.4.0, 2026-11-05

- Added the network restrictions review endpoint: `ClientInterface.V1ReviewNetworkRestrictions` and `V1ReviewNetworkRestrictionsWithBody`, with matching methods on `Client`, `ClientWithResponses` (`V1ReviewNetworkRestrictionsWithResponse`, `V1ReviewNetworkRestrictionsWithBodyWithResponse`), and the helpers `NewV1ReviewNetworkRestrictionsRequest`, `NewV1ReviewNetworkRestrictionsRequestWithBody`, `ParseV1ReviewNetworkRestrictionsResponse`.
- Added the request type `NetworkRestrictionsReviewRequest` (alias `V1ReviewNetworkRestrictionsJSONRequestBody`) and the enum `NetworkRestrictionsReviewRequestAction` with `NetworkRestrictionsReviewRequestActionApply` and `NetworkRestrictionsReviewRequestActionReject`.
- Added `PendingConfig` and `PendingRequestedAt` to `NetworkRestrictionsResponseOutput`.
- Added `ColumnName`, `SchemaName`, `TableName` and `TypeName` to `ProjectUpgradeEligibilityResponseOutputValidationErrors8`, and the constant `ProjectUpgradeEligibilityResponseOutputValidationErrors8TypeTableUsingRegType`.
- Added the union variant `ProjectUpgradeEligibilityResponseOutputValidationErrors10` with `AsProjectUpgradeEligibilityResponseOutputValidationErrors10`, `FromProjectUpgradeEligibilityResponseOutputValidationErrors10` and `MergeProjectUpgradeEligibilityResponseOutputValidationErrors10`.
- Breaking: removed `ProjectUpgradeEligibilityResponseOutputValidationErrors8TypeX86Architecture` and `ProjectUpgradeEligibilityResponseOutputValidationErrors9TypeProjectHibernating`. The values moved to `ProjectUpgradeEligibilityResponseOutputValidationErrors9TypeX86Architecture` and `ProjectUpgradeEligibilityResponseOutputValidationErrors10TypeProjectHibernating`.
- Added `ClientInterface.V1ReviewNetworkRestrictions` methods, so custom implementations of `ClientInterface` and `ClientWithResponsesInterface` must add them.

## v0.3.7, 2026-11-04

- Added the project status value `PROMOTING` to five enums: `BranchDetailResponseOutputStatusPROMOTING`, `BranchResponseOutputPreviewProjectStatusPROMOTING`, `OrganizationProjectsResponseOutputProjectsStatusPROMOTING`, `V1ProjectResponseOutputStatusPROMOTING` and `V1ProjectWithDatabaseResponseOutputStatusPROMOTING`. Their `Valid` methods accept it.
- Added a fifth warning variant to the project upgrade eligibility response: `ProjectUpgradeEligibilityResponseOutputWarnings4`, with a `Type` field of the new type `ProjectUpgradeEligibilityResponseOutputWarnings4Type` (value `ProjectUpgradeEligibilityResponseOutputWarnings4TypeCollationReindexRequired`, `collation_reindex_required`).
- Added the union methods `AsProjectUpgradeEligibilityResponseOutputWarnings4`, `FromProjectUpgradeEligibilityResponseOutputWarnings4` and `MergeProjectUpgradeEligibilityResponseOutputWarnings4` on `ProjectUpgradeEligibilityResponseOutput_Warnings_Item`.
- No existing exported names changed.
- Generated from the Supabase Management API spec, still version 1.0.0.

## v0.3.6, 2026-10-31

- Added the PostgREST pool settings `DbPoolMaxIdletime` and `DbPoolMaxLifetime` (both `*int`) to `V1UpdatePostgrestConfigBody`, `V1PostgrestConfigResponseOutput` and `PostgrestConfigWithJWTSecretResponseOutput`.
- On the response types the fields are `db_pool_max_idletime` and `db_pool_max_lifetime`. A `null` value means the server uses 30 for idle time and 1800 for lifetime.
- No existing exported names changed.
- Generated from the Supabase Management API spec, still version 1.0.0.

## v0.3.5, 2026-10-28

- Regenerated the generated code (1 file changed, 47 insertions(+), 44 deletions(-)).
- `V1ProjectAdvisorsResponseOutputLintsNameRealtimeConcurrentUsersNearLimit`: added.

## v0.3.4, 2026-10-28

- Updated Semgrep in the development image to 1.182.1 and refreshed its locked dependencies. certifi 2026.10.23 drops the emSign root CAs whose intermediate keys were compromised (GHSA-6q48-6c7w-9gw2).
- Dependabot no longer opens pull requests for `mcp` or the OpenTelemetry packages in `requirements-semgrep.txt`. Semgrep pins their versions itself, so bumping one of them on its own cannot install.
- Updated kin-openapi, which the code generator uses, from 0.142.0 to 0.149.0. The newer version fixes several panics on malformed specs reported in GitHub security advisories.
- The security scan merges its two gitleaks runs into one SARIF run. GitHub code scanning rejects a report with two runs from the same tool, so no scan results had been uploaded.

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
