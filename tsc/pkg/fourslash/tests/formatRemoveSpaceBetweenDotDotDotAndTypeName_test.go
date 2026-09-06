package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestFormatRemoveSpaceBetweenDotDotDotAndTypeName(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `let a: [... any[]];
let b: [...   number[]];
let c: [...     string[]];`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.FormatDocument(t, "")
	f.VerifyCurrentFileContent(t, `let a: [...any[]];
let b: [...number[]];
let c: [...string[]];`)
}
