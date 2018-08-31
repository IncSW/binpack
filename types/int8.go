package types

import "strconv"

type Int8 struct {
	Index     uint64
	Name      string
	IsPointer bool
}

func (v *Int8) GenerateMarshalLength() string {
	data := ""
	if v.IsPointer {
		data += "\n\t" + "if v." + v.Name + " != nil {"
	} else {
		data += "\n\t" + "if v." + v.Name + " != 0 {"
	}
	data += "\n\t\t" + "length += 2"
	data += "\n\t" + "}"
	return data
}

func (v *Int8) GenerateMarshal() string {
	data := ""
	if v.IsPointer {
		data += "\n\t" + "if v." + v.Name + " != nil {"
	} else {
		data += "\n\t" + "if v." + v.Name + " != 0 {"
	}
	data += "\n\t\t" + "data[i] = " + strconv.FormatUint(v.Index, 10)
	data += "\n\t\t" + "i++"
	if v.IsPointer {
		data += "\n\t\t" + "data[i] = byte(*v." + v.Name + ")"
	} else {
		data += "\n\t\t" + "data[i] = byte(v." + v.Name + ")"
	}
	data += "\n\t\t" + "i++"
	data += "\n\t" + "}"
	return data
}

func (v *Int8) GenerateUnmarshal() string {
	data := "\n\t\t" + "case " + strconv.FormatUint(v.Index, 10) + ":"
	data += "\n\t\t\t" + "i++"
	data += "\n\t\t\t" + "if len(data) < i {"
	data += "\n\t\t\t\t" + `return errors.New("Invalid ` + v.Name + ` data size")`
	data += "\n\t\t\t" + "}"
	if v.IsPointer {
		data += "\n\t\t\t" + "value := int8(data[i])"
		data += "\n\t\t\t" + "v." + v.Name + " = &value"
	} else {
		data += "\n\t\t\t" + "v." + v.Name + " = int8(data[i])"
	}
	data += "\n\t\t\t" + "i++"
	return data
}

func NewInt8(index uint64, name string, isPointer bool) Type {
	return &Int8{
		Index:     index,
		Name:      name,
		IsPointer: isPointer,
	}
}
