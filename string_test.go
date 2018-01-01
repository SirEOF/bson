package bson

import "testing"

func TestCString(t *testing.T) {
	a1 := CString("test123!")

	SerializableTest(t, &a1)
}

func TestString(t *testing.T) {
	a1 := String("test123!")

	SerializableTest(t, &a1)
}
