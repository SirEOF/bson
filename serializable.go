package bson

import "bytes"

type Serializable interface {
	Serialize() ([]byte, error)
	Deserialize(*bytes.Reader) error
}

func SerializeArray(arr []Serializable) ([]byte, error) {
	out := []byte{}
	for _, a := range arr {
		d, err := a.Serialize()
		if err != nil {
			return nil, err
		}

		out = append(out, d...)
	}
	return out, nil
}

func DeserializeArray(arr []Serializable, in *bytes.Reader) error {
	for _, a := range arr {
		err := a.Deserialize(in)
		if err != nil {
			return err
		}
	}
	return nil
}
