package jsonfeed

func Ptr[T any](in T) *T {
	return &in
}

func validStr(str *string) bool {
	return str != nil && len(*str) > 0
}
