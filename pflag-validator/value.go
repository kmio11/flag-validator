package pflagvalidator

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

func (t typed[T]) value(fs *flag.FlagSet) T {
	var val any
	var err error
	switch t.org.value(fs).Type() {
	case "bool":
		panic(fmt.Sprintf("not implemented : %s", t.org.value(fs).Type()))
	case "int":
		val, err = fs.GetInt(t.org.flagName)
		if err != nil {
			panic(err)
		}
	case "int64":
		panic(fmt.Sprintf("not implemented : %s", t.org.value(fs).Type()))
	case "uint":
		panic(fmt.Sprintf("not implemented : %s", t.org.value(fs).Type()))
	case "uint64":
		panic(fmt.Sprintf("not implemented : %s", t.org.value(fs).Type()))
	case "string":
		panic(fmt.Sprintf("not implemented : %s", t.org.value(fs).Type()))
	default:
		panic(fmt.Sprintf("not implemented : %s", t.org.value(fs).Type()))
	}
	vv, ok := any(val).(T)
	if !ok {
		panic(fmt.Sprintf("the value of the flag %s is unexpected type %T", t.org.flagName, val))
	}
	return vv
}
