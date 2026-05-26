package sprig

// Reflection is used in these functions so that slices and arrays of strings,
// ints, and other types not implementing []interface{} can be worked with.
// For example, this is useful if you need to work on the output of regexs.

func list(v ...interface{}) []interface{} { _ = "STUB: not implemented"; return nil }

func push(list interface{}, v interface{}) []interface{} { _ = "STUB: not implemented"; return nil }

func mustPush(list interface{}, v interface{}) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func prepend(list interface{}, v interface{}) []interface{} { _ = "STUB: not implemented"; return nil }

func mustPrepend(list interface{}, v interface{}) ([]interface{}, error) {
	_ = "STUB: not implemented"
	//return append([]interface{}{v}, list...)
	return nil, nil
}

func chunk(size int, list interface{}) [][]interface{} { _ = "STUB: not implemented"; return nil }

func mustChunk(size int, list interface{}) ([][]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func last(list interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func mustLast(list interface{}) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func first(list interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func mustFirst(list interface{}) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func rest(list interface{}) []interface{} { _ = "STUB: not implemented"; return nil }

func mustRest(list interface{}) ([]interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func initial(list interface{}) []interface{} { _ = "STUB: not implemented"; return nil }

func mustInitial(list interface{}) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sortAlpha(list interface{}) []string { _ = "STUB: not implemented"; return nil }

func reverse(v interface{}) []interface{} { _ = "STUB: not implemented"; return nil }

func mustReverse(v interface{}) ([]interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

// We do not sort in place because the incoming array should not be altered.

func compact(list interface{}) []interface{} { _ = "STUB: not implemented"; return nil }

func mustCompact(list interface{}) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func uniq(list interface{}) []interface{} { _ = "STUB: not implemented"; return nil }

func mustUniq(list interface{}) ([]interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func inList(haystack []interface{}, needle interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func without(list interface{}, omit ...interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func mustWithout(list interface{}, omit ...interface{}) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func has(needle interface{}, haystack interface{}) bool { _ = "STUB: not implemented"; return false }

func mustHas(needle interface{}, haystack interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// $list := [1, 2, 3, 4, 5]
// slice $list     -> list[0:5] = list[:]
// slice $list 0 3 -> list[0:3] = list[:3]
// slice $list 3 5 -> list[3:5]
// slice $list 3   -> list[3:5] = list[3:]
func slice(list interface{}, indices ...interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func mustSlice(list interface{}, indices ...interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concat(lists ...interface{}) interface{} { _ = "STUB: not implemented"; return nil }
