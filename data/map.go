//Data provides extended data operations over maps and slices
package data

//makes union of two maps. If existing map and added map contains
//same key than value of existing is replaced by value of added
func Union(existing map[string]string, added map[string]string) (labels map[string]string) {
	newMap := make(map[string]string, len(existing)+len(added))
	for key, value := range existing {
		newMap[key] = value
	}
	for key, value := range added {
		newMap[key] = value
	}
	return newMap
}
