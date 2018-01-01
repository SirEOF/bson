package bson

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Double float64

func (d *Double) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, d)
	checkError(err)
	return buf.Bytes(), nil
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
	checkError(err)
	return buf.Bytes(), nil
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
	checkError(err)
	return buf.Bytes(), nil
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
