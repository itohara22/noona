package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode"
)

// dic := "d3:cow3:moo4:spam4:eggsi42el4:test3:one3:twoe3:fooi99ee"
func main() {
	strrr := "4:spam3:eggi42e"
	i := 0
	var res []any
	for i < len(strrr) {
		r := rune(strrr[i])
		if unicode.IsDigit(r) {
			res = append(res, decodeString(strrr, &i))
		} else if r == 'i' {
			res = append(res, decodeInteger(strrr, &i))
		} else {
			i++
		}
	}
	fmt.Println(res)
}

func decodeInteger(s string, i *int) int {
	res := ""
	*i++ // skip i
	for string(s[*i]) != "e" {
		res += string(s[*i])
		*i++
	}
	resInt, err := strconv.Atoi(res)
	fmt.Println(res)
	if err != nil {
		log.Fatal("invalid bencode")
	}
	return resInt
}

func decodeString(s string, i *int) string {
	res := ""
	r := rune(s[*i])
	if r > '0' && r < '9' {
		length, err := strconv.Atoi(string(r))
		if err != nil {
			log.Fatal(err.Error())
		}
		*i++ //move to colon
		if s[*i] != ':' {
			log.Fatal("expected colon")
		}
		*i++ // skip colon
		strParsed := s[*i : *i+length]
		res += strParsed
	} else {
		*i++
	}
	return res
}

// func decodeDict(s string) map[string]string {
// 	m := make(map[string]string)
// 	return m
// }
