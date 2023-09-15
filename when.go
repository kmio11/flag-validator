package flagvalidator

import (
	"flag"
)

// When validates the given list of rules when the condition is true.
func When(condition Condition) func(rules ...Rule) Rule {
	return func(rules ...Rule) Rule {
		return &whenRule{
			condition: condition,
			rules:     rules,
		}
	}
}

var _ Rule = (*whenRule)(nil)

type whenRule struct {
	condition Condition
	rules     []Rule
}

func (r *whenRule) Validate(fs *flag.FlagSet) error {
	if !r.condition.Evaluate(fs) {
		return nil
	}
	for _, rule := range r.rules {
		err := rule.Validate(fs)
		if err != nil {
			return WrapError(err)
		}
	}
	return nil
}

func (r *whenRule) Error(message string) Rule {
	return errorRuleFunc(r, message)
}

type (
	Condition interface {
		Evaluate(fs *flag.FlagSet) bool
	}

	ConditionFunc func(fs *flag.FlagSet) bool

	ConditionBool bool
)

func (c ConditionFunc) Evaluate(fs *flag.FlagSet) bool {
	return c(fs)
}

func (c ConditionBool) Evaluate(fs *flag.FlagSet) bool {
	return bool(c)
}

func Not(condition Condition) Condition {
	return ConditionFunc(func(fs *flag.FlagSet) bool {
		return !condition.Evaluate(fs)
	})
}
