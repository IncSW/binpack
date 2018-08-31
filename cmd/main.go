package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/IncSW/binpack"
)

const suffix = "_binpack"

func main() {
	pathname := "testdata"
	err := generate(pathname)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func generate(pathname string) error {
	fileInfo, err := os.Stat(pathname)
	if err != nil {
		return err
	}

	if !fileInfo.IsDir() && !strings.HasSuffix(pathname, ".go") {
		return errors.New("Filename must end in '.go'")
	}

	parser := &binpack.Parser{}
	err = parser.Parse(pathname, fileInfo.IsDir())
	if err != nil {
		return err
	}

	outName := ""
	if fileInfo.IsDir() {
		outName = filepath.Join(pathname, parser.PackageName+suffix+".go")
	} else {
		outName = strings.TrimSuffix(pathname, ".go") + suffix + ".go"
	}

	generator := &binpack.Generator{
		PackageName: parser.PackageName,
		Types:       parser.Types,
		OutName:     outName,
	}

	return generator.Generate()
}
