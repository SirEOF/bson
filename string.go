package bson

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type CString string
type String string

// cstring ::= (byte*) "\x00"
// Zero or more modified UTF-8 encoded characters followed by '\x00'.
// The (byte*) MUST NOT contain '\x00', hence it is not full UTF-8.
func (s CString) Serialize() ([]byte, error) {
	return append([]byte(s), byte(0)), nil
}

func (s *CString) Deserialize(in *bytes.Reader) error {
	*s = ""

	out := ""
	for {
		c, err := in.ReadByte()
		if err != nil {
			return err
		}

		if c == byte(0) {
			break
		}

		out += string(c)
	}

	*s = CString(out)
	return nil
}

func (s CString) String() string {
	return "\"" + string(s) + "\""
}

// string  ::= int32 (byte*) "\x00"
// The int32 is the number bytes in the (byte*) + 1 (for the trailing '\x00').
// The (byte*) is zero or more UTF-8 encoded characters.
func (s String) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	length := len(s) + 1

	err := binary.Write(buf, binary.LittleEndian, uint32(length))
	if err != nil {
		return nil, err
	}

	_, err = buf.Write([]byte(s))
	if err != nil {
		return nil, err
	}

	err = buf.WriteByte(0)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (s *String) Deserialize(in *bytes.Reader) error {

	length := uint32(0)
	err := binary.Read(in, binary.LittleEndian, &length)

	if err != nil {
		return err
	}

	out := make([]byte, length)
	n, err := in.Read(out)

	if err != nil {
		return err
	}

	if uint32(n) != length {
		return errors.New("invalid string length")
	}

	*s = String(out[0 : len(out)-1])
	return nil
}

func (s String) String() string {
	return "\"" + string(s) + "\""
}
