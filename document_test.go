package bson

import (
	"bytes"
	"fmt"
	"testing"
)

func TestDocuments(t *testing.T) {
	in1 := bytes.NewReader([]byte("\x16\x00\x00\x00\x02hello\x00\x06\x00\x00\x00world\x00\x00"))
	document1 := new(Document)
	err := document1.Deserialize(in1)

	if err != nil {
		t.Error(err)
	}
	SerializableTest(t, document1)

	in2 := bytes.NewReader([]byte("\x31\x00\x00\x00\x04BSON\x00\x26\x00\x00\x00\x02\x30\x00\x08\x00\x00\x00awesome\x00\x01\x31\x00\x33\x33\x33\x33\x33\x33\x14\x40\x10\x32\x00\xc2\x07\x00\x00\x00\x00"))
	document2 := new(Document)
	err = document2.Deserialize(in2)

	if err != nil {
		t.Error(err)
	}

	SerializableTest(t, document2)

}

func TestDocumentRead(t *testing.T) {
	in1 := bytes.NewReader([]byte("\x16\x00\x00\x00\x02hello\x00\x06\x00\x00\x00world\x00\x00"))
	document1 := new(Document)
	err := document1.Deserialize(in1)

	if err != nil {
		t.Error(err)
	}

	if document1.Get(0).EName != CString("hello") {
		t.Error("key value should of been hello!")
	}

	if document1.Get(0).Data.String() != String("world").String() {
		t.Error("value should of been world!")
	}
}

func TestDocumentEdit(t *testing.T) {
	in1 := bytes.NewReader([]byte("\x16\x00\x00\x00\x02hello\x00\x06\x00\x00\x00world\x00\x00"))
	document1 := new(Document)
	err := document1.Deserialize(in1)

	if err != nil {
		t.Error(err)
	}

	if document1.Key("hello") == nil {
		t.Error("Failed to find key")
	}

	replace := Int32(1337)
	document1.Key("hello").Data = &replace

	if document1.Get(0).Data.String() != replace.String() {
		t.Error("value should of been world!")
	}
}
func TestDocumentManual(t *testing.T) {

	// {"BSON": {"0": "awesome", "1": double{5.050000}, "2": int32{1986}}}

	d := Document{
		NewStringElement("hello", "world"),
	}
	hello1 := []byte("\x16\x00\x00\x00\x02hello\x00\x06\x00\x00\x00world\x00\x00")

	hello1Guess, err := d.Serialize()
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(hello1, hello1Guess) {
		fmt.Println(hello1)
		fmt.Println(hello1Guess)
		t.Error("manually created faulty document")
	}

	// Try 1

	arr := []Element{
		NewStringElement("0", "awesome"),
		NewDoubleElement("1", 5.05000),
		NewInt32Element("2", 1986),
	}

	doc := Document{
		NewArrayElement("BSON", arr),
	}

	correct := []byte("\x31\x00\x00\x00\x04BSON\x00\x26\x00\x00\x00\x02\x30\x00\x08\x00\x00\x00awesome\x00\x01\x31\x00\x33\x33\x33\x33\x33\x33\x14\x40\x10\x32\x00\xc2\x07\x00\x00\x00\x00")
	guess, err := doc.Serialize()
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(correct, guess) {
		fmt.Println(correct)
		fmt.Println(guess)
		t.Error("manually created faulty document")
	}

	SerializableTest(t, &doc)
}
