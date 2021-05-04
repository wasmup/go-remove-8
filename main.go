package main

import (
	"bytes"
	"fmt"
)

func remove8(buf []byte) []byte {
	const somechar = 8
	pos := bytes.IndexByte(buf, somechar) // find first somechar postion
	if pos < 0 {
		return buf
	}
	dst := pos
	var src int
	for {
		src = pos + 1 // next non somechar
		for ; src < len(buf) && buf[src] == somechar; src++ {
		}
		dst -= src - pos // number of somechar
		if dst < 0 {
			dst = 0 // home
		}
		// copy non somechar
		for ; src < len(buf) && buf[src] != somechar; src++ {
			buf[dst] = buf[src]
			dst++
		}
		if src >= len(buf) {
			return buf[:dst]
		}
		pos = src // somechar postion
	}
}

func main() {
	stdOutputResult := []byte{'1', 8, 8, 8, 8, 8, '2', 8, 8, '3', '4', 8}
	fmt.Println(stdOutputResult)
	txt := remove8(stdOutputResult)
	fmt.Println(txt)
	fmt.Println(string(txt))
}
