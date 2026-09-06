package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestCodeFixClassImplementInterfaceIndexSignaturesNoFix(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface I4 {
    [x: string, y: number]: number;
}

class C implements I {[|  |]}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCodeFixNotAvailable(t)
}
