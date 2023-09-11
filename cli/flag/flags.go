package flag

import (
	goflag "flag"
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

var underscoreWarnings = make(map[string]bool)

type PrintLogger interface {
	Printf(template string, args ...interface{})
}

// WordSepNormalizeFunc changes all flags that contain "_" separators.
func WordSepNormalizeFunc(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.ReplaceAll(name, "_", "-"))
	}
	return pflag.NormalizedName(name)
}

// WarnWordSepNormalizeFunc changes and warns for flags that contain "_" separators.
func WarnWordSepNormalizeFunc(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		nname := strings.ReplaceAll(name, "_", "-")
		if _, alreadyWarned := underscoreWarnings[name]; !alreadyWarned {
			fmt.Printf("using an underscore in a flag name is not supported. %s has been converted to %s.\n", name, nname)
			underscoreWarnings[name] = true
		}

		return pflag.NormalizedName(nname)
	}
	return pflag.NormalizedName(name)
}

// InitFlags normalizes, parses the command line flags.
func InitFlags(flags *pflag.FlagSet) {
	flags.SetNormalizeFunc(WordSepNormalizeFunc)
	flags.AddGoFlagSet(goflag.CommandLine)
}

// PrintFlags logs the flags in the pflag.FlagSet.
func PrintFlags(flags *pflag.FlagSet, logger PrintLogger) {
	flags.VisitAll(func(flag *pflag.Flag) {
		logger.Printf("FLAG: --%s=%q", flag.Name, flag.Value)
	})
}
