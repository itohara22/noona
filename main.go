package main

import (
	"fmt"
)

func main() {
	dic := "d3:cow3:moo4:spam4:eggs3:numi42e4:listl4:test3:one3:twoe3:fooi99ee"
	bencodeDecoder := newBencodeDecoder()
	res := bencodeDecoder.Decode(dic)
	fmt.Println(res)
}
