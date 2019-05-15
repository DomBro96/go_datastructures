package greedy

// 给定长度为N字符串S，求长度为N的字符串T，期初T是一个空串，每次删除并复制S的开头或结尾较小的字符，目的是构造
// 字典序较小的T。
func BestCowLine(s string) string {
	// 反转后的字符串
	s1 := ReverseString(s)
	sIndex, s1Index, tIndex := 0, 0, len(s1) - 1
	// ACDBCB
	// BCBDCA
	res := make([]byte, len(s))
	for i := 0; i < len(res); i++ {
		if s[sIndex] < s1[s1Index] {
			res[i] =  s[sIndex]
			sIndex++
		}else {
			res[i] = s[tIndex]
			tIndex--
			s1Index++
		}
	}
	return string(res)
}

func ReverseString(s string) string {
	chars := []byte(s)
	for i, j := 0, len(chars)-1; i < j; {
		chars[i], chars[j] = chars[j], chars[i]
		i++
		j--
	}
	return string(chars)
}
