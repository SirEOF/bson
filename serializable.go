package bson

import "bytes"

type Serializable interface {
	Serialize() ([]byte, error)
	Deserialize(*bytes.Reader) error
}
