package reverseprefixofword

import "fmt"

func reversePrefix(word string, ch byte) string {
	q := []byte{}
	for i := 0; i < len(word); i++ {
		q = append([]byte{word[i]}, q...)
		if word[i] == ch {
			fmt.Println(q, word, i)
			word = string(q[:]) + word[i+1:]
			break
		}
	}
	return word
}

func ReversePrefix(word string, ch byte) string {
	return reversePrefix(word, ch)
}
