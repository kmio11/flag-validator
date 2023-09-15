package flagvalidator_test

import (
	"fmt"

	"flag"
	"testing"

	fv "github.com/kmio11/flag-validator"
	"github.com/stretchr/testify/assert"
)

func assertError(t *testing.T, expected string, err error) {
	if expected == "" {
		assert.Nil(t, err)
	} else if assert.NotNil(t, err) {
		assert.Equal(t, expected, err.Error())
	}
}

func TestNumberOfArgs(t *testing.T) {
	tests := []struct {
		num  int
		args []string
		err  string
	}{
		{
			num:  0,
			args: []string{},
			err:  "",
		},
		{
			num:  0,
			args: []string{"--flag1", "value1"},
			err:  "",
		},
		{
			num:  0,
			args: []string{"arg1"},
			err:  "this commands needs 0 arguments but specified 1 arguments : [arg1]",
		},
		{
			num:  1,
			args: []string{"arg1"},
			err:  "",
		},
		{
			num:  1,
			args: []string{"--flag1", "value1", "arg1"},
			err:  "",
		},
		{
			num:  1,
			args: []string{},
			err:  "this commands needs 1 arguments but specified 0 arguments : []",
		},
		{
			num:  1,
			args: []string{"arg1", "arg2"},
			err:  "this commands needs 1 arguments but specified 2 arguments : [arg1 arg2]",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			fs := flag.NewFlagSet("test", flag.ContinueOnError)
			fs.String("flag1", "", "usage")

			ruleSet := fv.NewRuleSet(
				fv.NumberOfArgs(tt.num),
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

func TestFlag_Required(t *testing.T) {
	tests := []struct {
		name string
		args []string
		err  string
	}{
		{
			name: "flag1",
			args: []string{"--flag1", "value1"},
			err:  "",
		},
		{
			name: "flag1",
			args: []string{"--flag2", "value1"},
			err:  "The flag [--flag1] is required",
		},
		{
			name: "flag1",
			args: []string{},
			err:  "The flag [--flag1] is required",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			fs := flag.NewFlagSet("test", flag.ContinueOnError)
			fs.String("flag1", "", "usage")
			fs.String("flag2", "", "usage")
			fs.String("flag3", "", "usage")

			ruleSet := fv.NewRuleSet(
				fv.Flag(tt.name).Required(),
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

func TestAnyOf_Required(t *testing.T) {
	tests := []struct {
		names []string
		args  []string
		err   string
	}{
		{
			names: []string{"flag1"},
			args:  []string{"--flag1", "value1"},
			err:   "",
		},
		{
			names: []string{"flag1", "flag2"},
			args:  []string{"--flag1", "value1"},
			err:   "",
		},
		{
			names: []string{"flag1", "flag2"},
			args:  []string{"--flag1", "value1", "--flag2", "value2"},
			err:   "",
		},
		{
			names: []string{"flag3"},
			args:  []string{"--flag1", "value1", "--flag2", "value2"},
			err:   "At least one of the flags [--flag3] is required",
		},
		{
			names: []string{"flag3", "flag4"},
			args:  []string{"--flag1", "value1", "--flag2", "value2"},
			err:   "At least one of the flags [--flag3 --flag4] is required",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			fs := flag.NewFlagSet("test", flag.ContinueOnError)
			fs.String("flag1", "", "usage")
			fs.String("flag2", "", "usage")
			fs.String("flag3", "", "usage")
			fs.String("flag4", "", "usage")

			ruleSet := fv.NewRuleSet(
				fv.AnyOf(tt.names...).Required(),
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

func TestAllOf_Required(t *testing.T) {
	tests := []struct {
		names []string
		args  []string
		err   string
	}{
		{
			names: []string{"flag1"},
			args:  []string{"--flag1", "value1"},
			err:   "",
		},
		{
			names: []string{"flag1", "flag2"},
			args:  []string{"--flag1", "value1"},
			err:   "All the flags [--flag1 --flag2] is required",
		},
		{
			names: []string{"flag1", "flag2"},
			args:  []string{"--flag1", "value1", "--flag2", "value2"},
			err:   "",
		},
		{
			names: []string{"flag3"},
			args:  []string{"--flag1", "value1", "--flag2", "value2"},
			err:   "All the flags [--flag3] is required",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			fs := flag.NewFlagSet("test", flag.ContinueOnError)
			fs.String("flag1", "", "usage")
			fs.String("flag2", "", "usage")
			fs.String("flag3", "", "usage")
			fs.String("flag4", "", "usage")

			ruleSet := fv.NewRuleSet(
				fv.AllOf(tt.names...).Required(),
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
