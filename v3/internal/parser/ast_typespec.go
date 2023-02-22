package parser

import "go/ast"

type AstTypeSpec ast.TypeSpec

func (spec *AstTypeSpec) String() string {
	return spec.Name.Name
}
