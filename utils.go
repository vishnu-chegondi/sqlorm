package sqlorm

/*
Find is useful for finding if element is in array and display index if there.
*/
func Find(key string, slice []string) (int, bool) {
	for i, item := range slice {
		if item == key {
			return i, true
		}
	}
	return -1, false
}
