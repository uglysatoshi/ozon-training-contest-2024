package main

import (
	"fmt"
	"unicode"
)

func main3() {
	var N int
	_, err := fmt.Scan(&N)
	if err != nil {
		return
	}
	for i := 0; i < N; i++ {
		var str string
		var answers []string
		_, err2 := fmt.Scan(&str)
		if err2 != nil {
			return
		}
		for len(str) > 4 {
			if unicode.IsLetter(rune(str[0])) && unicode.IsDigit(rune(str[1])) && unicode.IsDigit(rune(str[2])) && unicode.IsLetter(rune(str[3])) && unicode.IsLetter(rune(str[4])) {
				answers = append(answers, str[0:5])
				str = str[5:]
				continue
			}

			if unicode.IsLetter(rune(str[0])) && unicode.IsDigit(rune(str[1])) && unicode.IsLetter(rune(str[2])) && unicode.IsLetter(rune(str[3])) {
				answers = append(answers, str[0:4])
				str = str[4:]
				continue
			}
			break
		}
		if len(str) > 0 {
			if len(str) >= 4 {
				if !(unicode.IsLetter(rune(str[0])) && unicode.IsDigit(rune(str[1])) && unicode.IsLetter(rune(str[2])) && unicode.IsLetter(rune(str[3]))) {
					fmt.Println("-")
					answers = nil
				} else {
					answers = append(answers, str[0:4])
				}
			} else {
				fmt.Println("-")
				answers = nil
			}
		}

		if answers != nil {
			for _, value := range answers {
				fmt.Printf("%v ", value)
			}
			fmt.Println()
		}
	}
}
