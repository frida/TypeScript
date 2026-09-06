package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/core"
	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/ls/lsutil"
	"github.com/frida/TypeScript/tsc/pkg/lsp/lsproto"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestOrganizeImportsPathsUnicode2(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `import * as a2 from "./a2";
import * as a100 from "./a100";
import * as a1 from "./a1";

console.log(a1, a2, a100);`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(t,
		`import * as a1 from "./a1";
import * as a100 from "./a100";
import * as a2 from "./a2";

console.log(a1, a2, a100);`,
		lsproto.CodeActionKindSourceOrganizeImportsTs,
		&lsutil.UserPreferences{
			OrganizeImportsIgnoreCase:       core.TSFalse,
			OrganizeImportsCollation:        lsutil.OrganizeImportsCollationUnicode,
			OrganizeImportsNumericCollation: core.TSFalse,
		},
	)
	f.VerifyOrganizeImports(t,
		`import * as a1 from "./a1";
import * as a2 from "./a2";
import * as a100 from "./a100";

console.log(a1, a2, a100);`,
		lsproto.CodeActionKindSourceOrganizeImportsTs,
		&lsutil.UserPreferences{
			OrganizeImportsIgnoreCase:       core.TSFalse,
			OrganizeImportsCollation:        lsutil.OrganizeImportsCollationUnicode,
			OrganizeImportsNumericCollation: core.TSTrue,
		},
	)
}
