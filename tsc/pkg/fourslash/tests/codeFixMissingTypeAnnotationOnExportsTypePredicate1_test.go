package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestCodeFixMissingTypeAnnotationOnExportsTypePredicate1(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @isolatedDeclarations: true
// @declaration: true
// @filename: index.ts
export function isString(value: unknown) {
  return typeof value === "string";
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCodeFix(t, fourslash.VerifyCodeFixOptions{
		Description: "Add return type 'value is string'",
		NewFileContent: `export function isString(value: unknown): value is string {
  return typeof value === "string";
}`,
		Index: 0,
	})
}
