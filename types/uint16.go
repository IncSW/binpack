package types

import "strconv"

type UInt16 struct {
	Index     uint64
	Name      string
	IsPointer bool
}

func (v *UInt16) GenerateMarshalLength() string {
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

func (v *UInt16) GenerateMarshal() string {
	data := ""
	if v.IsPointer {
		data += "\n\t" + "if v." + v.Name + " != nil {"
	} else {
		data += "\n\t" + "if v." + v.Name + " != 0 {"
	}
	data += "\n\t\t" + "data[i] = " + strconv.FormatUint(v.Index, 10)
	data += "\n\t\t" + "i++"
	if v.IsPointer {
		data += "\n\t\t" + "binary.LittleEndian.PutUint16(data[i:], *v." + v.Name + ")"
	} else {
		data += "\n\t\t" + "binary.LittleEndian.PutUint16(data[i:], v." + v.Name + ")"
	}
	data += "\n\t\t" + "i += 2"
	data += "\n\t" + "}"
	return data
}

func (v *UInt16) GenerateUnmarshal() string {
	data := "\n\t\t" + "case " + strconv.FormatUint(v.Index, 10) + ":"
	data += "\n\t\t\t" + "i++"
	data += "\n\t\t\t" + "if len(data) < i + 2 {"
	data += "\n\t\t\t\t" + `return errors.New("Invalid ` + v.Name + ` data size")`
	data += "\n\t\t\t" + "}"
	if v.IsPointer {
		data += "\n\t\t\t" + "value := binary.LittleEndian.Uint16(data[i:])"
		data += "\n\t\t\t" + "v." + v.Name + " = &value"
	} else {
		data += "\n\t\t\t" + "v." + v.Name + " = binary.LittleEndian.Uint16(data[i:])"
	}
	data += "\n\t\t\t" + "i += 2"
	return data
}

func NewUInt16(index uint64, name string, isPointer bool) Type {
	return &UInt16{
		Index:     index,
		Name:      name,
		IsPointer: isPointer,
	}
}
