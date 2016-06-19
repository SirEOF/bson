package bson

import (
	"bytes"
	"encoding/binary"
)

//http://bsonspec.org/spec.html

type BSON interface {
	ToBSON() []byte
	ToString() string
}

type Document ElementList
type ElementList []Element

// e_list  ::= element e_list
// | ""
func (elist ElementList) ToBSON() []byte {
	out := []byte{}
	for _, e := range elist {
		out = append(out, e.ToBSON()...)
	}
	return out
}

func (elist ElementList) ToString() string {
	out := ""
	for _, e := range elist {
		out += e.ToString() + ", "
	}

	return out[0 : len(out)-2]
}

// document  ::= int32 e_list "\x00"
// int32 is the total number of bytes comprising the document.
func (d Document) ToBSON() []byte {
	buf := new(bytes.Buffer)

	eListBSON := ElementList(d).ToBSON()
	binary.Write(buf, binary.LittleEndian, int32(len(eListBSON)+5))
	buf.Write(eListBSON)
	buf.WriteByte(0)

	return buf.Bytes()
}

func (d Document) ToString() string {
	return "{" + ElementList(d).ToString() + "}"
}
