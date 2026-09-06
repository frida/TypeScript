package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/ls/lsutil"
	"github.com/frida/TypeScript/tsc/pkg/lsp/lsproto"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestOrganizeImportsType2(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowSyntheticDefaultImports: true
// @moduleResolution: bundler
// @noUnusedLocals: true
// @target: es2018
type A = string;
type B = string;
const C = "hello";
export { A, type B, C };`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(t,
		`type A = string;
type B = string;
const C = "hello";
export { A, C, type B };
`,
		lsproto.CodeActionKindSourceOrganizeImportsTs,
		nil,
	)
	f.VerifyOrganizeImports(t,
		`type A = string;
type B = string;
const C = "hello";
export { A, type B, C };
`,
		lsproto.CodeActionKindSourceOrganizeImportsTs,
		&lsutil.UserPreferences{
			OrganizeImportsTypeOrder: lsutil.OrganizeImportsTypeOrderInline,
		},
	)
	f.VerifyOrganizeImports(t,
		`type A = string;
type B = string;
const C = "hello";
export { type B, A, C };
`,
		lsproto.CodeActionKindSourceOrganizeImportsTs,
		&lsutil.UserPreferences{
			OrganizeImportsTypeOrder: lsutil.OrganizeImportsTypeOrderFirst,
		},
	)
	f.VerifyOrganizeImports(t,
		`type A = string;
type B = string;
const C = "hello";
export { A, C, type B };
`,
		lsproto.CodeActionKindSourceOrganizeImportsTs,
		&lsutil.UserPreferences{
			OrganizeImportsTypeOrder: lsutil.OrganizeImportsTypeOrderLast,
		},
	)
}
