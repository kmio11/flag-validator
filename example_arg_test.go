package flagvalidator_test

import (
	"flag"
	"fmt"

	"github.com/asaskevich/govalidator"
	fv "github.com/kmio11/flag-validator"
)

func ExampleArgument_Is() {
	fs := flag.NewFlagSet("fs", flag.ContinueOnError)

	ruleSet := fv.NewRuleSet(
		fv.Arg(0).Is(fv.EqualTo("value_expected")),
	)

	_ = fs.Parse([]string{"arg1"})
	err := ruleSet.Validate(fs)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// 0th argument must be a valid value
}

func ExampleArgument_Is_govalidator() {
	fs := flag.NewFlagSet("fs", flag.ContinueOnError)

	ruleSet := fv.NewRuleSet(
		// you can use govalidator.IsXxx functions
		fv.Arg(0).Is(govalidator.IsEmail),
		fv.Arg(1).Is(govalidator.IsNumeric),
	)

	_ = fs.Parse([]string{
		"email@example.com",
		"invalid",
	})
	err := ruleSet.Validate(fs)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// 1th argument must be a valid value
}
