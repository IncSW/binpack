package binpack

import (
	"go/ast"
	"strings"
)

const explicitComment = "binpack:gen"

type Visitor struct {
	*Parser
	explicit bool
}

func (v *Visitor) Visit(node ast.Node) ast.Visitor {
	switch node := node.(type) {
	case *ast.Package:
		return v
	case *ast.File:
		v.PackageName = node.Name.String()
		return v
	case *ast.GenDecl:
		v.explicit = isExplicit(node.Doc.Text())
		if !v.explicit {
			return nil
		}
		return v
	case *ast.TypeSpec:
		if v.explicit {
			v.Types = append(v.Types, node)
			return nil
		}
		return v
	}
	return nil
}

func isExplicit(comments string) bool {
	for _, comment := range strings.Split(comments, "\n") {
		if strings.HasPrefix(comment, explicitComment) {
			return true
		}
	}
	return false
}
