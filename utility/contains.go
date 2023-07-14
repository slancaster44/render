package utility

func Contains[T comparable](list []T, val T) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}

	return false
}
