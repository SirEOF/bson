package bson

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

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

func (b Binary) Serialize() ([]byte, error) {

	length, err := Int32(len(b.Data) + 1).Serialize()
	if err != nil {
		return nil, err
	}

	out := append(length, b.Subtype)
	out = append(out, b.Data...)
	return out, nil
}

func (b *Binary) Deserialize(in *bytes.Reader) error {

	length := int(0)
	err := binary.Read(in, binary.LittleEndian, &length)

	if err != nil {
		return err
	}

	b.Data = make([]byte, length)
	n, err := in.Read(b.Data)

	if err != nil {
		return err
	}

	if n != length {
		return errors.New("invalid string length")
	}

	return nil
}

func (b Binary) String() string {
	return fmt.Sprintf("binary{%s, %h}", BinaryTypeMap[b.Subtype], b.Data)
}
