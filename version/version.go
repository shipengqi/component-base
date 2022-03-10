// Package version supplies version information collected at build time to apps.
package version

import (
	"fmt"
	"runtime"

	"github.com/gosuri/uitable"

	"github.com/shipengqi/component-base/json"
)

var (
	// Version is semantic version.
	Version = "v0.0.0-master+$Format:%h$"
	// BuildDate in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ').
	BuildDate = "1970-01-01T00:00:00Z"
	// GitCommit sha1 from git, output of $(git rev-parse HEAD).
	GitCommit = "$Format:%H$"
)

// Info contains versioning information.
type Info struct {
	Version   string `json:"version"`
	GoVersion string `json:"goVersion"`
	GitCommit string `json:"gitCommit"`
	BuildDate string `json:"buildDate"`
	Platform  string `json:"platform"`
}

// String returns info as a human-friendly version string.
func (info Info) String() string {
	if s, err := info.Text(); err == nil {
		return string(s)
	}

	return info.Version
}

// ToJSON returns the JSON string of version information.
func (info Info) ToJSON() string {
	s, _ := json.Marshal(info)

	return string(s)
}

// Text encodes the version information into UTF-8-encoded text and
// returns the result.
func (info Info) Text() ([]byte, error) {
	table := uitable.New()
	table.RightAlign(0)
	table.MaxColWidth = 80
	table.Separator = " "
	table.AddRow("Version:", info.Version)
	table.AddRow("Go Version:", info.GoVersion)
	table.AddRow("Git Commit:", info.GitCommit)
	table.AddRow("Built:", info.BuildDate)
	table.AddRow("OS/Arch:", info.Platform)

	return table.Bytes(), nil
}

// Get returns the overall codebase version. It's for detecting
// what code a binary was built from.
func Get() Info {
	// These variables typically come from -ldflags settings and in
	// their absence fallback to the settings in pkg/version/base.go
	return Info{
		Version:   Version,
		GitCommit: GitCommit,
		BuildDate: BuildDate,
		GoVersion: runtime.Version(),
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
