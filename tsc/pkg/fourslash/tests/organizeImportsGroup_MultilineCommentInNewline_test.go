package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/lsp/lsproto"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestOrganizeImportsGroup_MultilineCommentInNewline(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// polyfill
import c from "C";
/*
* demo
*/
import d from "D";
import a from "A";
import b from "B";

console.log(a, b, c, d)`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(t,
		`// polyfill
import c from "C";
/*
* demo
*/
import a from "A";
import b from "B";
import d from "D";

console.log(a, b, c, d)`,
		lsproto.CodeActionKindSourceOrganizeImportsTs,
		nil,
	)
}
