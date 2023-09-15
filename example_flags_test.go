package flagvalidator_test

import (
	"flag"
	"fmt"

	fv "github.com/kmio11/flag-validator"
)

func ExampleFlags_Required() {
	fs := flag.NewFlagSet("fs", flag.ContinueOnError)
	_ = fs.String("flag1", "", "")
	_ = fs.String("flag2", "", "")

	ruleSet := fv.NewRuleSet(
		fv.Flag("flag1").Required(),
	)

	_ = fs.Parse([]string{"--flag2", "value2"})
	err := ruleSet.Validate(fs)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// The flag [--flag1] is required
}

func ExampleFlags_Required_anyOf() {
	fs := flag.NewFlagSet("fs", flag.ContinueOnError)
	_ = fs.String("flag1", "", "")
	_ = fs.String("flag2", "", "")
	_ = fs.String("flag3", "", "")

	ruleSet := fv.NewRuleSet(
		fv.AnyOf("flag1", "flag2").Required(),
	)

	_ = fs.Parse([]string{"--flag3", "value3"})
	err := ruleSet.Validate(fs)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// At least one of the flags [--flag1 --flag2] is required
}

func ExampleFlags_Required_allOf() {
	fs := flag.NewFlagSet("fs", flag.ContinueOnError)
	_ = fs.String("flag1", "", "")
	_ = fs.String("flag2", "", "")
	_ = fs.String("flag3", "", "")

	ruleSet := fv.NewRuleSet(
		fv.AllOf("flag1", "flag2").Required(),
	)

	_ = fs.Parse([]string{"--flag1", "value1"})
	err := ruleSet.Validate(fs)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// All the flags [--flag1 --flag2] is required
}
