# Copilot Instructions for `goark/gocli`

## Project purpose

`gocli` is a collection of small, focused packages for building command-line tools in Go.
Each sub-package addresses one concern: I/O handling, exit codes, signal handling, file globbing, and XDG directories.

## Design principles

- Keep each sub-package small and focused on a single responsibility.
- Prefer explicit, simple APIs over feature-rich abstractions.
- Preserve backward compatibility when possible; avoid breaking public symbols.
- Use functional options (`OptFunc`) where configuration is needed.

## Coding style

- Write idiomatic Go. Keep implementation straightforward.
- Always add a space after `//` in comments: `// Foo`, not `//Foo`.
- Add a blank line before `Deprecated:` in doc comments.
- Avoid unnecessary external dependencies.
- Keep comments concise and in English.

## Error handling

- Return errors explicitly; do not panic in library code.
- Keep error sentinel behavior compatible with `errors.Is`.

## Testing and validation

- Add or update tests for behavior changes.
- Run local validation with:

```text
task test
```

- `task test` runs: `go mod verify`, `go test -shuffle on ./...`, and golangci-lint.

## Release process

- Create release tags from `master`.
- Use semantic versioning tags in `vMAJOR.MINOR.PATCH` format.
- Ensure the local repository is clean and synced before tagging.

Release steps:

1. Ensure `master` is up to date.
2. Create annotated tag:
   - `git tag -a vX.Y.Z -m "Release vX.Y.Z"`
3. Push tag:
   - `git push origin vX.Y.Z`
4. Create GitHub release with auto-generated notes:
   - `gh release create vX.Y.Z --generate-notes`
5. Verify the release on GitHub.
