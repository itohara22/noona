package main

import (
	"log"
	"strconv"
	"unicode"
)

type BencodeDecoder struct{}

func newBencodeDecoder() *BencodeDecoder {
	return &BencodeDecoder{}
}

func (bencodeDecoder BencodeDecoder) Decode(str string) any {
	i := 0
	res := parseValue(str, &i)
	if i != len(str) {
		log.Fatal("trailing data after valid bencode")
	}
	return res
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
