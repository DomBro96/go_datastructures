package string

// 用 %20 替换字符串里的空格

func ReplaceBlank(s string) string {
	originalLen := len(s)
	blankNum := 0
	for _, c := range s {
		if c == ' ' {
			blankNum++
		}
	}

	// 新字符串长度 = 原字符串长度 + 2*空格长度
	newStr := make([]byte, originalLen+2*blankNum, originalLen+2*blankNum)
	for i := range s {
		newStr[i] = s[i]
	}
	// 原字符串下标
	oIndex := originalLen - 1
	// 新字符串下标
	nIndex := len(newStr) - 1
	for oIndex >= 0 && nIndex > oIndex {
		if newStr[oIndex] == ' ' {
			newStr[nIndex] = '0'
			nIndex--
			newStr[nIndex] = '2'
			nIndex--
			newStr[nIndex] = '%'
			nIndex--
		} else {
			newStr[nIndex] = newStr[oIndex]
			nIndex--
		}
		oIndex--
	}
	return string(newStr)
}
