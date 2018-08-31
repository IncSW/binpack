package testdata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalUnmarshal(t *testing.T) {
	value := &Struct{
		Bool: true,

		Byte: 1,

		UInt:   2,
		UInt8:  3,
		UInt16: 4,
		UInt32: 5,
		UInt64: 6,

		Int:   7,
		Int8:  8,
		Int16: 9,
		Int32: 10,
		Int64: 11,

		Float32: 12.1,
		Float64: 13.2,

		String:          "foobar",
		StringFixedSize: "foo",
	}

	data := value.Marshal()
	empty := &Struct{}
	err := empty.Unmarshal(data)
	if err != nil {
		t.Fatal(err)
	}
	if !assert.Equal(t, value, empty) {
		return
	}

	value.BoolPointer = &value.Bool

	value.BytePointer = &value.Byte

	value.UIntPointer = &value.UInt
	value.UInt8Pointer = &value.UInt8
	value.UInt16Pointer = &value.UInt16
	value.UInt32Pointer = &value.UInt32
	value.UInt64Pointer = &value.UInt64

	value.IntPointer = &value.Int
	value.Int8Pointer = &value.Int8
	value.Int16Pointer = &value.Int16
	value.Int32Pointer = &value.Int32
	value.Int64Pointer = &value.Int64

	value.Float32Pointer = &value.Float32
	value.Float64Pointer = &value.Float64

	value.StringPointer = &value.String

	data = value.Marshal()
	empty = &Struct{}
	err = empty.Unmarshal(data)
	if err != nil {
		t.Fatal(err)
	}
	if !assert.Equal(t, value, empty) {
		return
	}
}

func BenchmarkMarshalUnmarshal(b *testing.B) {
	value := &Struct{
		Bool: true,

		Byte: 1,

		UInt:   2,
		UInt8:  3,
		UInt16: 4,
		UInt32: 5,
		UInt64: 6,

		Int:   7,
		Int8:  8,
		Int16: 9,
		Int32: 10,
		Int64: 11,

		Float32: 12.1,
		Float64: 13.2,

		String:          "foobar",
		StringFixedSize: "foo",
	}
	data := value.Marshal()
	empty := &Struct{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		value.Marshal()
		empty.Unmarshal(data)
	}
}
