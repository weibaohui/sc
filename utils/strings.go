package utils

func InArray(key string, array []string) bool {
	if len(key) == 0 {
		return false
	}
	for _, v := range array {
		if key == v {
			return true
		}
	}
	return false
}
