package util

func Unique(list []int) []int {
	tmp := map[int]bool{}
	result := []int{}
	for _, v := range list {
		tmp[v] = true
	}
	for k, v := range tmp {
		if v {
			result = append(result, k)
		}
	}
	return result
}

func Contains(list []int, target int) bool {
	for _, v := range list {
		if v == target {
			return true
		}
	}
	return false
}
