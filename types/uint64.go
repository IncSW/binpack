package types

import "strconv"

type UInt64 struct {
	Index     uint64
	Name      string
	IsPointer bool
}

func (v *UInt64) GenerateMarshalLength() string {
	data := ""
	if v.IsPointer {
		data += "\n\t" + "if v." + v.Name + " != nil {"
	} else {
		data += "\n\t" + "if v." + v.Name + " != 0 {"
	}
	data += "\n\t\t" + "length += 9"
	data += "\n\t" + "}"
	return data
}

func (v *UInt64) GenerateMarshal() string {
	data := ""
	if v.IsPointer {
		data += "\n\t" + "if v." + v.Name + " != nil {"
	} else {
		data += "\n\t" + "if v." + v.Name + " != 0 {"
	}
	data += "\n\t\t" + "data[i] = " + strconv.FormatUint(v.Index, 10)
	data += "\n\t\t" + "i++"
	if v.IsPointer {
		data += "\n\t\t" + "binary.LittleEndian.PutUint64(data[i:], *v." + v.Name + ")"
	} else {
		data += "\n\t\t" + "binary.LittleEndian.PutUint64(data[i:], v." + v.Name + ")"
	}
	data += "\n\t\t" + "i += 8"
	data += "\n\t" + "}"
	return data
}

func (v *UInt64) GenerateUnmarshal() string {
	data := "\n\t\t" + "case " + strconv.FormatUint(v.Index, 10) + ":"
	data += "\n\t\t\t" + "i++"
	data += "\n\t\t\t" + "if len(data) < i + 8 {"
	data += "\n\t\t\t\t" + `return errors.New("Invalid ` + v.Name + ` data size")`
	data += "\n\t\t\t" + "}"
	if v.IsPointer {
		data += "\n\t\t\t" + "value := binary.LittleEndian.Uint64(data[i:])"
		data += "\n\t\t\t" + "v." + v.Name + " = &value"
	} else {
		data += "\n\t\t\t" + "v." + v.Name + " = binary.LittleEndian.Uint64(data[i:])"
	}
	data += "\n\t\t\t" + "i += 8"
	return data
}

func NewUInt64(index uint64, name string, isPointer bool) Type {
	return &UInt64{
		Index:     index,
		Name:      name,
		IsPointer: isPointer,
	}
}
