package bson

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
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

type Double float64

func (d *Double) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, d)
	return buf.Bytes(), err
}

func (d *Double) Deserialize(in *bytes.Reader) error {
	err := binary.Read(in, binary.LittleEndian, d)

	return err
}

func (d Double) String() string {
	return fmt.Sprintf("double{%f}", d)
}

type Int32 int32

func (i Int32) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, i)

	return buf.Bytes(), err
}

func (i *Int32) Deserialize(in *bytes.Reader) error {
	err := binary.Read(in, binary.LittleEndian, i)

	return err
}

func (i Int32) String() string {
	return fmt.Sprintf("int32{%d}", i)
}

type Int64 int64

func (i Int64) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, i)
	return buf.Bytes(), err
}

func (i *Int64) Deserialize(in *bytes.Reader) error {
	err := binary.Read(in, binary.LittleEndian, i)

	return err
}

func (i Int64) String() string {
	return fmt.Sprintf("int64{%d}", i)
}

type Null int

func (n Null) Serialize() ([]byte, error) {
	return []byte{}, nil
}

func (n *Null) Deserialize(in *bytes.Reader) error {
	return nil
}

func (n Null) String() string {
	return fmt.Sprintf("null{}")
}

type Byte byte

func (b Byte) Serialize() ([]byte, error) {
	return []byte{byte(b)}, nil
}

func (b *Byte) Deserialize(in *bytes.Reader) error {
	out, err := in.ReadByte()
	*b = Byte(out)
	return err
}

func (b Byte) String() string {
	return fmt.Sprintf("byte{%d}", b)
}

type RegExp struct {
	RegExp  CString
	Options CString
}

func (r RegExp) Serialize() ([]byte, error) {
	reg, err := r.RegExp.Serialize()
	if err != nil {
		return nil, err
	}

	options, err := r.Options.Serialize()
	if err != nil {
		return nil, err
	}

	return append(reg, options...), nil
}

func (r *RegExp) Deserialize(in *bytes.Reader) error {
	err := r.RegExp.Deserialize(in)
	if err != nil {
		return err
	}

	err = r.Options.Deserialize(in)
	if err != nil {
		return err
	}

	return nil
}

func (r RegExp) String() string {
	return fmt.Sprintf("regexp{%s, %s}", r.RegExp.String(), r.Options.String())
}

type DBPointer struct {
	Namespace String
	ObjectId
}

func (p DBPointer) Serialize() ([]byte, error) {
	namespace, err := p.Namespace.Serialize()
	if err != nil {
		return nil, err
	}

	objectid, err := p.ObjectId.Serialize()
	if err != nil {
		return nil, err
	}
	return append(namespace, objectid...), nil
}

func (d *DBPointer) Deserialize(in *bytes.Reader) error {
	err := d.Namespace.Deserialize(in)
	if err != nil {
		return err
	}

	err = d.ObjectId.Deserialize(in)
	if err != nil {
		return err
	}

	return nil
}

type MinKey byte
type MaxKey byte

func (m MinKey) Serialize() ([]byte, error) {
	return []byte{}, nil
}

func (m MinKey) Deserialize(in *bytes.Reader) error {
	return nil
}

func (m MinKey) String() string {
	return "MinKey{}"
}

func (m MaxKey) Serialize() ([]byte, error) {
	return []byte{}, nil
}

func (m MaxKey) Deserialize(in *bytes.Reader) error {
	return nil
}

func (m MaxKey) String() string {
	return "MaxKey{}"
}

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

type ObjectId []byte

func (id ObjectId) Serialize() ([]byte, error) {
	return id, nil
}

func (id *ObjectId) Deserialize(in *bytes.Reader) error {
	data := make([]byte, 12)
	n, err := in.Read(data)
	if err != nil {
		return err
	}

	if n != 12 {
		return errors.New("invalid objectid length")
	}

	*id = data
	return nil
}

func (id ObjectId) String() string {
	return fmt.Sprintf("ObjectId{%x}", []byte(id))
}
