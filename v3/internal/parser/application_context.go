package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

type fileContext struct {
	pkgName  string
	fileName string
	file     *AstFile
}

type ApplicationContext struct {
	files []fileContext
}

func NewApplicationContext(root string) *ApplicationContext {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, root, nil, parser.AllErrors)
	if err != nil {
		log.Printf("unable to parse directory %s: %v", root, err)
		return nil
	}

	var contexts []fileContext
	for name, pkg := range pkgs {
		for _, f := range pkg.Files {
			file := NewAstFile(f)
			context := fileContext{
				pkgName:  name,
				fileName: file.Name.Name,
				file:     file,
			}
			contexts = append(contexts, context)

			wailsImport, ok := file.ImportMap[WailsApplicationPackage]
			if !ok {
				continue
			}

			var target string
			switch {
			// default import
			case wailsImport == nil:
				target = WailsApplicationImport + ".Options"
			case *wailsImport == "_":
				target = "Options"
			default:
				target = *wailsImport + ".Options"
			}

			fmt.Printf("Looking for %s\n", target)
			var options *AstCompositeLit
			ast.Inspect(f, func(node ast.Node) bool {
				compositeLit, ok := node.(*ast.CompositeLit)
				if !ok {
					return true
				}

				t := AstCompositeLit(*compositeLit)
				if t.String() == target {
					options = &t
					return true
				}

				return true
			})

			fmt.Printf("Found what we're looking for: %s\n", options)
		}
	}

	return &ApplicationContext{files: contexts}
}