package utils

// Must _
func Must(i interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return i
}
