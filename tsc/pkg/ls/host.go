package ls

import (
	"github.com/frida/TypeScript/tsc/pkg/ls/autoimport"
	"github.com/frida/TypeScript/tsc/pkg/ls/lsconv"
	"github.com/frida/TypeScript/tsc/pkg/ls/lsutil"
	"github.com/frida/TypeScript/tsc/pkg/sourcemap"
)

type Host interface {
	UseCaseSensitiveFileNames() bool
	ReadFile(path string) (contents string, ok bool)
	Converters() *lsconv.Converters
	GetPreferences(activeFile string) lsutil.UserPreferences
	GetECMALineInfo(fileName string) *sourcemap.ECMALineInfo
	AutoImportRegistry() *autoimport.Registry

	// Used for module specifier completions.
	// ! Do not use for anything else, as this violates the principle that
	// the host is a snapshot-in-time.
	ReadDirectory(currentDir string, path string, extensions []string, excludes []string, includes []string, depth int) []string
	GetDirectories(path string) []string
	DirectoryExists(path string) bool
	FileExists(path string) bool
}
