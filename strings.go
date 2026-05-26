package sprig

func base64encode(v string) string { _ = "STUB: not implemented"; return "" }

func base64decode(v string) string { _ = "STUB: not implemented"; return "" }

func base32encode(v string) string { _ = "STUB: not implemented"; return "" }

func base32decode(v string) string { _ = "STUB: not implemented"; return "" }

func abbrev(width int, s string) string { _ = "STUB: not implemented"; return "" }

func abbrevboth(left, right int, s string) string { _ = "STUB: not implemented"; return "" }

func initials(s string) string {
	_ = "STUB: not implemented"
	// Wrap this just to eliminate the var args, which templates don't do well.
	return ""
}

func randAlphaNumeric(count int) string {
	_ = "STUB: not implemented"
	// It is not possible, it appears, to actually generate an error here.
	return ""
}

func randAlpha(count int) string { _ = "STUB: not implemented"; return "" }

func randAscii(count int) string { _ = "STUB: not implemented"; return "" }

func randNumeric(count int) string { _ = "STUB: not implemented"; return "" }

func untitle(str string) string { _ = "STUB: not implemented"; return "" }

func quote(str ...interface{}) string { _ = "STUB: not implemented"; return "" }

func squote(str ...interface{}) string { _ = "STUB: not implemented"; return "" }

func cat(v ...interface{}) string { _ = "STUB: not implemented"; return "" }

func indent(spaces int, v string) string { _ = "STUB: not implemented"; return "" }

func nindent(spaces int, v string) string { _ = "STUB: not implemented"; return "" }

func replace(old, new, src string) string { _ = "STUB: not implemented"; return "" }

func plural(one, many string, count int) string { _ = "STUB: not implemented"; return "" }

func strslice(v interface{}) []string { _ = "STUB: not implemented"; return nil }

func removeNilElements(v []interface{}) []interface{} { _ = "STUB: not implemented"; return nil }

func strval(v interface{}) string { _ = "STUB: not implemented"; return "" }

func trunc(c int, s string) string { _ = "STUB: not implemented"; return "" }

func join(sep string, v interface{}) string { _ = "STUB: not implemented"; return "" }

func split(sep, orig string) map[string]string { _ = "STUB: not implemented"; return nil }

func splitn(sep string, n int, orig string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// substring creates a substring of the given string.
//
// If start is < 0, this calls string[:end].
//
// If start is >= 0 and end < 0 or end bigger than s length, this calls string[start:]
//
// Otherwise, this calls string[start, end].
func substring(start, end int, s string) string { _ = "STUB: not implemented"; return "" }
