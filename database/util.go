package database

func FormToArray(m map[string][]string) [][]string {
	arr := [][]string{}
	for k, v := range m {
		newarr := []string{
			k, v[0],
		}
		arr = append(arr, newarr)
	}
	return arr
}

func Contains[T comparable](arr []T, c T) bool {
	for _, item := range arr {
		if item == c {
			return true
		}
	}
	return false
}
