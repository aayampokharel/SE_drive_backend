package functions

func RemoveDuplicatesFromList(strList []string) []string {
	// Create a map to track seen elements
	seen := make(map[string]struct{})
	list := []string{}
	for _, value := range strList {
		seen[value] = struct{}{}
	}
	for key, _ := range seen {
		list = append(list, key)
	}
	return list
}
