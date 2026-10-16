#!/bin/bash
# Rewrites the downloaded Supabase Management API spec into a form
# oapi-codegen accepts. Reads api/openapi.json and writes
# api/openapi.fixed.json.
#
# Two fixes:
#
# - A parameter whose schema is anyOf a comma-separated string or an array
#   of enum values makes oapi-codegen emit a type that refers to itself and
#   is declared twice. Such a parameter keeps only its array branch; the API
#   accepts the repeated `name=a&name=b` form that the array produces.
# - Some schemas carry both `connectionString` (deprecated) and
#   `connection_string`. Both become the Go field ConnectionString, which
#   does not compile, so the deprecated one is named
#   DeprecatedConnectionString through x-go-name.
set -euo pipefail
trap 'on_error "$?" "$LINENO"' ERR

LOG_FILE="${LOG_FILE:-/tmp/$(basename "$0" .sh).log}"
exec > >(tee -a "$LOG_FILE") 2>&1

SPEC_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)/api"
readonly SPEC_DIR
readonly SPEC_IN="${SPEC_DIR}/openapi.json"
readonly SPEC_OUT="${SPEC_DIR}/openapi.fixed.json"

# Matches a parameter object with a string-or-array anyOf schema.
readonly STRING_OR_ARRAY_PARAM='type == "object" and has("in") and (.schema.anyOf | type == "array" and (map(.type) | index("array") != null and index("string") != null))'

# Matches a properties object holding both connection string spellings.
readonly BOTH_CONNECTION_STRINGS='type == "object" and (.connectionString | type == "object") and (.connection_string | type == "object")'

on_error() {
	local exit_code="$1"
	local line="$2"

	printf '{"level":"ERROR","file":"%s","line":%d,"msg":"command failed","exit_code":%d}\n' \
		"${BASH_SOURCE[1]}" "$line" "$exit_code" >&2
}

log() {
	local level="$1"
	shift

	jq -nc \
		--arg time "$(date -u '+%Y-%m-%dT%H:%M:%SZ')" \
		--arg level "$level" \
		--arg file "${BASH_SOURCE[1]##*/}" \
		--arg line "${BASH_LINENO[0]}" \
		--arg func "${FUNCNAME[1]:-main}" \
		--arg msg "$*" \
		'{time:$time,level:$level,file:$file,line:($line|tonumber),func:$func,msg:$msg}' >&2
}

main() {
	local tmp params props
	tmp="$(mktemp "${SPEC_DIR}/.openapi.fixed.XXXXXX")"
	# shellcheck disable=SC2064 # expand tmp now, the local is gone when EXIT fires
	trap "rm -f -- '${tmp}'" EXIT

	params="$(jq "[.. | select(${STRING_OR_ARRAY_PARAM})] | length" "$SPEC_IN")"
	props="$(jq "[.. | select(${BOTH_CONNECTION_STRINGS})] | length" "$SPEC_IN")"

	jq "walk(
		if ${STRING_OR_ARRAY_PARAM} then
			.schema = (.schema.anyOf | map(select(.type == \"array\")) | first)
		elif ${BOTH_CONNECTION_STRINGS} then
			.connectionString[\"x-go-name\"] = \"DeprecatedConnectionString\"
		else . end
	)" "$SPEC_IN" >"$tmp"

	chmod 0644 "$tmp"
	mv -- "$tmp" "$SPEC_OUT"
	log INFO "rewrote string-or-array params=${params} connection string pairs=${props} out=${SPEC_OUT##*/}"
}

main "$@"
