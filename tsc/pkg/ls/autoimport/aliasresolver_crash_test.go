package autoimport

import (
	"context"
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/ast"
	"github.com/frida/TypeScript/tsc/pkg/binder"
	"github.com/frida/TypeScript/tsc/pkg/checker"
	"github.com/frida/TypeScript/tsc/pkg/compiler"
	"github.com/frida/TypeScript/tsc/pkg/core"
	"github.com/frida/TypeScript/tsc/pkg/module"
	"github.com/frida/TypeScript/tsc/pkg/packagejson"
	"github.com/frida/TypeScript/tsc/pkg/parser"
	"github.com/frida/TypeScript/tsc/pkg/tspath"
	"github.com/frida/TypeScript/tsc/pkg/vfs"
	"github.com/frida/TypeScript/tsc/pkg/vfs/vfstest"
)

type fakeCloneHost struct {
	fs vfs.FS
}

func (h *fakeCloneHost) FS() vfs.FS                  { return h.fs }
func (h *fakeCloneHost) GetCurrentDirectory() string { return "/" }
func (h *fakeCloneHost) GetDefaultProject(path tspath.Path) (tspath.Path, *compiler.Program) {
	return "", nil
}

func (h *fakeCloneHost) GetProgramForProject(projectPath tspath.Path) *compiler.Program { return nil }

func (h *fakeCloneHost) GetPackageJson(fileName string) *packagejson.InfoCacheEntry { return nil }

func (h *fakeCloneHost) GetSourceFile(fileName string, path tspath.Path) *ast.SourceFile { return nil }
func (h *fakeCloneHost) Dispose()                                                        {}

var _ RegistryCloneHost = (*fakeCloneHost)(nil)

// Regression test for microsoft/typescript-go#4322.
//
// During auto-import export extraction, the checker is built on top of an
// aliasResolver standing in for a real program. This file has a type error, and
// extracting exports should still complete without crashing.
func TestAliasResolverGetDiagnosticsDoesNotPanic(t *testing.T) {
	t.Parallel()

	const fileName = "/pkg/index.ts"
	text := "declare function f(arg: { a: string }): () => void;\nexport const x = f({ a: 1 });\n"

	fs := vfstest.FromMap(map[string]string{fileName: text}, true /*useCaseSensitiveFileNames*/)
	host := &fakeCloneHost{fs: fs}

	sourceFile := parser.ParseSourceFile(ast.SourceFileParseOptions{
		FileName: fileName,
		Path:     tspath.Path(fileName),
	}, text, core.ScriptKindTS)
	binder.BindSourceFile(sourceFile)

	resolver := module.NewResolver(host, core.EmptyCompilerOptions, "", "", nil)
	r := newAliasResolver(
		[]*ast.SourceFile{sourceFile},
		nil,
		host,
		resolver,
		func(f string) tspath.Path { return tspath.Path(f) },
		func(ast.HasFileName, string) {},
	)

	ch, _ := checker.NewChecker(r, nil)

	// Type-checking this file's diagnostics must not panic.
	ch.GetDiagnostics(context.Background(), sourceFile)
}
