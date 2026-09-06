package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestCodeFixClassImplementInterfaceTypeParamInstantiateT(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface I<T> { x: T; }
class C<T> implements I<T> {}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCodeFix(t, fourslash.VerifyCodeFixOptions{
		Description: "Implement interface 'I<T>'",
		NewFileContent: `interface I<T> { x: T; }
class C<T> implements I<T> {
    x: T;
}`,
		Index: 0,
	})
}
