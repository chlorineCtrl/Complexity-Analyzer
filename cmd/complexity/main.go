package main

import (
	complexity "github.com/chlorineCtrl/Complexity-Analyzer"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(complexity.Analyzer) }
