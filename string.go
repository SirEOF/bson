package bson

import (
	"bytes"
	"encoding/binary"
)

type CString string
type String string

// cstring ::= (byte*) "\x00"
// Zero or more modified UTF-8 encoded characters followed by '\x00'.
// The (byte*) MUST NOT contain '\x00', hence it is not full UTF-8.
func (s CString) ToBSON() []byte {
	return append([]byte(s), byte(0))
}

func (s CString) ToString() string {
	return "\"" + string(s) + "\""
}

// string  ::= int32 (byte*) "\x00"
// The int32 is the number bytes in the (byte*) + 1 (for the trailing '\x00').
// The (byte*) is zero or more UTF-8 encoded characters.
func (s String) ToBSON() []byte {
	buf := new(bytes.Buffer)
	length := len(s) + 1

	err := binary.Write(buf, binary.LittleEndian, uint32(length))
	if err != nil {
		panic(err)
	}

	_, err = buf.Write([]byte(s))
	if err != nil {
		panic(err)
	}

	err = buf.WriteByte(0)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func (s String) ToString() string {
	return "\"" + string(s) + "\""
}
