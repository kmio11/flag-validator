package flagvalidator

import (
	"bytes"
	"flag"
	"html/template"
)

var defaultErrorMessages = map[string]string{
	"NumberOfArgs":          "this commands needs {{.want}} arguments but specified {{.actual}} arguments : {{.actual_args}}",
	"MutuallyExclusive":     "exclusive flags are specified at the same time",
	"Flags_Required":        "{{.subject}} {{.required_flags}} is required",
	"Value_Is":              "The value of {{.flag_name}} must be a valid value",
	"IntValue_Min":          "The value of {{.flag_name}} must be greater than or equal to {{.min}}",
	"IntValue_ExclusiveMin": "The value of {{.flag_name}} must be greater than {{.min}}",
	"IntValue_Max":          "The value of {{.flag_name}} must be less than or equal to {{.max}}",
	"IntValue_ExclusiveMax": "The value of {{.flag_name}} must be less than {{.max}}",
	"Argument_Is":           "{{.arg_index}}th argument must be a valid value",
	"Argument_In":           "{{.arg_index}}th argument must be a valid value",
	"nth_argument_not_set":  "{{.arg_index}}th argument is not set",
}

func OverwriteDefaultMessages(messages map[string]string) {
	defaultErrorMessages = merge(defaultErrorMessages, messages)
}

func getMessage(key string) string {
	if v, ok := defaultErrorMessages[key]; ok {
		return v
	}
	return "validation error"
}

type (
	Error interface {
		error
		SetMessage(message string) Error
		AddParams(params map[string]any) Error
	}

	ValidationError struct {
		message string
		params  map[string]any
	}
)

var _ (Error) = (*ValidationError)(nil)

func NewError(message string, params ...map[string]any) Error {
	return ValidationError{
		message: message,
		params:  merge(params...),
	}
}

// WrapError returns Error. When err is Error, return err as-is.
func WrapError(err error) Error {
	if v, ok := err.(Error); ok {
		return v
	}
	return NewError(err.Error(), map[string]any{
		"error": err,
	})
}

func (e ValidationError) Error() string {
	if len(e.params) == 0 {
		return e.message
	}

	errMessage := new(bytes.Buffer)
	_ = template.Must(template.New("err").Parse(e.message)).Execute(errMessage, e.params)

	return errMessage.String()
}

func (e ValidationError) SetMessage(message string) Error {
	return ValidationError{
		message: message,
		params:  e.params,
	}
}

func (a ValidationError) AddParams(params map[string]any) Error {
	if a.params == nil {
		a.params = params
		return a
	}
	a.params = merge(a.params, params)
	return a
}

func presetParams(fs *flag.FlagSet) map[string]any {
	return map[string]any{
		"actual_flags": actualFlagNames(fs),
		"actual_args":  fs.Args(),
	}
}

func merge[K comparable, V any](m ...map[K]V) map[K]V {
	merged := map[K]V{}
	for _, c := range m {
		for k, v := range c {
			merged[k] = v
		}
	}
	return merged
}