package cpp

//#include "parser.h"
//TSLanguage *tree_sitter_cpp();
import "C"
import (
	"unsafe"

	sitter "github.com/octoberswimmer/go-tree-sitter"
)

func GetLanguage() *sitter.Language {
	ptr := unsafe.Pointer(C.tree_sitter_cpp())
	return sitter.NewLanguage(ptr)
}
