package prefixsuffixcheck

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

func NewAnalyzer() (Analyzer *analysis.Analyzer) {
	Analyzer = &analysis.Analyzer{
		Name:     "firstParamtester",
		Doc:      "Check prefix and suffix of functions",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
	return Analyzer
}

// var Analyzer = MakeAnalyzer("Fuck_U")

func CheckPrefifSuffix(template string, pattern1 string, pattern2 string) bool {
	len1 := len(pattern1)
	len2 := len(pattern2)
	len := len(template)
	if len < len1+len2 {
		return false
	}
	var i, j int
	for ; i < len1; i++ {
		if template[i] != pattern1[i] {
			break
		}
	}
	for ; j < len2; j++ {
		if template[len-1-j] != pattern2[len2-1-j] {
			break
		}
	}
	if i < len1 || j < len2 {
		return false
	}
	return true
}
func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}
	prefix := "AA"
	suffix := "BB"
	ff := func(node ast.Node) {
		funcDecl, ok := node.(*ast.FuncDecl)
		if !ok {
			return
		}
		params := funcDecl.Type.Params.List
		if len(params) == 0 {
			return
		}
		curIdent := params[0].Names[0]
		if !CheckPrefifSuffix(curIdent.Name, prefix, suffix) {
			pass.Reportf(node.Pos(), "''%s' function first parameter : %s is not accord with prefix : %s or suffix : %s!\n", funcDecl.Name.Name, curIdent.Name, prefix, suffix)
		}
		return
	}
	inspect.Preorder(nodeFilter, ff)
	return nil, nil
}
