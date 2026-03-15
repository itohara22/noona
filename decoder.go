package main

import (
	"strings"
	"log"
	"strconv"
	"unicode"
)

type BencodeDecoder struct{}

func newBencodeDecoder() *BencodeDecoder {
	return &BencodeDecoder{}
}

func (bencodeDecoder *BencodeDecoder) Decode(str []byte) any {
	i := 0
	res := parseValue(str, &i)
	if i != len(str) {
		log.Fatal("trailing data after valid bencode")
	}
	return res
}

func parseValue(s []byte, i *int) any {
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


func decodeList(byteStr []byte, pos *int) []any {
	*pos++ // skip l
	res := []any{}

	for byteStr[*pos] != 'e' {
		res = append(res, parseValue(byteStr, pos))
	}
	*pos++ // skip the last e
	return res
}


func decodeInteger(byteStr []byte, pos *int) int {
	var res strings.Builder
	*pos++ // skip i

	for byteStr[*pos] != 'e' {
		res.WriteString(string(byteStr[*pos]))
		*pos++
	}
	*pos++ // move from last e
	resInt, err := strconv.Atoi(res.String())
	if err != nil {
		log.Fatal("decode integer invalid bencode")
	}
	return resInt
}


func decodeString(byteStr []byte, pos *int) []byte {
	// strings can be byte strings so its best to return bytes
	var lenghtString strings.Builder

	for byteStr[*pos] >= '0' && byteStr[*pos] <= '9' { // comapring ascii values
		// we will check for digits as string lenght can be multiple digits
		lenghtString.WriteString(string(byteStr[*pos]))
		*pos++
	}

	if byteStr[*pos] != ':' {
		log.Fatal("expected colon")
	}
	*pos++ // skip colon
	length, err := strconv.Atoi(lenghtString.String())
	if err != nil {
		log.Fatal(err.Error())
	}
	strParsed := byteStr[*pos : *pos+length]
	res := strParsed
	*pos += length // move i to the next digit
	return []byte(res)
}

func decodeDict(byteStr []byte, pos *int) map[string]any {
	*pos++ // skip d
	m := make(map[string]any)

	for byteStr[*pos] != 'e' {
		key := string(decodeString(byteStr, pos))
		val := parseValue(byteStr, pos)
		m[key] = val
	}

	*pos++ // skip the last e
	return m
}
