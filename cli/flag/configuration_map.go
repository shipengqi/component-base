package flag

import (
	"fmt"
	"sort"
	"strings"
)

type ConfigurationMap map[string]string

func (m *ConfigurationMap) String() string {
	var pairs []string
	for k, v := range *m {
		pairs = append(pairs, fmt.Sprintf("%s=%s", k, v))
	}
	sort.Strings(pairs)
	return strings.Join(pairs, ",")
}

func (m *ConfigurationMap) Set(value string) error {
	for _, s := range strings.Split(value, ",") {
		if len(s) == 0 {
			continue
		}
		arr := strings.SplitN(s, "=", 2)
		if len(arr) == 2 {
			(*m)[strings.TrimSpace(arr[0])] = strings.TrimSpace(arr[1])
		} else {
			(*m)[strings.TrimSpace(arr[0])] = ""
		}
	}
	return nil
}

func (*ConfigurationMap) Type() string {
	return "mapStringString"
}
