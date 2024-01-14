package complexity_test

import (
	"testing"

	complexity "github.com/chlorineCtrl/Complexity-Analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, complexity.Analyzer, []string{"a", ""}...)
}
