package bson

import (
	"bytes"
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
