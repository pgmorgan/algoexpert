import "strings"

func toGoatLatin(S string) string {
	words := strings.Fields(S)
	for i := range words {
		vowelStart := false
		for _, vowel := range []byte{'a', 'e', 'i', 'o', 'u'} {
			if words[i][0] == vowel {
				tmp := []byte(words[i])
				tmp = append(tmp, 'm', 'a')
				words[i] = string(tmp)
				vowelStart = true
			}
		}
		if vowelStart {
			continue
		}
		words[i] = transformRegular(words[i], i)
	}
	return strings.Join(words, " ")
}

func transformRegular(str string, i int) string {
	word := []byte(str)
	tmp := word[0]
	for i := 1; i < len(word); i++ {
		word[i-1] = word[i]
	}
	word[len(word)-1] = tmp
	word = append(word, 'm', 'a')
	for j := i; j >= 0; j-- {
		word = append(word, 'a')
	}
	return string(word)
}