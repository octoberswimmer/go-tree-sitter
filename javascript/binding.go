package javascript

//#include "parser.h"
//TSLanguage *tree_sitter_javascript();
import "C"
import (
	"unsafe"

	sitter "github.com/octoberswimmer/go-tree-sitter"
)

func GetLanguage() *sitter.Language {
	ptr := unsafe.Pointer(C.tree_sitter_javascript())
	return sitter.NewLanguage(ptr)
}
