package dp

import "testing"

func TestIsMatch(t *testing.T) {
	t.Logf("s: %s, p: %s, is match: %t", "aa", "a", isMatch("aa", "a"))
	t.Logf("s: %s, p: %s, is match: %t", "aa", "a*", isMatch("aa", "a*"))
	t.Logf("s: %s, p: %s, is match: %t", "ab", ".*", isMatch("ab", ".*"))
	t.Logf("s: %s, p: %s, is match: %t", "abb", ".*", isMatch("abb", ".*"))
	t.Logf("s: %s, p: %s, is match: %t", "abbc", ".*", isMatch("abbc", ".*"))
	t.Logf("s: %s, p: %s, is match: %t", "abbc", ".b*", isMatch("abbc", ".b*"))
	t.Logf("s: %s, p: %s, is match: %t", "abbc", "..*", isMatch("abbc", "..*"))
	t.Logf("s: %s, p: %s, is match: %t", "abbc", ".**", isMatch("abbc", ".**"))
	t.Logf("s: %s, p: %s, is match: %t", "aab", "c*a*b", isMatch("aab", "c*a*b"))
	t.Logf("s: %s, p: %s, is match: %t", "ab", ".*c", isMatch("ab", ".*c"))
}
