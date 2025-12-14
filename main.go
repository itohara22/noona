package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode"
)

func main() {
	dic := "d3:cow3:moo4:spam4:eggs3:numi42e4:listl4:test3:one3:twoe3:fooi99ee"
	i := 0
	res := parseValue(dic, &i)
	fmt.Println(res)
}

func parseValue(s string, i *int) any {
	r := rune(s[*i])
	if unicode.IsDigit(r) {
		return decodeString(s, i)
	} else if r == 'i' {
		return decodeInteger(s, i)
	} else if r == 'l' {
		return decodeList(s, i)
	} else if r == 'd' {
		return decodeDict(s, i)
	} else {
		log.Fatal("parse value invalid bencode")
	}
	return nil
}

func decodeList(s string, i *int) []any {
	*i++
	res := []any{}
	for s[*i] != 'e' {
		res = append(res, parseValue(s, i))
	}
	*i++ // skip the last e
	return res
}

func decodeInteger(s string, i *int) int {
	res := ""
	*i++ // skip i
	for s[*i] != 'e' {
		res += string(s[*i])
		*i++
	}
	*i++ // move from last e
	resInt, err := strconv.Atoi(res)
	if err != nil {
		log.Fatal("decode integer invalid bencode")
	}
	return resInt
}

func decodeString(s string, i *int) string {
	res := ""
	lenghtString := ""

	for s[*i] >= '0' && s[*i] <= '9' {
		lenghtString += string(s[*i])
		*i++
	}

	if s[*i] != ':' {
		log.Fatal("expected colon")
	}
	*i++ // skip colon

	length, err := strconv.Atoi(lenghtString)
	if err != nil {
		log.Fatal(err.Error())
	}
	strParsed := s[*i : *i+length]
	res += strParsed
	*i += length // move i to the next digit
	return res
}

func decodeDict(s string, i *int) map[string]any {
	*i++ // skip d
	m := make(map[string]any)

	for s[*i] != 'e' {
		key := decodeString(s, i)
		val := parseValue(s, i)
		m[key] = val
	}

	*i++ // skip the last e
	return m
}
