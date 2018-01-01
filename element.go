package bson

import (
	"bytes"
	"errors"
	"fmt"
)

// element	::=	"\x01" e_name double	64-bit binary floating point
// |	"\x02" e_name string	UTF-8 string
// |	"\x03" e_name document	Embedded document
// |	"\x04" e_name document	Array
// |	"\x05" e_name binary	Binary data
// |	"\x06" e_name	Undefined (value) — Deprecated
// |	"\x07" e_name (byte*12)	ObjectId
// |	"\x08" e_name "\x00"	Boolean "false"
// |	"\x08" e_name "\x01"	Boolean "true"
// |	"\x09" e_name int64	UTC datetime
// |	"\x0A" e_name	Null value
// |	"\x0B" e_name cstring cstring	Regular expression - The first cstring is the regex pattern, the second is the regex options string. Options are identified by characters, which must be stored in alphabetical order. Valid options are 'i' for case insensitive matching, 'm' for multiline matching, 'x' for verbose mode, 'l' to make \w, \W, etc. locale dependent, 's' for dotall mode ('.' matches everything), and 'u' to make \w, \W, etc. match unicode.
// |	"\x0C" e_name string (byte*12)	DBPointer — Deprecated
// |	"\x0D" e_name string	JavaScript code
// |	"\x0E" e_name string	Deprecated
// |	"\x0F" e_name code_w_s	JavaScript code w/ scope
// |	"\x10" e_name int32	32-bit integer
// |	"\x11" e_name int64	Timestamp
// |	"\x12" e_name int64	64-bit integer
// |	"\xFF" e_name	Min key
// |	"\x7F" e_name	Max key

type BSONElement interface {
	Serializable
	String() string
}

type Element struct {
	Identifier byte
	EName      CString
	Data       BSONElement
}

func (e *Element) Deserialize(in *bytes.Reader) error {
	var err error
	e.Identifier, err = in.ReadByte()
	if err != nil {
		return err
	}

	err = e.EName.Deserialize(in)

	switch e.Identifier {
	case 0x01: // Double
		e.Data = new(Double)
		err := e.Data.Deserialize(in)
		if err != nil {
			return err
		}
	case 0x02, 0x0D: // String, COde
		e.Data = new(String)
		err := e.Data.Deserialize(in)
		if err != nil {
			return err
		}
	case 0x03, 0x04: // Document embeeded / array
		e.Data = new(Document)
		err := e.Data.Deserialize(in)
		if err != nil {
			return err
		}
	case 0x05: //binary
		e.Data = new(Binary)
		err := e.Data.Deserialize(in)
		if err != nil {
			return err
		}
	// case 0x06: //undefined
	// 	e.Data = new(String)
	// 	err := e.Data.Deserialize(in)
	// 	if err != nil {
	// 		return err
	// 	}
	case 0x07: //obejctid
		e.Data = new(ObjectId)
		err := e.Data.Deserialize(in)
		if err != nil {
			return err
		}
	case 0x08: //boolean
		e.Data = new(Byte)
		err := e.Data.Deserialize(in)
		if err != nil {
			return err
		}
	case 0x09, 0x11, 0x12: // UTC Datetime, 64bit uint (i know i know, i'm lazy), 64bit
		e.Data = new(Int64)
		err := e.Data.Deserialize(in)
		if err != nil {
			return err
		}
	case 0x10: // int32
		e.Data = new(Int32)
		err := e.Data.Deserialize(in)
		if err != nil {
			return err
		}
	case 0x0A: //Null
		e.Data = new(Null)
		err := e.Data.Deserialize(in)
		if err != nil {
			return err
		}
	case 0x0B: // Regex
		e.Data = new(RegExp)
		err := e.Data.Deserialize(in)
		if err != nil {
			return err
		}

	// case 0x0F: // code_w_s
	// 	e.Data = new()
	// 	err := e.Data.Deserialize(in)
	// 	if err != nil {
	// 		return err
	// 	}
	// case 0x13: // String
	// 	e.Data = new(Int32)
	// 	err := e.Data.Deserialize(in)
	// 	if err != nil {
	// 		return err
	// 	}
	case 0xFF: // MinKey
		e.Data = new(MinKey)
		err := e.Data.Deserialize(in)
		if err != nil {
			return err
		}

	case 0x7F: // MaxKey
		e.Data = new(MaxKey)
		err := e.Data.Deserialize(in)
		if err != nil {
			return err
		}
	default:
		return errors.New(fmt.Sprintf("unknown type %x", e.Identifier))
	}

	return nil
}

func (e Element) Serialize() ([]byte, error) {

	ename, err := e.EName.Serialize()
	if err != nil {
		return nil, err
	}

	data, err := e.Data.Serialize()
	if err != nil {
		return nil, err
	}

	out := []byte{}
	out = append(out, e.Identifier)
	out = append(out, ename...)
	out = append(out, data...)
	return out, nil
}

func (e Element) String() string {
	return e.EName.String() + ": " + e.Data.String()
}
