package binpack

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

type Parser struct {
	PackageName string
	Types       []*ast.TypeSpec
}

func (p *Parser) Parse(pathname string, isDir bool) error {
	fileSet := token.NewFileSet()
	if isDir {
		packages, err := parser.ParseDir(fileSet, pathname, func(info os.FileInfo) bool {
			return !strings.HasSuffix(info.Name(), "_binpack.go")
		}, parser.ParseComments)
		if err != nil {
			return err
		}
		for _, astPackage := range packages {
			visitor := &Visitor{
				Parser: p,
			}
			ast.Walk(visitor, astPackage)
		}
		return nil
	}
	file, err := parser.ParseFile(fileSet, pathname, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	visitor := &Visitor{
		Parser: p,
	}
	ast.Walk(visitor, file)
	return nil
}
