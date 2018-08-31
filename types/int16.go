package types

import "strconv"

type Int16 struct {
	Index     uint64
	Name      string
	IsPointer bool
}

func (v *Int16) GenerateMarshalLength() string {
	data := ""
	if v.IsPointer {
		data += "\n\t" + "if v." + v.Name + " != nil {"
	} else {
		data += "\n\t" + "if v." + v.Name + " != 0 {"
	}
	data += "\n\t\t" + "length += 3"
	data += "\n\t" + "}"
	return data
}

func (v *Int16) GenerateMarshal() string {
	data := ""
	if v.IsPointer {
		data += "\n\t" + "if v." + v.Name + " != nil {"
	} else {
		data += "\n\t" + "if v." + v.Name + " != 0 {"
	}
	data += "\n\t\t" + "data[i] = " + strconv.FormatUint(v.Index, 10)
	data += "\n\t\t" + "i++"
	if v.IsPointer {
		data += "\n\t\t" + "binary.LittleEndian.PutUint16(data[i:], uint16(*v." + v.Name + "))"
	} else {
		data += "\n\t\t" + "binary.LittleEndian.PutUint16(data[i:], uint16(v." + v.Name + "))"
	}
	data += "\n\t\t" + "i += 2"
	data += "\n\t" + "}"
	return data
}

func (v *Int16) GenerateUnmarshal() string {
	data := "\n\t\t" + "case " + strconv.FormatUint(v.Index, 10) + ":"
	data += "\n\t\t\t" + "i++"
	data += "\n\t\t\t" + "if len(data) < i + 2 {"
	data += "\n\t\t\t\t" + `return errors.New("Invalid ` + v.Name + ` data size")`
	data += "\n\t\t\t" + "}"
	if v.IsPointer {
		data += "\n\t\t\t" + "value := int16(binary.LittleEndian.Uint16(data[i:]))"
		data += "\n\t\t\t" + "v." + v.Name + " = &value"
	} else {
		data += "\n\t\t\t" + "v." + v.Name + " = int16(binary.LittleEndian.Uint16(data[i:]))"
	}
	data += "\n\t\t\t" + "i += 2"
	return data
}

func NewInt16(index uint64, name string, isPointer bool) Type {
	return &Int16{
		Index:     index,
		Name:      name,
		IsPointer: isPointer,
	}
}
