package main

import (
	"test/prefixsuffixcheck"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(prefixsuffixcheck.NewAnalyzer())
}
