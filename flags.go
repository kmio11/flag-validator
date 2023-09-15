package flagvalidator

import (
	"flag"
	"fmt"

	"golang.org/x/exp/slices"
)

// Flags represents Flags.
// This struct has the methods returning Rule.
type Flags struct {
	names     []string
	isSetFunc func(fs *flag.FlagSet, names []string) bool
	subject   string
}

func Flag(name string) *Flags {
	return &Flags{
		names: []string{name},
		isSetFunc: func(fs *flag.FlagSet, names []string) bool {
			return slices.Contains(actualFlagNames(fs), name)
		},
		subject: "The flag",
	}
}

func AnyOf(names ...string) *Flags {
	return &Flags{
		names: names,
		isSetFunc: func(fs *flag.FlagSet, names []string) bool {
			actualNames := actualFlagNames(fs)
			for _, name := range names {
				if slices.Contains(actualNames, name) {
					return true
				}
			}
			return false
		},
		subject: "At least one of the flags",
	}
}

func AllOf(names ...string) *Flags {
	return &Flags{
		names: names,
		isSetFunc: func(fs *flag.FlagSet, names []string) bool {
			actualNames := actualFlagNames(fs)
			for _, name := range names {
				if !slices.Contains(actualNames, name) {
					return false
				}
			}
			return true
		},
		subject: "All the flags",
	}
}

func actualFlagNames(fs *flag.FlagSet) []string {
	names := []string{}
	fs.Visit(func(f *flag.Flag) {
		names = append(names, f.Name)
	})
	return names
}

func (f *Flags) IsSet() Condition {
	return ConditionFunc(func(fs *flag.FlagSet) bool {
		return f.isSetFunc(fs, f.names)
	})
}

func (f *Flags) DisplayStrings() []string {
	ret := []string{}
	for _, name := range f.names {
		ret = append(ret, fmt.Sprintf("--%s", name))
	}
	return ret
}

func (f *Flags) Required() Rule {
	return RuleFunc(func(fs *flag.FlagSet) error {
		if len(f.names) == 0 {
			return nil
		}
		if !f.IsSet().Evaluate(fs) {
			return NewError(
				getMessage("Flags_Required"),
				presetParams(fs),
				map[string]any{
					"subject":        f.subject,
					"required_flags": f.DisplayStrings(),
				},
			)
		}
		return nil
	})
}
