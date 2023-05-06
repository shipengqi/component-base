//go:build jsoniter

package json

import jsoniter "github.com/json-iterator/go"

// RawMessage is exported by component-base/json package.
type RawMessage = jsoniter.RawMessage

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
	// Marshal is exported by component-base/json package.
	Marshal = json.Marshal
	// Unmarshal is exported by component-base/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by component-base/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by component-base/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by component-base/json package.
	NewEncoder = json.NewEncoder
)
