package _14

//https://leetcode-cn.com/problems/longest-common-prefix/

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func commonPrefix(a string, b string) string {
	if len(a) == 0 || len(b) == 0 {
		return ""
	}

	for i := 0; i < min(len(a), len(b)); i++ {
		if a[i] != b[i] {
			return a[0:i]
		}
	}
	if len(a) > len(b) {
		return b
	} else {
		return a
	}
}

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := range strs {
		current := strs[i]
		println(current, prefix)
		prefix = commonPrefix(current, prefix)
		println(prefix)
	}

	return prefix
}
