package bson

import (
	"bytes"
	"fmt"
	"testing"
)

func SerializableTest(t *testing.T, e Serializable) {
	data, err := e.Serialize()

	if err != nil {
		t.Error(err)
	}

	err = e.Deserialize(bytes.NewReader(data))
	if err != nil {
		t.Error(err)
	}

	data2, err := e.Serialize()
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(data, data2) {
		fmt.Println(data)
		fmt.Println(data2)
		t.Error("The serialize/deserialize operation was not reversable")
	}
}
