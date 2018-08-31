package binpack

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"os"
	"sort"
	"strings"
)

type Generator struct {
	PackageName string
	Types       []*ast.TypeSpec
	OutName     string
	Buffer      bytes.Buffer
}

func (g *Generator) WriteHeader() {
	g.Buffer.WriteString("// DO NOT EDIT." + "\n")
	g.Buffer.WriteString("\n" + "package " + g.PackageName + "\n")
	g.Buffer.WriteString("\n" + "import (" + "\n")
	g.Buffer.WriteString("\t" + `"encoding/binary"` + "\n")
	g.Buffer.WriteString("\t" + `"errors"` + "\n")
	g.Buffer.WriteString("\t" + `"math"` + "\n")
	g.Buffer.WriteString("\t" + `"strconv"` + "\n")
	g.Buffer.WriteString(")\n")
	g.Buffer.WriteString("\n" + "var (" + "\n")
	g.Buffer.WriteString("\t" + "_ binary.ByteOrder" + "\n")
	g.Buffer.WriteString("\t" + "_ = math.Pi" + "\n")
	g.Buffer.WriteString("\t" + "_ strconv.NumError" + "\n")
	g.Buffer.WriteString(")\n")
}

func (g *Generator) GenerateType(typeSpec *ast.TypeSpec) error {
	switch expression := typeSpec.Type.(type) {
	case *ast.StructType:
		err := g.GenerateStruct(typeSpec.Name.String(), expression)
		if err != nil {
			return err
		}
	default:
		return errors.New("unsupported type: " + fmt.Sprintf("%T", expression))
	}
	return nil
}

func (g *Generator) GenerateStruct(name string, structType *ast.StructType) error {
	var fields []*Field
	for _, astField := range structType.Fields.List {
		field, err := NewField(astField)
		if err != nil {
			return err
		}
		fields = append(fields, field)
	}
	sort.Slice(fields, func(i int, j int) bool {
		return fields[i].Index < fields[j].Index
	})
	lengthParts := []string{}
	marshalParts := []string{}
	unmarshalParts := []string{}
	for _, field := range fields {
		lengthParts = append(lengthParts, field.GenerateMarshalLength())
		marshalParts = append(marshalParts, field.GenerateMarshal())
		unmarshalParts = append(unmarshalParts, field.GenerateUnmarshal())
	}
	g.Buffer.WriteString(`
func (v *` + name + `) MarshalLength() int {
	length := 0` + strings.Join(lengthParts, "") + `
	return length
}

func (v *` + name + `) Marshal() []byte {
	if v == nil {
		return nil
	}
	length := v.MarshalLength()
	if length == 0 {
		return nil
	}
	data := make([]byte, length)
	v.MarshalTo(data)
	return data
}

func (v *` + name + `) MarshalTo(data []byte) {
	i := 0` + strings.Join(marshalParts, "") + `
}

func (v *` + name + `) Unmarshal(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	i := 0
	for {
		switch data[i] {` + strings.Join(unmarshalParts, "") + `
		default:
			return errors.New("Invalid ` + name + ` index: " + strconv.Itoa(int(data[i])))
		}
		if len(data) == i {
			break
		}
	}
	return nil
}
`)
	return nil
}

func (g *Generator) Generate() error {
	g.WriteHeader()
	for _, astType := range g.Types {
		err := g.GenerateType(astType)
		if err != nil {
			return err
		}
	}
	file, err := os.Create(g.OutName)
	if err != nil {
		return err
	}
	_, err = file.Write(g.Buffer.Bytes())
	return err
}
