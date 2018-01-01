package bson

import (
	"bytes"
	"errors"
	"fmt"
)

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
	return fmt.Sprintf("ObjectId{%h}", id)
}
