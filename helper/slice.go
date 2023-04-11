package helper

func SlicesContain[T comparable](s []T, contain T) bool {
	for _, v := range s {
		if v == contain {
			return true
		}
	}

	return false
}
