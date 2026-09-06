package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/core"
	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/ls/lsutil"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestInlayHintsReturnType(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function foo1 () {
    return 1
}
function foo2 (): number {
    return 1
}
class C {
    foo() {
        return 1
    }
}
const a = () => 1
const b = function () { return 1 }
const c = (b) => 1
const d = b => 1`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineInlayHints(t, nil /*span*/, &lsutil.UserPreferences{InlayHints: lsutil.InlayHintsPreferences{IncludeInlayFunctionLikeReturnTypeHints: core.TSTrue}})
}
