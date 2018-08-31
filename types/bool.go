package types

import "strconv"

type Bool struct {
	Index     uint64
	Name      string
	IsPointer bool
}

func (v *Bool) GenerateMarshalLength() string {
	data := ""
	if v.IsPointer {
		data += "\n\t" + "if v." + v.Name + " != nil {"
		data += "\n\t\t" + "length += 2"
	} else {
		data += "\n\t" + "if v." + v.Name + " {"
		data += "\n\t\t" + "length += 1"
	}
	data += "\n\t" + "}"
	return data
}

func (v *Bool) GenerateMarshal() string {
	data := ""
	if v.IsPointer {
		data += "\n\t" + "if v." + v.Name + " != nil {"
	} else {
		data += "\n\t" + "if v." + v.Name + " {"
	}
	data += "\n\t\t" + "data[i] = " + strconv.FormatUint(v.Index, 10)
	data += "\n\t\t" + "i++"
	if v.IsPointer {
		data += "\n\t\t" + "if *v." + v.Name + " {"
		data += "\n\t\t\t" + "data[i] = 1"
		data += "\n\t\t" + "}"
		data += "\n\t\t" + "i++"
	}
	data += "\n\t" + "}"
	return data
}

func (v *Bool) GenerateUnmarshal() string {
	data := "\n\t\t" + "case " + strconv.FormatUint(v.Index, 10) + ":"
	data += "\n\t\t\t" + "i++"
	data += "\n\t\t\t" + "if len(data) < i {"
	data += "\n\t\t\t\t" + `return errors.New("Invalid ` + v.Name + ` data size")`
	data += "\n\t\t\t" + "}"
	if v.IsPointer {
		data += "\n\t\t\t" + "value := data[i] == 1"
		data += "\n\t\t\t" + "v." + v.Name + " = &value"
		data += "\n\t\t\t" + "i++"
	} else {
		data += "\n\t\t\t" + "v." + v.Name + " = true"
	}
	return data
}

func NewBool(index uint64, name string, isPointer bool) Type {
	return &Bool{
		Index:     index,
		Name:      name,
		IsPointer: isPointer,
	}
}
