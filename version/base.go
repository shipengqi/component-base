package version

// Base version information.
var (
	// Version is the current version of crtctl, set by the go linker's -X flag at build time.
	Version = "v0.0.0-master+$Format:%h$"

	// GitCommit is the actual commit that is being built, set by the go linker's -X flag at build time.
	GitCommit = "$Format:%H$"

	// GitTreeState indicates if the git tree is clean or dirty, set by the go linker's -X flag at build
	// time.
	GitTreeState string

	// BuildTime in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ').
	BuildTime = "1970-01-01T00:00:00Z"
)
