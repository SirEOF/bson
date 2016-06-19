package bson

import "fmt"

type ObjectId []byte

func (id ObjectId) ToBSON() []byte {
	return id
}

func (id ObjectId) ToString() string {
	return fmt.Sprintf("ObjectId{%h}", id)
}
