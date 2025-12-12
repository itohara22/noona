package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// dic := "d3:cow3:moo4:spam4:eggsi42el4:test3:one3:twoe3:fooi99ee"
	// decodeDict(dic)
	strrr := "4:spam3:egg"
	for _, s := range decodeString(strrr) {
		fmt.Println(s)
	}
}

func decodeString(s string) []string { // Implemnet chr counter to go over the whole bencode *i
	res := []string{}
	i := 0
	for i < len(s) {
		r := rune(s[i])
		if r > '0' && r < '9' {
			length, err := strconv.Atoi(string(r))
			if err != nil {
				log.Fatal(err.Error())
			}
			i++
			if s[i] != ':' {
				log.Fatal("expected colon")
			}
			i++
			strParsed := s[i : i+length]
			res = append(res, strParsed)
		} else {
			i++
		}
	}
	return res
}

// func decodeDict(s string) map[string]string {
// 	m := make(map[string]string)
// 	return m
// }
