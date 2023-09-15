package flagvalidator

import (
	"flag"
	"fmt"
	"time"
)

type (
	// Value represents the value of a flag.
	Value struct {
		flagName string
	}

	valueType interface {
		bool | int | int64 | uint | uint64 | string | time.Duration
	}

	typed[T valueType] struct {
		org Value
	}

	IntValue struct {
		typed[int]
	}
)

func ValueOf(name string) *Value {
	return &Value{
		flagName: name,
	}
}

func (v Value) value(fs *flag.FlagSet) flag.Value {
	f := fs.Lookup(v.flagName)
	if f == nil {
		panic(fmt.Sprintf("the flag %s is not defined", v.flagName))
	}
	return f.Value
}

func (t typed[T]) value(fs *flag.FlagSet) T {
	getter, ok := t.org.value(fs).(flag.Getter)
	if !ok {
		panic("unable to get the value of the flag %s")
	}
	vv, ok := getter.Get().(T)
	if !ok {
		panic(fmt.Sprintf("the value of the flag %s is unexpected type %T", t.org.flagName, getter.Get()))
	}
	return vv
}

// Is returns a validation rule that checks if a value is a valid value.
func (v Value) Is(is IsFunc) Rule {
	return RuleFunc(func(fs *flag.FlagSet) error {
		vStr := v.value(fs).String()
		if ok := is(vStr); !ok {
			return NewError(
				getMessage("Value_Is"),
				presetParams(fs),
				map[string]any{
					"flag_name":    v.flagName,
					"value_string": v.value(fs).String(),
				},
			)
		}
		return nil
	})
}

func (v Value) TypeInt() *IntValue {
	return &IntValue{
		typed: typed[int]{
			org: v,
		},
	}
}

// Min returns a validation rule that checks if a value is less than or equal to the threshold value.
func (v IntValue) Min(min int) Rule {
	return RuleFunc(func(fs *flag.FlagSet) error {
		val := v.typed.value(fs)
		if val < min {
			return NewError(
				getMessage("IntValue_Min"),
				presetParams(fs),
				map[string]any{
					"flag_name": v.org.flagName,
					"value":     val,
					"min":       min,
				},
			)
		}
		return nil
	})
}

// ExclusiveMin returns a validation rule that checks if a value is less than the threshold value.
func (v IntValue) ExclusiveMin(min int) Rule {
	return RuleFunc(func(fs *flag.FlagSet) error {
		val := v.typed.value(fs)
		if val <= min {
			return NewError(
				getMessage("IntValue_ExclusiveMin"),
				presetParams(fs),
				map[string]any{
					"flag_name": v.org.flagName,
					"value":     val,
					"min":       min,
				},
			)
		}
		return nil
	})
}

// Max returns a validation rule that checks if a value is greater than or equal to the threshold value.
func (v IntValue) Max(max int) Rule {
	return RuleFunc(func(fs *flag.FlagSet) error {
		val := v.typed.value(fs)
		if val > max {
			return NewError(
				getMessage("IntValue_Max"),
				presetParams(fs),
				map[string]any{
					"flag_name": v.org.flagName,
					"value":     val,
					"max":       max,
				},
			)
		}
		return nil
	})
}

// ExclusiveMax returns a validation rule that checks if a value is greater than the threshold value.
func (v IntValue) ExclusiveMax(max int) Rule {
	return RuleFunc(func(fs *flag.FlagSet) error {
		val := v.typed.value(fs)
		if val >= max {
			return NewError(
				getMessage("IntValue_ExclusiveMax"),
				presetParams(fs),
				map[string]any{
					"flag_name": v.org.flagName,
					"value":     val,
					"max":       max,
				},
			)
		}
		return nil
	})
}
