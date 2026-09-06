package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/lsp/lsproto"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestOrganizeImportsGroup_MultiNewlines(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `import c from "C";


import d from "D";
import a from "A";
import b from "B";

console.log(a, b, c, d)`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(t,
		`import c from "C";


import a from "A";
import b from "B";
import d from "D";

console.log(a, b, c, d)`,
		lsproto.CodeActionKindSourceOrganizeImportsTs,
		nil,
	)
}
