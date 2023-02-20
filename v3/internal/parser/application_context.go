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
	FileContexts    []fileContext
	Options         *AstCompositeLit
	BoundCandidates []ast.Expr
}

func NewApplicationContext(root string) *ApplicationContext {
	var contexts []fileContext
	contexts = append(contexts, loadDir(root)...)

	appContext := ApplicationContext{
		FileContexts: contexts,
	}
	appContext.build()
	return &appContext
}

func loadDir(path string) []fileContext {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, path, nil, parser.AllErrors)
	if err != nil {
		log.Printf("unable to parse directory %s: %v", path, err)
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
		}
	}

	return contexts
}

func (app *ApplicationContext) build() {
	for _, f := range app.FileContexts {
		file := f.file
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

		app.Options = lits[0]
		for _, option := range app.Options.Elts {
			temp, ok := option.(*ast.KeyValueExpr)
			if !ok {
				msg := fmt.Sprintf("Found non-key for %s", app.Options.String())
				panic(msg)
			}

			key := AstKeyValueExpr(*temp)
			if key.Name() != "Bind" {
				continue
			}

			app.BoundCandidates = key.ArrayValue()
			fmt.Printf("Found %d candidates for bound structures\n", len(app.BoundCandidates))
		}
	}
}
