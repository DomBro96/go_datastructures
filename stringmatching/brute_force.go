package stringmatching

// 字符串匹配BF算法: 蛮力算法, 简单粗暴好理解, 最常用
// 模式串 从头匹配 主串 时间复杂度 O(N*M)
func BruteForceMatch(s, m string) (bool, int) {
	sLen := len(s)
	mLen := len(m)
	// 主串长度小于模式串长度 false
	if sLen < mLen {
		return false, -1
	}
	// 0 -> (mLen - sLen + 1)
	for i := 0; i < sLen-mLen+1; i++ {
		subS := s[i : i+mLen] // 子串
		isMatch := true
		for j := 0; j < mLen; j++ {
			if subS[j] != m[j] {
				isMatch = false
				break
			}
		}
		if isMatch {
			return true, i
		}
	}
	return false, -1
}
