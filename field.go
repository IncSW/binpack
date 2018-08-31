package binpack

import (
	"fmt"
	"go/ast"
	"reflect"
	"strconv"
	"strings"

	"github.com/IncSW/binpack/types"
)

type Field struct {
	Field       *ast.Field
	Name        string
	TypePackage string
	TypeName    string
	Index       uint64
	IsPointer   bool
	Size        uint64
	Type        types.Type
}

func (f *Field) GenerateMarshalLength() string {
	return f.Type.GenerateMarshalLength()
}

func (f *Field) GenerateMarshal() string {
	return f.Type.GenerateMarshal()
}

func (f *Field) GenerateUnmarshal() string {
	return f.Type.GenerateUnmarshal()
}

func (f *Field) Parse() {
	switch expression := f.Field.Type.(type) {
	case *ast.Ident:
		f.TypeName = expression.Name
	case *ast.StarExpr:
		f.IsPointer = true
		switch expression := expression.X.(type) {
		case *ast.Ident:
			f.TypeName = expression.Name
		case *ast.ArrayType:
			f.ParseArrayType(expression)
		default:
			panic("unknown field type\n" + fmt.Sprintf("%#v\n", expression))
		}
	case *ast.ArrayType:
		f.ParseArrayType(expression)
	default:
		panic("unknown field type\n" + fmt.Sprintf("%#v\n", expression))
	}
}

func (f *Field) ParseArrayType(expression *ast.ArrayType) {
	panic("unknown field type\n" + fmt.Sprintf("%#v\n", expression))
}

func NewField(astField *ast.Field) (*Field, error) {
	field := &Field{
		Field: astField,
		Name:  astField.Names[0].Name,
	}

	for _, value := range strings.Split(reflect.StructTag(strings.Trim(astField.Tag.Value, "`")).Get("binpack"), ",") {
		value = strings.TrimSpace(value)
		switch value {
		default:
			keyValue := strings.Split(value, "=")
			if len(keyValue) != 2 {
				continue
			}
			switch strings.ToLower(keyValue[0]) {
			case "index":
				value, err := strconv.ParseUint(keyValue[1], 10, 8)
				if err != nil {
					return nil, err
				}
				field.Index = value
			case "size":
				value, err := strconv.ParseUint(keyValue[1], 10, 8)
				if err != nil {
					return nil, err
				}
				field.Size = value
			}
		}
	}

	field.Parse()

	packType, err := types.NewType(
		field.TypePackage,
		field.TypeName,
		field.Index,
		field.Name,
		field.IsPointer,
		field.Size,
	)
	if err != nil {
		return nil, err
	}
	field.Type = packType

	return field, nil
}
