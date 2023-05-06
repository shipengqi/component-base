package flag

import (
	goflag "flag"
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

// StringSlice implements goflag.Value and plfag.Value,
// and allows set to be invoked repeatedly to accumulate values.
type StringSlice struct {
	value   *[]string
	changed bool
}

func NewStringSlice(s *[]string) *StringSlice {
	return &StringSlice{value: s}
}

var _ goflag.Value = &StringSlice{}
var _ pflag.Value = &StringSlice{}

func (s *StringSlice) String() string {
	if s == nil || s.value == nil {
		return ""
	}
	return strings.Join(*s.value, " ")
}

func (s *StringSlice) Set(val string) error {
	if s.value == nil {
		return fmt.Errorf("no target (nil pointer to []string)")
	}
	if !s.changed {
		*s.value = make([]string, 0)
	}
	*s.value = append(*s.value, val)
	s.changed = true
	return nil
}

func (StringSlice) Type() string {
	return "sliceString"
}