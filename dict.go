package sprig

func get(d map[string]interface{}, key string) interface{} { _ = "STUB: not implemented"; return nil }

func set(d map[string]interface{}, key string, value interface{}) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func unset(d map[string]interface{}, key string) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func hasKey(d map[string]interface{}, key string) bool { _ = "STUB: not implemented"; return false }

func pluck(key string, d ...map[string]interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func keys(dicts ...map[string]interface{}) []string { _ = "STUB: not implemented"; return nil }

func pick(dict map[string]interface{}, keys ...string) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func omit(dict map[string]interface{}, keys ...string) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func dict(v ...interface{}) map[string]interface{} { _ = "STUB: not implemented"; return nil }

func merge(dst map[string]interface{}, srcs ...map[string]interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// Swallow errors inside of a template.

func mustMerge(dst map[string]interface{}, srcs ...map[string]interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mergeOverwrite(dst map[string]interface{}, srcs ...map[string]interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// Swallow errors inside of a template.

func mustMergeOverwrite(dst map[string]interface{}, srcs ...map[string]interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func values(dict map[string]interface{}) []interface{} { _ = "STUB: not implemented"; return nil }

func deepCopy(i interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func mustDeepCopy(i interface{}) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func dig(ps ...interface{}) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func digFromDict(dict map[string]interface{}, d interface{}, ks []string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
