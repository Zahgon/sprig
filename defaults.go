package sprig

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// dfault checks whether `given` is set, and returns default if not set.
//
// This returns `d` if `given` appears not to be set, and `given` otherwise.
//
// For numeric types 0 is unset.
// For strings, maps, arrays, and slices, len() = 0 is considered unset.
// For bool, false is unset.
// Structs are never considered unset.
//
// For everything else, including pointers, a nil value is unset.
func dfault(d interface{}, given ...interface{}) interface{} { _ = "STUB: not implemented"; return nil }

// empty returns true if the given value has the zero value for its type.
func empty(given interface{}) bool { _ = "STUB: not implemented"; return false }

// Basically adapted from text/template.isTrue

// coalesce returns the first non-empty value.
func coalesce(v ...interface{}) interface{} { _ = "STUB: not implemented"; return nil }

// all returns true if empty(x) is false for all values x in the list.
// If the list is empty, return true.
func all(v ...interface{}) bool { _ = "STUB: not implemented"; return false }

// any returns true if empty(x) is false for any x in the list.
// If the list is empty, return false.
func any(v ...interface{}) bool { _ = "STUB: not implemented"; return false }

// fromJson decodes JSON into a structured value, ignoring errors.
func fromJson(v string) interface{} { _ = "STUB: not implemented"; return nil }

// mustFromJson decodes JSON into a structured value, returning errors.
func mustFromJson(v string) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

// toJson encodes an item into a JSON string
func toJson(v interface{}) string { _ = "STUB: not implemented"; return "" }

func mustToJson(v interface{}) (string, error) { _ = "STUB: not implemented"; return "", nil }

// toPrettyJson encodes an item into a pretty (indented) JSON string
func toPrettyJson(v interface{}) string { _ = "STUB: not implemented"; return "" }

func mustToPrettyJson(v interface{}) (string, error) { _ = "STUB: not implemented"; return "", nil }

// toRawJson encodes an item into a JSON string with no escaping of HTML characters.
func toRawJson(v interface{}) string { _ = "STUB: not implemented"; return "" }

// mustToRawJson encodes an item into a JSON string with no escaping of HTML characters.
func mustToRawJson(v interface{}) (string, error) { _ = "STUB: not implemented"; return "", nil }

// ternary returns the first value if the last value is true, otherwise returns the second value.
func ternary(vt interface{}, vf interface{}, v bool) interface{} {
	_ = "STUB: not implemented"
	return nil
}
