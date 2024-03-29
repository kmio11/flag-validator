package flagvalidator

import (
	"flag"
)

type (
	Rule interface {
		Validate(fs *flag.FlagSet) error
		Error(message string) Rule
	}

	RuleFunc func(fs *flag.FlagSet) error

	RuleSet struct {
		rules []Rule
	}
)

func (v RuleFunc) Validate(fs *flag.FlagSet) error {
	return v(fs)
}

func (v RuleFunc) Error(message string) Rule {
	return errorRuleFunc(v, message)
}

// errorRuleFunc is processing for Rule.Error() implementation.
func errorRuleFunc(r Rule, message string) Rule {
	return RuleFunc(func(fs *flag.FlagSet) error {
		err := r.Validate(fs)
		if err != nil {
			return WrapError(err).SetMessage(message)
		}
		return nil
	})
}

func NewRuleSet(rules ...Rule) *RuleSet {
	return &RuleSet{
		rules: rules,
	}
}

func (rs *RuleSet) Validate(fs *flag.FlagSet) error {
	for _, rule := range rs.rules {
		err := rule.Validate(fs)
		if err != nil {
			return err
		}
	}
	return nil
}

// NumberOfArgs returns the validation rule that checks if
// the number of specified argments is equal to n.
func NumberOfArgs(n int) Rule {
	return RuleFunc(func(fs *flag.FlagSet) error {
		if fs.NArg() != n {
			return NewError(
				getMessage("NumberOfArgs"),
				presetParams(fs),
				map[string]any{
					"want":   n,
					"actual": fs.NArg(),
				},
			)
		}
		return nil
	})
}

// MutuallyExclusive returns the validation rule that checks if
// the given list of flags are not specified more than once.
func MutuallyExclusive(flags ...*Flags) Rule {
	return RuleFunc(func(fs *flag.FlagSet) error {
		var isOtherGrouopSet bool
		for _, f := range flags {
			isSet := f.IsSet().Evaluate(fs)
			if isOtherGrouopSet && isSet {
				return NewError(
					getMessage("MutuallyExclusive"),
					presetParams(fs),
				)
			}
			isOtherGrouopSet = isOtherGrouopSet || isSet
		}
		return nil
	})
}
