package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/core"
	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/ls/lsutil"
	"github.com/frida/TypeScript/tsc/pkg/lsp/lsproto"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestOrganizeImports16(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `import { a, A, b } from "foo";
interface Use extends A {}
console.log(a, b);`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(t,
		`import { a, A, b } from "foo";
interface Use extends A {}
console.log(a, b);`,
		lsproto.CodeActionKindSourceOrganizeImportsTs,
		nil,
	)
	f.ReplaceLine(t, 0, "import { a, A, b } from \"foo1\";")
	f.VerifyOrganizeImports(t,
		`import { a, A, b } from "foo1";
interface Use extends A {}
console.log(a, b);`,
		lsproto.CodeActionKindSourceOrganizeImportsTs,
		&lsutil.UserPreferences{
			OrganizeImportsIgnoreCase: core.TSUnknown,
		},
	)
	f.ReplaceLine(t, 0, "import { a, A, b } from \"foo2\";")
	f.VerifyOrganizeImports(t,
		`import { a, A, b } from "foo2";
interface Use extends A {}
console.log(a, b);`,
		lsproto.CodeActionKindSourceOrganizeImportsTs,
		&lsutil.UserPreferences{
			OrganizeImportsIgnoreCase: core.TSTrue,
		},
	)
	f.ReplaceLine(t, 0, "import { a, A, b } from \"foo3\";")
	f.VerifyOrganizeImports(t,
		`import { A, a, b } from "foo3";
interface Use extends A {}
console.log(a, b);`,
		lsproto.CodeActionKindSourceOrganizeImportsTs,
		&lsutil.UserPreferences{
			OrganizeImportsIgnoreCase: core.TSFalse,
		},
	)
}
