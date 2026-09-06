package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestCodeFixClassImplementInterfaceMultipleImplements2(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @strict: false
interface I1 {
    x: number;
}
interface I2 {
    y: "𣋝ઢȴ¬⏊";
}

class C implements I1,I2 {[|
    |]x: number;
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyRangeAfterCodeFix(t, `
y: "𣋝ઢȴ¬⏊";
`, false, 0, 0)
	f.VerifyCodeFixNotAvailable(t)
}
