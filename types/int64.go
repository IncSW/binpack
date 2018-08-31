package types

import "strconv"

type Int64 struct {
	Index     uint64
	Name      string
	IsPointer bool
}

func (v *Int64) GenerateMarshalLength() string {
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

func (v *Int64) GenerateMarshal() string {
	data := ""
	if v.IsPointer {
		data += "\n\t" + "if v." + v.Name + " != nil {"
	} else {
		data += "\n\t" + "if v." + v.Name + " != 0 {"
	}
	data += "\n\t\t" + "data[i] = " + strconv.FormatUint(v.Index, 10)
	data += "\n\t\t" + "i++"
	if v.IsPointer {
		data += "\n\t\t" + "binary.LittleEndian.PutUint64(data[i:], uint64(*v." + v.Name + "))"
	} else {
		data += "\n\t\t" + "binary.LittleEndian.PutUint64(data[i:], uint64(v." + v.Name + "))"
	}
	data += "\n\t\t" + "i += 8"
	data += "\n\t" + "}"
	return data
}

func (v *Int64) GenerateUnmarshal() string {
	data := "\n\t\t" + "case " + strconv.FormatUint(v.Index, 10) + ":"
	data += "\n\t\t\t" + "i++"
	data += "\n\t\t\t" + "if len(data) < i + 8 {"
	data += "\n\t\t\t\t" + `return errors.New("Invalid ` + v.Name + ` data size")`
	data += "\n\t\t\t" + "}"
	if v.IsPointer {
		data += "\n\t\t\t" + "value := int64(binary.LittleEndian.Uint64(data[i:]))"
		data += "\n\t\t\t" + "v." + v.Name + " = &value"
	} else {
		data += "\n\t\t\t" + "v." + v.Name + " = int64(binary.LittleEndian.Uint64(data[i:]))"
	}
	data += "\n\t\t\t" + "i += 8"
	return data
}

func NewInt64(index uint64, name string, isPointer bool) Type {
	return &Int64{
		Index:     index,
		Name:      name,
		IsPointer: isPointer,
	}
}
