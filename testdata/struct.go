package testdata

// binpack:gen
type Struct struct {
	Bool        bool  `binpack:"index=1"`
	BoolPointer *bool `binpack:"index=2"`

	Byte        byte  `binpack:"index=3"`
	BytePointer *byte `binpack:"index=4"`

	UInt          uint    `binpack:"index=5"`
	UIntPointer   *uint   `binpack:"index=6"`
	UInt8         uint8   `binpack:"index=7"`
	UInt8Pointer  *uint8  `binpack:"index=8"`
	UInt16        uint16  `binpack:"index=9"`
	UInt16Pointer *uint16 `binpack:"index=10"`
	UInt32        uint32  `binpack:"index=11"`
	UInt32Pointer *uint32 `binpack:"index=12"`
	UInt64        uint64  `binpack:"index=13"`
	UInt64Pointer *uint64 `binpack:"index=14"`

	Int          int    `binpack:"index=15"`
	IntPointer   *int   `binpack:"index=16"`
	Int8         int8   `binpack:"index=17"`
	Int8Pointer  *int8  `binpack:"index=18"`
	Int16        int16  `binpack:"index=19"`
	Int16Pointer *int16 `binpack:"index=20"`
	Int32        int32  `binpack:"index=21"`
	Int32Pointer *int32 `binpack:"index=22"`
	Int64        int64  `binpack:"index=23"`
	Int64Pointer *int64 `binpack:"index=24"`

	Float32        float32  `binpack:"index=25"`
	Float32Pointer *float32 `binpack:"index=26"`
	Float64        float64  `binpack:"index=27"`
	Float64Pointer *float64 `binpack:"index=28"`

	String          string  `binpack:"index=29"`
	StringPointer   *string `binpack:"index=30"`
	StringFixedSize string  `binpack:"index=31, size=1"`

	// ByteArray        [2]byte  `binpack:"index=32"`
	// ByteArrayPointer *[2]byte `binpack:"index=33"`
	// ByteSlice        []byte   `binpack:"index=34"`
	// ByteSlicePointer *[]byte  `binpack:"index=35"`
}
