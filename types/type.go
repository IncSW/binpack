package types

import "errors"

type Type interface {
	GenerateMarshalLength() string
	GenerateMarshal() string
	GenerateUnmarshal() string
}

func NewType(packageName string, typeName string, index uint64, name string, isPointer bool, size uint64) (Type, error) {
	switch packageName {
	case "":
		switch typeName {
		case "bool":
			return NewBool(index, name, isPointer), nil

		case "byte":
			return NewByte(index, name, isPointer), nil

		case "uint":
			return NewUInt(index, name, isPointer), nil
		case "uint8":
			return NewUInt8(index, name, isPointer), nil
		case "uint16":
			return NewUInt16(index, name, isPointer), nil
		case "uint32":
			return NewUInt32(index, name, isPointer), nil
		case "uint64":
			return NewUInt64(index, name, isPointer), nil

		case "int":
			return NewInt(index, name, isPointer), nil
		case "int8":
			return NewInt8(index, name, isPointer), nil
		case "int16":
			return NewInt16(index, name, isPointer), nil
		case "int32":
			return NewInt32(index, name, isPointer), nil
		case "int64":
			return NewInt64(index, name, isPointer), nil

		case "float32":
			return NewFloat32(index, name, isPointer), nil
		case "float64":
			return NewFloat64(index, name, isPointer), nil

		case "string":
			return NewString(index, name, isPointer, size), nil

		default:
			return nil, errors.New("type not found: " + typeName)
		}
	default:
		return nil, errors.New("package not found: " + packageName)
	}
}
