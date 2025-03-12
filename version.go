package main

// Base version information.
// We use semantic version (see https://semver.org/ for more information).
var (
	// Version is updated by the Makefile to reflect the new version when releasing; a git-annotated tag sets this version.
	Version = "" // git tag, output of $(git describe --tags --always --dirty)

	// DefaultVersion is used as a fallback when version information from git is not
	// provided via go ldflags. It approximates the application version for adhoc builds
	// (e.g., `go build`) that cannot retrieve the version information from git.
	DefaultVersion = "0.0.0-git"
)
