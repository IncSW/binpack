package types

import "strconv"

type String struct {
	Index     uint64
	Name      string
	IsPointer bool
	Size      uint64
}

func (v *String) GenerateMarshalLength() string {
	data := ""
	if v.IsPointer {
		data += "\n\t" + "if v." + v.Name + " != nil {"
		data += "\n\t\t" + "length += " + strconv.FormatUint(1+v.Size, 10) + " + len(*v." + v.Name + ")"
	} else {
		data += "\n\t" + "if len(v." + v.Name + ") != 0 {"
		data += "\n\t\t" + "length += " + strconv.FormatUint(1+v.Size, 10) + " + len(v." + v.Name + ")"
	}
	data += "\n\t" + "}"
	return data
}

func (v *String) GenerateMarshal() string {
	data := ""
	if v.IsPointer {
		data += "\n\t" + "if v." + v.Name + " != nil {"
	} else {
		data += "\n\t" + "if len(v." + v.Name + ") != 0 {"
	}
	data += "\n\t\t" + "data[i] = " + strconv.FormatUint(v.Index, 10)
	data += "\n\t\t" + "i++"
	field := "v." + v.Name
	if v.IsPointer {
		field = "*" + field
	}
	switch v.Size {
	case 1:
		data += "\n\t\t" + "data[i] = len(" + field + ")"
	case 2:
		data += "\n\t\t" + "binary.LittleEndian.PutUint16(data[i:], uint16(len(" + field + ")))"
	case 4:
		data += "\n\t\t" + "binary.LittleEndian.PutUint32(data[i:], uint32(len(" + field + ")))"
	case 8:
		data += "\n\t\t" + "binary.LittleEndian.PutUint64(data[i:], uint64(len(" + field + ")))"
	}
	data += "\n\t\t" + "i += " + strconv.FormatUint(v.Size, 10)
	data += "\n\t\t" + "copy(data[i:], " + field + ")"
	data += "\n\t\t" + "i += len(" + field + ")"
	data += "\n\t" + "}"
	return data
}

func (v *String) GenerateUnmarshal() string {
	data := "\n\t\t" + "case " + strconv.FormatUint(v.Index, 10) + ":"
	data += "\n\t\t\t" + "i++"
	data += "\n\t\t\t" + "if len(data) < i + " + strconv.FormatUint(v.Size, 10) + " {"
	data += "\n\t\t\t\t" + `return errors.New("Invalid ` + v.Name + ` data size")`
	data += "\n\t\t\t" + "}"
	switch v.Size {
	case 1:
		data += "\n\t\t\t" + "length := int(data[i])"
	case 2:
		data += "\n\t\t\t" + "length := int(binary.LittleEndian.Uint16(data[i:]))"
	case 4:
		data += "\n\t\t\t" + "length := int(binary.LittleEndian.Uint32(data[i:]))"
	case 8:
		data += "\n\t\t\t" + "length := int(binary.LittleEndian.Uint64(data[i:]))"
	}
	data += "\n\t\t\t" + "i += " + strconv.FormatUint(v.Size, 10)
	data += "\n\t\t\t" + "if len(data) < i + length {"
	data += "\n\t\t\t\t" + `return errors.New("Invalid ` + v.Name + ` length size")`
	data += "\n\t\t\t" + "}"
	if v.IsPointer {
		data += "\n\t\t\t" + "value := string(data[i : i+length])"
		data += "\n\t\t\t" + "v." + v.Name + " = &value"
	} else {
		data += "\n\t\t\t" + "v." + v.Name + " = string(data[i : i+length])"
	}
	data += "\n\t\t\t" + "i += length"
	return data
}

func NewString(index uint64, name string, isPointer bool, size uint64) Type {
	switch size {
	case 1:
	case 2:
	case 4:
	case 8:
	default:
		size = 4
	}
	return &String{
		Index:     index,
		Name:      name,
		IsPointer: isPointer,
		Size:      size,
	}
}
