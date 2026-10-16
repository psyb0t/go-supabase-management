# Testing

Refines `~/.claude/rule-details/go/writing-tests.md`.

- **TST1** Tests are black-box, in package `supabase_test`, and use only the exported API.
- **TST2** Every test runs the client against an `httptest` server and asserts the method, the path, the `Authorization` header and the decoded response. No test talks to the real API.
- **TST3** `example_test.go` is the README usage example and keeps an `// Output:` line against a local server. Change the README and the example together.
- **TST4** Test data uses obvious placeholders: fake tokens, the spec's example project ref, made-up names.
