package flagvalidator_test

import (
	"flag"
	"fmt"

	fv "github.com/kmio11/flag-validator"
)

func ExampleIntValue_Max() {
	fs := flag.NewFlagSet("fs", flag.ContinueOnError)
	_ = fs.Int("flag1", 0, "")

	ruleSet := fv.NewRuleSet(
		fv.ValueOf("flag1").TypeInt().Max(10),
	)

	_ = fs.Parse([]string{"--flag1", "15"})
	err := ruleSet.Validate(fs)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// The value of flag1 must be less than or equal to 10
}
