package version

// Base version information.
var (
	// Version is the current version of application.
	Version = "v0.0.0-master+$Format:%h$"

	// GitCommit sha1 from git, output of $(git rev-parse HEAD).
	GitCommit = "$Format:%H$"

	// GitTreeState state of git tree, either "clean" or "dirty".
	GitTreeState string

	// BuildTime in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ').
	BuildTime = "1970-01-01T00:00:00Z"
)
