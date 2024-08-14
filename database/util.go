package util

import strconv

func MapContains(m map[string][]string, key string) bool {
	v, exists := m[key]

	intV, err := strconv.Atoi(v[0])
	if err == nil {
		if intV != 0 {
			return true
		} else {
			return false
		}
	}

	if Contains(v, "") {
		return false
	} else if !exists {
		return false
	} else {
		return true
	}

}


func Contains[T comparable](arr []T, c T) bool {
	for _, item := range arr {
		if item == c {
			return true
		}
	}
	return false
}
