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
	Files           []fileContext
	Options         *AstCompositeLit
	BoundCandidates []ast.Expr
}

func NewApplicationContext(root string) *ApplicationContext {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, root, nil, parser.AllErrors)
	if err != nil {
		log.Printf("unable to parse directory %s: %v", root, err)
		return nil
	}

	var contexts []fileContext
	var options *AstCompositeLit
	var candidates []ast.Expr
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

			lits, ok := file.CompositeLits[target]
			if !ok {
				continue
			}

			if len(lits) > 1 {
				fmt.Printf("Found %d CompositeLits named %s\n", len(lits), target)
			}

			options = lits[0]
			for _, option := range options.Elts {
				temp, ok := option.(*ast.KeyValueExpr)
				if !ok {
					msg := fmt.Sprintf("Found non-key for %s", options.String())
					panic(msg)
				}

				key := AstKeyValueExpr(*temp)
				if key.Name() != "Bind" {
					continue
				}

				candidates = key.ArrayValue()
				fmt.Printf("Found %d candidates for bound structures\n", len(candidates))
			}
		}
	}

	return &ApplicationContext{
		Files:           contexts,
		Options:         options,
		BoundCandidates: candidates,
	}
}
