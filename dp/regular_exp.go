package dp

// 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
// '.' 匹配任意单个字符
// '*' 匹配零个或多个前面的那一个元素
// 所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

func isMatch(s string, p string) bool {
	return doIsMatch(0, s, 0, p)
}

func doIsMatch(si int, s string, pi int, p string) bool {
	if pi >= len(p) {
		return si >= len(s)
	}

	match := si < len(s) && (s[si] == p[pi] || p[pi] == '.')

	if pi+1 < len(p) && p[pi+1] == '*' {
		if match {
			// 匹配若干次或0次
			return doIsMatch(si+1, s, pi, p) || doIsMatch(si, s, pi+2, p)
		}
		// 匹配0次
		return doIsMatch(si, s, pi+2, p)
	}

	if match {
		return doIsMatch(si+1, s, pi+1, p)
	}

	return false

}
