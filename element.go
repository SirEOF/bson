package bson

import "time"

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

type Element struct {
	Identifier byte
	EName      CString
	Data       BSON
}

func (e Element) ToBSON() []byte {
	out := []byte{}
	out = append(out, e.Identifier)
	out = append(out, e.EName.ToBSON()...)
	out = append(out, e.Data.ToBSON()...)
	return out
}

func (e Element) ToString() string {
	return e.EName.ToString() + ": " + e.Data.ToString()
}

func NewDoubleElement(ename string, data float64) Element {
	e := Element{}
	e.Identifier = byte(1)
	e.EName = CString(ename)
	e.Data = Double(data)
	return e
}

func NewStringElement(ename, data string) Element {
	e := Element{}
	e.Identifier = byte(2)
	e.EName = CString(ename)
	e.Data = String(data)

	return e
}

func NewEmbeddedDocument(ename string, data Document) Element {
	e := Element{}
	e.Identifier = byte(3)
	e.EName = CString(ename)
	e.Data = data
	return e
}

func NewArrayElement(ename string, data []Element) Element {
	e := Element{}
	e.Identifier = byte(4)
	e.EName = CString(ename)
	e.Data = Document(data)
	return e
}

func NewBinaryElement(ename string, data []byte, subtype byte) Element {
	e := Element{}
	e.Identifier = byte(4)
	e.EName = CString(ename)
	e.Data = Binary{subtype, data}
	return e
}

func NewUndefinedElement(ename string) Element {
	e := Element{}
	e.Identifier = byte(6)
	e.EName = CString(ename)
	e.Data = Null(0)

	return e
}

////
func NewObjectIdElement(ename string, data []byte) Element {
	e := Element{}
	e.Identifier = byte(7)
	e.EName = CString(ename)
	e.Data = ObjectId(data)

	return e
}

func NewTrueElement(ename string) Element {
	e := Element{}
	e.Identifier = byte(8)
	e.EName = CString(ename)
	e.Data = Byte(1)

	return e
}

func NewFalseElement(ename string) Element {
	e := Element{}
	e.Identifier = byte(8)
	e.EName = CString(ename)
	e.Data = Byte(0)

	return e
}

func NewUTCElement(ename string, tc time.Time) Element {
	e := Element{}
	e.Identifier = byte('\x09')
	e.EName = CString(ename)
	e.Data = Int64(tc.Unix())

	return e
}

func NewNullElement(ename string) Element {
	e := Element{}
	e.Identifier = byte('\x0A')
	e.EName = CString(ename)
	e.Data = Null(0)

	return e
}

func NewRegularExpressionElement(ename, regexp, options string) Element {
	e := Element{}
	e.Identifier = byte('\x0B')
	e.EName = CString(ename)
	e.Data = RegExp{CString(regexp), CString(options)}

	return e
}

func NewDBPointerElement(ename, namespace string, objectid []byte) Element {
	e := Element{}
	e.Identifier = byte('\x0C')
	e.EName = CString(ename)
	e.Data = DBPointer{String(namespace), ObjectId(objectid)}

	return e
}

func NewJavascriptCodeElement(ename string, code string) Element {
	e := Element{}
	e.Identifier = byte('\x0D')
	e.EName = CString(ename)
	e.Data = String(code)

	return e
}

func NewSymbolElement(ename string, data string) Element {
	e := Element{}
	e.Identifier = byte('\x0E')
	e.EName = CString(ename)
	e.Data = String(data)

	return e
}

// func NewJavascriptWithScopeElement(ename string, data int32) Element {
// 	e := Element{}
// 	e.Identifier = byte('\x0F')
// 	e.EName = CString(ename)
// 	e.Data = Int32(data)

// 	return e
// }

func NewInt32Element(ename string, data int32) Element {
	e := Element{}
	e.Identifier = byte('\x10')
	e.EName = CString(ename)
	e.Data = Int32(data)

	return e
}

func NewTimeStampElement(ename string, tc time.Time) Element {
	e := Element{}
	e.Identifier = byte('\x11')
	e.EName = CString(ename)
	e.Data = Int64(tc.Unix())

	return e
}

func NewInt64Element(ename string, data int64) Element {
	e := Element{}
	e.Identifier = byte('\x12')
	e.EName = CString(ename)
	e.Data = Int64(data)

	return e
}

func NewMinKeyElement(ename string, data int32) Element {
	e := Element{}
	e.Identifier = byte('\xFF')
	e.EName = CString(ename)
	e.Data = Null(0)

	return e
}

func NewMaxKeyElement(ename string, data int32) Element {
	e := Element{}
	e.Identifier = byte('\x7F')
	e.EName = CString(ename)
	e.Data = Null(0)

	return e
}
