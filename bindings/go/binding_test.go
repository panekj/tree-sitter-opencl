package tree_sitter_opencl_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-opencl"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_opencl.Language())
	if language == nil {
		t.Errorf("Error loading Opencl grammar")
	}
}
