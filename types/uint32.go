package types

import "strconv"

type UInt32 struct {
	Index     uint64
	Name      string
	IsPointer bool
}

func (v *UInt32) GenerateMarshalLength() string {
	data := ""
	if v.IsPointer {
		data += "\n\t" + "if v." + v.Name + " != nil {"
	} else {
		data += "\n\t" + "if v." + v.Name + " != 0 {"
	}
	data += "\n\t\t" + "length += 5"
	data += "\n\t" + "}"
	return data
}

func (v *UInt32) GenerateMarshal() string {
	data := ""
	if v.IsPointer {
		data += "\n\t" + "if v." + v.Name + " != nil {"
	} else {
		data += "\n\t" + "if v." + v.Name + " != 0 {"
	}
	data += "\n\t\t" + "data[i] = " + strconv.FormatUint(v.Index, 10)
	data += "\n\t\t" + "i++"
	if v.IsPointer {
		data += "\n\t\t" + "binary.LittleEndian.PutUint32(data[i:], *v." + v.Name + ")"
	} else {
		data += "\n\t\t" + "binary.LittleEndian.PutUint32(data[i:], v." + v.Name + ")"
	}
	data += "\n\t\t" + "i += 4"
	data += "\n\t" + "}"
	return data
}

func (v *UInt32) GenerateUnmarshal() string {
	data := "\n\t\t" + "case " + strconv.FormatUint(v.Index, 10) + ":"
	data += "\n\t\t\t" + "i++"
	data += "\n\t\t\t" + "if len(data) < i + 4 {"
	data += "\n\t\t\t\t" + `return errors.New("Invalid ` + v.Name + ` data size")`
	data += "\n\t\t\t" + "}"
	if v.IsPointer {
		data += "\n\t\t\t" + "value := binary.LittleEndian.Uint32(data[i:])"
		data += "\n\t\t\t" + "v." + v.Name + " = &value"
	} else {
		data += "\n\t\t\t" + "v." + v.Name + " = binary.LittleEndian.Uint32(data[i:])"
	}
	data += "\n\t\t\t" + "i += 4"
	return data
}

func NewUInt32(index uint64, name string, isPointer bool) Type {
	return &UInt32{
		Index:     index,
		Name:      name,
		IsPointer: isPointer,
	}
}
