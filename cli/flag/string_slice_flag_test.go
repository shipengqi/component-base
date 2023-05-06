package flag

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/spf13/pflag"
)

func TestStringSlice(t *testing.T) {
	tests := []struct {
		args       []string
		def        []string
		expected   []string
		parseError string
		changed    bool
	}{
		{
			args:     []string{},
			expected: nil,
		},
		{
			args:     []string{"a"},
			expected: []string{"a"},
			changed:  true,
		},
		{
			args:     []string{"a", "b"},
			expected: []string{"a", "b"},
			changed:  true,
		},
		{
			def:      []string{"a"},
			args:     []string{"a", "b"},
			expected: []string{"a", "b"},
			changed:  true,
		},
		{
			def:      []string{"a", "b"},
			args:     []string{"a", "b"},
			expected: []string{"a", "b"},
			changed:  true,
		},
	}
	for i, test := range tests {
		fs := pflag.NewFlagSet("testStringSlice", pflag.ContinueOnError)
		var s []string
		s = append(s, test.def...)

		v := NewStringSlice(&s)
		fs.Var(v, "slice", "usage")

		var args []string
		for _, a := range test.args {
			args = append(args, fmt.Sprintf("--slice=%s", a))
		}

		err := fs.Parse(args)
		if test.parseError != "" {
			if err == nil {
				t.Errorf("%d: expected error %q, got nil", i, test.parseError)
			} else if !strings.Contains(err.Error(), test.parseError) {
				t.Errorf("%d: expected error %q, got %q", i, test.parseError, err)
			}
		} else if err != nil {
			t.Errorf("%d: expected nil error, got %v", i, err)
		}
		if !reflect.DeepEqual(s, test.expected) {
			t.Errorf("%d: expected %+v, got %+v", i, test.expected, s)
		}
		if v.changed != test.changed {
			t.Errorf("%d: expected %t got %t", i, test.changed, v.changed)
		}
	}
}
