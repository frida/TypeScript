package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestCodeFixClassImplementInterfaceMemberNestedTypeAlias(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `type Either<T> = { val: T } | Error;
interface I {
    x: Either<Either<string>>;
    foo(x: Either<Either<string>>): void;
}
class C implements I {}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCodeFix(t, fourslash.VerifyCodeFixOptions{
		Description: "Implement interface 'I'",
		NewFileContent: `type Either<T> = { val: T } | Error;
interface I {
    x: Either<Either<string>>;
    foo(x: Either<Either<string>>): void;
}
class C implements I {
    x: Either<Either<string>>;
    foo(x: Either<Either<string>>): void {
        throw new Error("Method not implemented.");
    }
}`,
		Index: 0,
	})
}
