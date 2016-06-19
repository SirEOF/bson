package bson

import (
	"bytes"
	"fmt"
	"testing"
)

func TestDocumentToBSON(t *testing.T) {
	out := []byte("\x16\x00\x00\x00\x02hello\x00\x06\x00\x00\x00world\x00\x00")
	document := Document([]Element{NewStringElement("hello", "world")})

	if bytes.Compare(document.ToBSON(), out) != 0 {
		t.Error("Failed to generate bson for hello world document\n", out, "\n", document.ToBSON())
	}

	out2 := []byte("\x31\x00\x00\x00\x04BSON\x00\x26\x00\x00\x00\x02\x30\x00\x08\x00\x00\x00awesome\x00\x01\x31\x00\x33\x33\x33\x33\x33\x33\x14\x40\x10\x32\x00\xc2\x07\x00\x00\x00\x00")
	a0 := NewStringElement("0", "awesome")
	a1 := NewDoubleElement("1", 5.05)
	a2 := NewInt32Element("2", 1986)
	a := []Element{a0, a1, a2}

	document2 := Document([]Element{NewArrayElement("BSON", a)})
	fmt.Println(document2.ToString())
	if bytes.Compare(document2.ToBSON(), out2) != 0 {
		t.Error("Failed to generate bson for array document\n", out2, "\n", document2.ToBSON())
	}
}

func TestBSONToString(t *testing.T) {
	document := Document([]Element{NewStringElement("hello", "world")})
	fmt.Println(document.ToString())
}

func TestStringElement(t *testing.T) {
	a := NewStringElement("hello", "world")
	out := []byte("\x02hello\x00\x06\x00\x00\x00world\x00")

	if bytes.Compare(a.ToBSON(), out) != 0 {
		t.Error("Failed to generate bson for hello world StringElement\n", out, "\n", a.ToBSON())
	}
}

func TestDoubleElement(t *testing.T) {
	a := NewDoubleElement("swag", 5.05)
	out := []byte("\x01swag\x00\x33\x33\x33\x33\x33\x33\x14\x40")
	if bytes.Compare(a.ToBSON(), out) != 0 {
		t.Error("Failed to generate bson for example double element\n", a.ToBSON(), "\n", out)
	}
}
