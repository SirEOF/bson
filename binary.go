package bson

import "fmt"

const (
	BINARY_GENERIC byte = iota
	BINARY_FUNCTION
	BINARY_BINARY_DEPRECIATED
	BINARY_UUID_DEPRECIATED
	BINARY_UUID
	BINARY_MD5
)

var BinaryTypeMap = map[byte]string{
	0: "GENERIC",
	1: "FUNCTION",
	2: "BINARY (OLD)",
	3: "UUID (OLD)",
	4: "UUID",
	5: "MD5",
}

type Binary struct {
	Subtype byte
	Data    []byte
}

func (b Binary) ToBSON() []byte {
	length := len(b.Data) + 1
	out := append(Int32(length).ToBSON(), b.Data...)
	return append(out, byte(0))
}

func (b Binary) ToString() string {
	return fmt.Sprintf("binary{%s, %h}", BinaryTypeMap[b.Subtype], b.Data)
}
