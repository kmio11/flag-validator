package flagvalidator

type (
	IsFunc func(value string) bool
)

func EqualTo(want string) IsFunc {
	return func(value string) bool {
		return want == value
	}
}
