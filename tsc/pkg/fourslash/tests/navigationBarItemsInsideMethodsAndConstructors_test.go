package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestNavigationBarItemsInsideMethodsAndConstructors(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class Class {
    constructor() {
        function LocalFunctionInConstructor() {}
        interface LocalInterfaceInConstrcutor {}
        enum LocalEnumInConstructor { LocalEnumMemberInConstructor }
    }

    method() {
        function LocalFunctionInMethod() {
            function LocalFunctionInLocalFunctionInMethod() {}
        }
        interface LocalInterfaceInMethod {}
        enum LocalEnumInMethod { LocalEnumMemberInMethod }
    }

    emptyMethod() { } // Non child functions method should not be duplicated
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineDocumentSymbol(t)
}
