package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestCodeFixClassImplementInterfaceObjectLiteral(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface IPerson {
    coordinate: {
        x: number;
        y: number;
    }
}
class Person implements IPerson { }`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCodeFix(t, fourslash.VerifyCodeFixOptions{
		Description: "Implement interface 'IPerson'",
		NewFileContent: `interface IPerson {
    coordinate: {
        x: number;
        y: number;
    }
}
class Person implements IPerson {
    coordinate: { x: number; y: number; };
}`,
		Index: 0,
	})
}
