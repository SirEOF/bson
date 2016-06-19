package bson

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Double float64

func (d Double) ToBSON() []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, d)
	checkError(err)
	return buf.Bytes()
}

func (d Double) ToString() string {
	return fmt.Sprintf("double{%f}", d)
}

type Int32 int32

func (i Int32) ToBSON() []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, i)
	checkError(err)
	return buf.Bytes()
}

func (i Int32) ToString() string {
	return fmt.Sprintf("int32{%d}", i)
}

type Int64 int64

func (i Int64) ToBSON() []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, i)
	checkError(err)
	return buf.Bytes()
}

func (i Int64) ToString() string {
	return fmt.Sprintf("int64{%d}", i)
}

type Null int

func (n Null) ToBSON() []byte {
	return []byte{}
}

func (n Null) ToString() string {
	return fmt.Sprintf("null{}")
}

type Byte byte

func (b Byte) ToBSON() []byte {
	return []byte{byte(b)}
}

func (b Byte) ToString() string {
	return fmt.Sprintf("byte{%d}", b)
}

type RegExp struct {
	RegExp  CString
	Options CString
}

func (r RegExp) ToBSON() []byte {
	return append(CString(r.RegExp).ToBSON(), CString(r.Options).ToBSON()...)
}

func (r RegExp) ToString() string {
	return fmt.Sprintf("regexp{%s, %s}", r.RegExp.ToString(), r.Options.ToString())
}

type DBPointer struct {
	Namespace String
	ObjectId
}

func (p DBPointer) ToBSON() []byte {
	return append(String(p.Namespace).ToBSON(), p.ObjectId.ToBSON()...)
}
