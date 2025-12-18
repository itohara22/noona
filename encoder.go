package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
)

type BencodeEncoder struct {
}

func NewBencodeEncoder() *BencodeEncoder {
	return &BencodeEncoder{}
}

func (encoder *BencodeEncoder) Encode(data any) []byte {
	res := encodeAny(data)
	return res
}

func encodeAny(data any) []byte {

	result := []byte{}
	switch dataType := data.(type) { // switch case according to data type
	case string:
		result = append(result, encodeString(data)...)
	case []byte:
		result = append(result, encodeBytes(data)...)
	case map[string]any:
		result = append(result, encodeDict(data)...)
	case int:
		result = append(result, encodeInt(data)...)
	case []any:
		result = append(result, encodeList(data)...)
	default:
		fmt.Printf("%v", dataType)
		log.Fatal("went wrong")
	}
	return result
}

func encodeBytes(data any) []byte {
	dataBytes, ok := data.([]byte)
	if !ok {
		log.Fatal("encode byte broke")
	}
	length := strconv.Itoa(len(dataBytes))
	res := []byte(length + ":")
	res = append(res, dataBytes...)
	return res
}

func encodeDict(data any) []byte {
	dataMap, ok := data.(map[string]any)
	if !ok {
		log.Fatal("shittt")
	}
	keys := make([]string, 0, len(dataMap))
	for k := range dataMap {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	res := []byte{'d'}
	for _, k := range keys {
		res = append(res, encodeString(k)...)       // we use ... to spread content
		res = append(res, encodeAny(dataMap[k])...) // encode map values
	}
	res = append(res, 'e')
	return res
}

func encodeInt(data any) []byte {
	dataInt, ok := data.(int)
	if !ok {
		log.Fatal("decode int broke")
	}
	res := "i" + strconv.Itoa(dataInt) + "e"
	return []byte(res)
}

func encodeString(data any) []byte {
	dataStr, ok := data.(string)
	if !ok {
		log.Fatal("encode string broke")
	}
	res := strconv.Itoa(len(dataStr)) + ":" + dataStr
	return []byte(res)
}

func encodeList(data any) []byte {
	res := []byte{'l'}
	dataList, ok := data.([]any)
	if !ok {
		log.Fatal("encode list broke")
	}
	for _, v := range dataList {
		res = append(res, encodeAny(v)...)
	}
	res = append(res, 'e')
	return res
}
