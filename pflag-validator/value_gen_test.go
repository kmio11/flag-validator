// Code generated by flag-validator; DO NOT EDIT.
package pflagvalidator_test

import (
	"fmt"
	"testing"

	fv "github.com/kmio11/flag-validator/pflag-validator"
	flag "github.com/spf13/pflag"
)

func TestTypeInt_Min(t *testing.T) {
	tests := []struct {
		args       []string
		min        int
		defaultVal int
		err        string
	}{
		{
			args: []string{"--flag-int", "10"},
			min:  5,
			err:  "",
		},
		{
			args: []string{"--flag-int", "5"},
			min:  5,
			err:  "",
		},
		{
			args: []string{"--flag-int", "4"},
			min:  5,
			err:  "The value of flag-int must be greater than or equal to 5",
		},
		{
			args: []string{"--flag-int", "5"},
			min:  10,
			err:  "The value of flag-int must be greater than or equal to 10",
		},
		{
			args:       []string{},
			min:        5,
			defaultVal: 5,
			err:        "",
		},
		{
			args:       []string{},
			min:        5,
			defaultVal: 4,
			err:        "The value of flag-int must be greater than or equal to 5",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			fs := flag.NewFlagSet("test", flag.ContinueOnError)
			fs.String("flag1", "", "usage")
			fs.Int("flag-int", tt.defaultVal, "usage")

			ruleSet := fv.NewRuleSet(
				fv.ValueOf("flag-int").TypeInt().Min(tt.min),
			)

			err := fs.Parse(tt.args)
			if err != nil {
				t.Fatal(err)
			}

			err = ruleSet.Validate(fs)
			assertError(t, tt.err, err)
		})
	}
}

func TestTypeInt_ExclusiveMin(t *testing.T) {
	tests := []struct {
		args       []string
		min        int
		defaultVal int
		err        string
	}{
		{
			args: []string{"--flag-int", "10"},
			min:  5,
			err:  "",
		},
		{
			args: []string{"--flag-int", "6"},
			min:  5,
			err:  "",
		},
		{
			args: []string{"--flag-int", "5"},
			min:  5,
			err:  "The value of flag-int must be greater than 5",
		},
		{
			args: []string{"--flag-int", "4"},
			min:  5,
			err:  "The value of flag-int must be greater than 5",
		},
		{
			args: []string{"--flag-int", "5"},
			min:  10,
			err:  "The value of flag-int must be greater than 10",
		},
		{
			args:       []string{},
			min:        5,
			defaultVal: 6,
			err:        "",
		},
		{
			args:       []string{},
			min:        5,
			defaultVal: 5,
			err:        "The value of flag-int must be greater than 5",
		},
		{
			args:       []string{},
			min:        5,
			defaultVal: 4,
			err:        "The value of flag-int must be greater than 5",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			fs := flag.NewFlagSet("test", flag.ContinueOnError)
			fs.String("flag1", "", "usage")
			fs.Int("flag-int", tt.defaultVal, "usage")

			ruleSet := fv.NewRuleSet(
				fv.ValueOf("flag-int").TypeInt().ExclusiveMin(tt.min),
			)

			err := fs.Parse(tt.args)
			if err != nil {
				t.Fatal(err)
			}

			err = ruleSet.Validate(fs)
			assertError(t, tt.err, err)
		})
	}
}

func TestTypeInt_Max(t *testing.T) {
	tests := []struct {
		args       []string
		max        int
		defaultVal int
		err        string
	}{
		{
			args: []string{"--flag-int", "10"},
			max:  15,
			err:  "",
		},
		{
			args: []string{"--flag-int", "15"},
			max:  15,
			err:  "",
		},
		{
			args: []string{"--flag-int", "16"},
			max:  15,
			err:  "The value of flag-int must be less than or equal to 15",
		},
		{
			args: []string{"--flag-int", "21"},
			max:  20,
			err:  "The value of flag-int must be less than or equal to 20",
		},
		{
			args:       []string{},
			max:        15,
			defaultVal: 15,
			err:        "",
		},
		{
			args:       []string{},
			max:        15,
			defaultVal: 16,
			err:        "The value of flag-int must be less than or equal to 15",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			fs := flag.NewFlagSet("test", flag.ContinueOnError)
			fs.String("flag1", "", "usage")
			fs.Int("flag-int", tt.defaultVal, "usage")

			ruleSet := fv.NewRuleSet(
				fv.ValueOf("flag-int").TypeInt().Max(tt.max),
			)

			err := fs.Parse(tt.args)
			if err != nil {
				t.Fatal(err)
			}

			err = ruleSet.Validate(fs)
			assertError(t, tt.err, err)
		})
	}
}

func TestTypeInt_ExclusiveMax(t *testing.T) {
	tests := []struct {
		args       []string
		max        int
		defaultVal int
		err        string
	}{
		{
			args: []string{"--flag-int", "10"},
			max:  15,
			err:  "",
		},
		{
			args: []string{"--flag-int", "14"},
			max:  15,
			err:  "",
		},
		{
			args: []string{"--flag-int", "15"},
			max:  15,
			err:  "The value of flag-int must be less than 15",
		},
		{
			args: []string{"--flag-int", "16"},
			max:  15,
			err:  "The value of flag-int must be less than 15",
		},
		{
			args: []string{"--flag-int", "20"},
			max:  20,
			err:  "The value of flag-int must be less than 20",
		},
		{
			args:       []string{},
			max:        15,
			defaultVal: 14,
			err:        "",
		},
		{
			args:       []string{},
			max:        15,
			defaultVal: 15,
			err:        "The value of flag-int must be less than 15",
		},
		{
			args:       []string{},
			max:        15,
			defaultVal: 16,
			err:        "The value of flag-int must be less than 15",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			fs := flag.NewFlagSet("test", flag.ContinueOnError)
			fs.String("flag1", "", "usage")
			fs.Int("flag-int", tt.defaultVal, "usage")

			ruleSet := fv.NewRuleSet(
				fv.ValueOf("flag-int").TypeInt().ExclusiveMax(tt.max),
			)

			err := fs.Parse(tt.args)
			if err != nil {
				t.Fatal(err)
			}

			err = ruleSet.Validate(fs)
			assertError(t, tt.err, err)
		})
	}
}
