package main

import (
	"fmt"
	"unsafe"
)

func main() {
	b := []byte{1, 2, 3, 4, 5}
	pb := &b[3]
	fmt.Println(unsafe.Slice(pb, 2))

	pb = (*byte)(unsafe.Add(unsafe.Pointer(pb), -1))
	fmt.Println(*pb)
}
