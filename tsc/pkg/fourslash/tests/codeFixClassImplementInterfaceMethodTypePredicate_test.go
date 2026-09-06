package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestCodeFixClassImplementInterfaceMethodTypePredicate(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface I {
    f(i: any): i is I;
    f(): this is I;
}

class C implements I {}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCodeFix(t, fourslash.VerifyCodeFixOptions{
		Description: "Implement interface 'I'",
		NewFileContent: `interface I {
    f(i: any): i is I;
    f(): this is I;
}

class C implements I {
    f(i: any): i is I;
    f(): this is I;
    f(i?: unknown): boolean {
        throw new Error("Method not implemented.");
    }
}`,
		Index: 0,
	})
}
