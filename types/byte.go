package types

import "strconv"

type Byte struct {
	Index     uint64
	Name      string
	IsPointer bool
}

func (v *Byte) GenerateMarshalLength() string {
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

func (v *Byte) GenerateMarshal() string {
	data := ""
	if v.IsPointer {
		data += "\n\t" + "if v." + v.Name + " != nil {"
	} else {
		data += "\n\t" + "if v." + v.Name + " != 0 {"
	}
	data += "\n\t\t" + "data[i] = " + strconv.FormatUint(v.Index, 10)
	data += "\n\t\t" + "i++"
	if v.IsPointer {
		data += "\n\t\t" + "data[i] = *v." + v.Name
	} else {
		data += "\n\t\t" + "data[i] = v." + v.Name
	}
	data += "\n\t\t" + "i++"
	data += "\n\t" + "}"
	return data
}

func (v *Byte) GenerateUnmarshal() string {
	data := "\n\t\t" + "case " + strconv.FormatUint(v.Index, 10) + ":"
	data += "\n\t\t\t" + "i++"
	data += "\n\t\t\t" + "if len(data) < i {"
	data += "\n\t\t\t\t" + `return errors.New("Invalid ` + v.Name + ` data size")`
	data += "\n\t\t\t" + "}"
	if v.IsPointer {
		data += "\n\t\t\t" + "value := data[i]"
		data += "\n\t\t\t" + "v." + v.Name + " = &value"
	} else {
		data += "\n\t\t\t" + "v." + v.Name + " = data[i]"
	}
	data += "\n\t\t\t" + "i++"
	return data
}

func NewByte(index uint64, name string, isPointer bool) Type {
	return &Byte{
		Index:     index,
		Name:      name,
		IsPointer: isPointer,
	}
}
