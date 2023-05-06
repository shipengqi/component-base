package flag

import (
	goflag "flag"
	"strings"

	"github.com/shipengqi/log"
	"github.com/spf13/pflag"
)

var underscoreWarnings = make(map[string]bool)

// WordSepNormalizeFunc changes all flags that contain "_" separators.
func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.ReplaceAll(name, "_", "-"))
	}
	return pflag.NormalizedName(name)
}

// WarnWordSepNormalizeFunc changes and warns for flags that contain "_" separators.
func WarnWordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		nname := strings.ReplaceAll(name, "_", "-")
		if _, alreadyWarned := underscoreWarnings[name]; !alreadyWarned {
			log.Warnf("using an underscore in a flag name is not supported. %s has been converted to %s.", name, nname)
			underscoreWarnings[name] = true
		}

		return pflag.NormalizedName(nname)
	}
	return pflag.NormalizedName(name)
}

// InitFlags normalizes, parses, then logs the command line flags.
func InitFlags(flags *pflag.FlagSet) {
	flags.SetNormalizeFunc(WarnWordSepNormalizeFunc)
	flags.AddGoFlagSet(goflag.CommandLine)
}

// PrintFlags logs the flags in the pflag.FlagSet.
func PrintFlags(flags *pflag.FlagSet) {
	flags.VisitAll(func(flag *pflag.Flag) {
		log.Debugf("FLAG: --%s=%q", flag.Name, flag.Value)
	})
}
