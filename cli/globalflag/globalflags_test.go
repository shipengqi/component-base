package globalflag

import (
	"flag"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/spf13/pflag"

	cliflag "github.com/shipengqi/component-base/cli/flag"
)

func TestAddGlobalFlags(t *testing.T) {
	namedFlagSets := &cliflag.NamedFlagSets{}
	nfs := namedFlagSets.FlagSet("global")
	nfs.SetNormalizeFunc(cliflag.WordSepNormalizeFunc)
	AddGlobalFlags(nfs, "test-cmd")

	var actualFlag []string
	nfs.VisitAll(func(flag *pflag.Flag) {
		actualFlag = append(actualFlag, flag.Name)
	})

	// Get all flags from flags.CommandLine, except flag `test.*`.
	wantedFlag := []string{"help"}
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	normalizeFunc := nfs.GetNormalizeFunc()
	pflag.VisitAll(func(flag *pflag.Flag) {
		if !strings.Contains(flag.Name, "test.") {
			wantedFlag = append(wantedFlag, string(normalizeFunc(nfs, flag.Name)))
		}
	})
	sort.Strings(wantedFlag)

	if !reflect.DeepEqual(wantedFlag, actualFlag) {
		t.Errorf("[Default]: expected %+v, got %+v", wantedFlag, actualFlag)
	}

	tests := []struct {
		expectedFlag  []string
		matchExpected bool
	}{
		{
			// Happy case
			expectedFlag:  []string{"help"},
			matchExpected: false,
		},
		{
			// Missing flag
			expectedFlag:  []string{"logtostderr", "log-dir"},
			matchExpected: true,
		},
		{
			// Empty flag
			expectedFlag:  []string{},
			matchExpected: true,
		},
		{
			// Invalid flag
			expectedFlag:  []string{"foo"},
			matchExpected: true,
		},
	}

	for i, test := range tests {
		if reflect.DeepEqual(test.expectedFlag, actualFlag) == test.matchExpected {
			t.Errorf("[%d]: expected %+v, got %+v", i, test.expectedFlag, actualFlag)
		}
	}
}
