package _387

//https://leetcode-cn.com/problems/first-unique-character-in-a-string/

type CountAndIndex struct {
	count int
	index int
}

func firstUniqChar(s string) int {

	charMap := make(map[uint8]*CountAndIndex)

	for i := range s {
		char := s[i]
		countAndIndex, contains := charMap[char]
		if contains {
			countAndIndex.count++
		} else {
			charMap[char] = &CountAndIndex{1, i}
		}
	}
	result := -1
	for u := range charMap {
		if charMap[u].count == 1 {
			if result == -1 {
				result = charMap[u].index
			} else if charMap[u].index < result {
				result = charMap[u].index
			}
		}
	}

	return result
}
