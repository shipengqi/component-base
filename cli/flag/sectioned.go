package flag

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var templateFuncs = template.FuncMap{
	"trim": strings.TrimSpace,
	"rpad": rpad,
	"gt":   cobra.Gt,
	"eq":   cobra.Eq,
}

const (
	usageFmt   = "Usage:\n  %s\n"
	aliasesFmt = `{{if gt (len .Aliases) 0}}
Aliases:
  {{.NameAndAliases}}
{{end}}`
	commandsFmt = `{{if .HasAvailableSubCommands}}
Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}
{{end}}`
	examplesFmt = `{{if .HasExample}}
Examples:
  {{.Example}}
{{end}}`
	moreFmt = `{{if .HasAvailableSubCommands}}
Use "{{.CommandPath}} [command] --help" for more information about a command.
{{end}}`
)

// NamedFlagSets stores named flag sets in the order of calling FlagSet.
type NamedFlagSets struct {
	// Order is an ordered list of flag set names.
	Order []string
	// FlagSets stores the flag sets by name.
	FlagSets map[string]*pflag.FlagSet
	// NormalizeNameFunc is the normalize function which used to initialize FlagSets created by NamedFlagSets.
	NormalizeNameFunc func(f *pflag.FlagSet, name string) pflag.NormalizedName
}

// FlagSet returns the flag set with the given name and adds it to the
// ordered name list if it is not in there yet.
func (nfs *NamedFlagSets) FlagSet(name string) *pflag.FlagSet {
	if nfs.FlagSets == nil {
		nfs.FlagSets = map[string]*pflag.FlagSet{}
	}
	if _, ok := nfs.FlagSets[name]; !ok {
		flagSet := pflag.NewFlagSet(name, pflag.ExitOnError)
		flagSet.SetNormalizeFunc(pflag.CommandLine.GetNormalizeFunc())
		if nfs.NormalizeNameFunc != nil {
			flagSet.SetNormalizeFunc(nfs.NormalizeNameFunc)
		}
		nfs.FlagSets[name] = flagSet
		nfs.Order = append(nfs.Order, name)
	}
	return nfs.FlagSets[name]
}

// PrintSections prints the given names flag sets in sections, with the maximal given column number.
// If cols is zero, lines are not wrapped.
func PrintSections(w io.Writer, fss NamedFlagSets, cols int) {
	for _, name := range fss.Order {
		fs := fss.FlagSets[name]
		if !fs.HasFlags() {
			continue
		}

		wideFS := pflag.NewFlagSet("", pflag.ExitOnError)
		wideFS.AddFlagSet(fs)

		var zzz string
		if cols > 24 {
			zzz = strings.Repeat("z", cols-24)
			wideFS.Int(zzz, 0, strings.Repeat("z", cols-24))
		}

		var buf bytes.Buffer
		_, _ = fmt.Fprintf(&buf, "\n%s flags:\n%s", strings.ToUpper(name[:1])+name[1:], wideFS.FlagUsagesWrapped(cols))

		if cols > 24 {
			i := strings.Index(buf.String(), zzz)
			lines := strings.Split(buf.String()[:i], "\n")
			_, _ = fmt.Fprint(w, strings.Join(lines[:len(lines)-1], "\n"))
			_, _ = fmt.Fprintln(w)
		} else {
			_, _ = fmt.Fprint(w, buf.String())
		}
	}
}

// PrintAliases prints the aliases.
func PrintAliases(w io.Writer, cmd *cobra.Command) {
	_ = tmpl(w, aliasesFmt, cmd)
}

// PrintSubCommands prints the sub commands.
func PrintSubCommands(w io.Writer, cmd *cobra.Command) {
	_ = tmpl(w, commandsFmt, cmd)
}

// PrintExamples prints the examples.
func PrintExamples(w io.Writer, cmd *cobra.Command) {
	_ = tmpl(w, examplesFmt, cmd)
}

// PrintMore prints the more information.
func PrintMore(w io.Writer, cmd *cobra.Command) {
	_ = tmpl(w, moreFmt, cmd)
}

// SetUsageAndHelpFunc set both usage and help function.
// Print the flag sets we need instead of all of them.
func SetUsageAndHelpFunc(cmd *cobra.Command, fss NamedFlagSets, cols int) {
	cmd.SetUsageFunc(func(cmd *cobra.Command) error {
		_, _ = fmt.Fprintf(cmd.OutOrStderr(), usageFmt, cmd.UseLine())
		PrintAliases(cmd.OutOrStderr(), cmd)
		PrintSubCommands(cmd.OutOrStderr(), cmd)
		PrintSections(cmd.OutOrStderr(), fss, cols)
		PrintExamples(cmd.OutOrStderr(), cmd)
		PrintMore(cmd.OutOrStderr(), cmd)
		return nil
	})
	cmd.SetHelpFunc(func(cmd *cobra.Command, _ []string) {
		_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usageFmt, cmd.Long, cmd.UseLine())
		PrintAliases(cmd.OutOrStderr(), cmd)
		PrintSubCommands(cmd.OutOrStderr(), cmd)
		PrintSections(cmd.OutOrStderr(), fss, cols)
		PrintExamples(cmd.OutOrStderr(), cmd)
		PrintMore(cmd.OutOrStderr(), cmd)
	})
}

// tmpl executes the given template text on data, writing the result to w.
func tmpl(w io.Writer, text string, data interface{}) error {
	t := template.New("top")
	t.Funcs(templateFuncs)
	template.Must(t.Parse(text))
	return t.Execute(w, data)
}

// rpad adds padding to the right of a string.
func rpad(s string, padding int) string {
	formattedString := fmt.Sprintf("%%-%ds", padding)
	return fmt.Sprintf(formattedString, s)
}
