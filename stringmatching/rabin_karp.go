package stringmatching

// RK 字符串匹配 BF 的升级版，不再逐一比较子串和模式串的字符，而是比较子串和模式串的哈希值。
// 时间复杂度 O(N)

var (
	// 26个小写字母和素数的映射
	key = map[byte]int{
		'a': 1,
		'b': 3,
		'c': 5,
		'd': 7,
		'e': 11,
		'f': 13,
		'g': 17,
		'h': 19,
		'i': 23,
		'j': 27,
		'k': 29,
		'l': 31,
		'm': 37,
		'n': 41,
		'o': 43,
		'p': 47,
		'q': 51,
		'r': 53,
		's': 57,
		't': 59,
		'u': 61,
		'v': 65,
		'w': 67,
		'x': 71,
		'y': 73,
		'z': 79,
	}
)

func RKMatch(s, m string) (bool, int) {
	sLen := len(s)
	mLen := len(m)
	if sLen < mLen {
		return false, -1
	}
	mHashValue := hash(m)
	for i := 0; i < sLen-mLen+1; i++ {
		subS := s[i : i+mLen]
		subHashValue := hash(subS)
		// 如果哈希值相同，判断
		if subHashValue == mHashValue {
			return true, i
		}
	}
	return false, -1
}

// hash 函数, 保证不会有哈希冲突
func hash(s string) int {
	hashValue := 0
	for i := len(s) - 1; i >= 0; i-- {
		power := 1
		for j := 0; j < len(s)-1-i; j++ {
			power *= 26
		}
		hashValue += key[s[i]] * power
	}
	return hashValue
}
