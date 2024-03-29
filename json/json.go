//go:build !jsoniter && !go_json && !(sonic && avx && (linux || windows || darwin) && amd64)

package json

import "encoding/json"

// RawMessage is exported by component-base/json package.
type RawMessage = json.RawMessage

var (
	// Marshal is exported by component-base/pkg/json package.
	Marshal = json.Marshal
	// Unmarshal is exported by component-base/pkg/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by component-base/pkg/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by component-base/pkg/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by component-base/pkg/json package.
	NewEncoder = json.NewEncoder
)
