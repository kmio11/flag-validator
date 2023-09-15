package flagvalidator

import (
	"flag"

	"golang.org/x/exp/slices"
)

// Argument represents non-flag arguments.
type Argument struct {
	index int
}

// Arg represents the index'th non-flag arguments.
func Arg(index int) *Argument {
	return &Argument{index: index}
}

func (a *Argument) value(fs *flag.FlagSet) string {
	return fs.Arg(a.index)
}

func (a *Argument) IsSet() Condition {
	return ConditionFunc(func(fs *flag.FlagSet) bool {
		return a.value(fs) != ""
	})
}

func (a *Argument) isSet(fs *flag.FlagSet) error {
	if !a.IsSet().Evaluate(fs) {
		return NewError(
			getMessage("nth_argument_not_set"),
			presetParams(fs),
			map[string]any{
				"arg_index": a.index,
			},
		)
	}
	return nil
}

// Is returns a validation rule that checks if a value is a valid value.
func (a *Argument) Is(is IsFunc) Rule {
	return RuleFunc(func(fs *flag.FlagSet) error {
		if err := a.isSet(fs); err != nil {
			return err
		}
		if ok := is(a.value(fs)); !ok {
			return NewError(
				getMessage("Argument_Is"),
				presetParams(fs),
				map[string]any{
					"arg_index": a.index,
				},
			)
		}
		return nil
	})
}

// In returns a validation rule that checks if a value is in the given list of values.
func (a *Argument) In(values ...string) Rule {
	return RuleFunc(func(fs *flag.FlagSet) error {
		if err := a.isSet(fs); err != nil {
			return err
		}
		actual := a.value(fs)
		ok := slices.Contains(values, actual)
		if !ok {
			return NewError(
				getMessage("Argument_In"),
				presetParams(fs),
				map[string]any{
					"want":      values,
					"arg_index": a.index,
				},
			)
		}
		return nil
	})
}
