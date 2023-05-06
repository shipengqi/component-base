package version

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/gosuri/uitable"
)

type Info struct {
	Version      string `json:"version"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildTime    string `json:"buildTime"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

// String returns info as a human-friendly version string.
func (i Info) String() string {
	if s, err := i.Text(); err == nil {
		return string(s)
	}

	return i.Version
}

// ToJSON returns the JSON string of version information.
func (i Info) ToJSON() string {
	s, _ := json.Marshal(i)

	return string(s)
}

// Text encodes the version information into UTF-8-encoded text and
// returns the result.
func (i Info) Text() ([]byte, error) {
	table := uitable.New()
	table.RightAlign(0)
	table.MaxColWidth = 80
	table.Separator = " "
	table.AddRow("gitVersion:", i.Version)
	table.AddRow("gitCommit:", i.GitCommit)
	table.AddRow("gitTreeState:", i.GitTreeState)
	table.AddRow("buildTime:", i.BuildTime)
	table.AddRow("goVersion:", i.GoVersion)
	table.AddRow("compiler:", i.Compiler)
	table.AddRow("platform:", i.Platform)

	return table.Bytes(), nil
}

// Get returns the overall codebase version. It's for detecting
// what code a binary was built from.
func Get() Info {
	// These variables typically come from -ldflags settings and in
	// their absence fallback to the settings in pkg/version/base.go
	return Info{
		Version:      Version,
		GitCommit:    GitCommit,
		GitTreeState: GitTreeState,
		BuildTime:    BuildTime,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
