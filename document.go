package bson

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

type Document []Element

func (d Document) Get(i int) *Element {
	return &[]Element(d)[i]
}

func (d Document) Key(name string) *Element {
	for i, e := range []Element(d) {
		if e.EName.String() == fmt.Sprintf("\"%s\"", name) {
			return &[]Element(d)[i]
		}
	}
	return nil
}

// document  ::= int32 e_list "\x00"
// int32 is the total number of bytes comprising the document.
func (d Document) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	out := []byte{}
	for _, e := range d {
		data, err := e.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, data...)
	}

	binary.Write(buf, binary.LittleEndian, int32(len(out)+5))
	buf.Write(out)
	buf.WriteByte(0)

	return buf.Bytes(), nil
}

func (d *Document) Deserialize(in *bytes.Reader) error {

	startOffset := in.Size() - int64(in.Len())

	length := uint32(0)
	err := binary.Read(in, binary.LittleEndian, &length)
	if err != nil {
		return err
	}

	elements := []Element{}
	for (in.Size()-int64(in.Len()))-startOffset < int64(length)-1 { //minus 1 is document null byte
		e := new(Element)
		err := e.Deserialize(in)
		if err != nil {
			return err
		}

		elements = append(elements, *e)
	}

	//chop last byte
	b, err := in.ReadByte()
	if err != nil {
		return err
	}

	if b != 0 {
		return errors.New("last byte of document should be null")
	}

	*d = elements
	return nil
}

func (d Document) String() string {
	if len(d) == 0 {
		return "{}"
	}

	out := "{"
	for _, e := range d {
		out += e.String() + ", "
	}

	return out[0:len(out)-2] + "}"
}
