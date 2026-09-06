package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestNavigationBarComputedPropertyName(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function F(key, value) {
    return {
        [key]: value,
        "prop": true
    }
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineDocumentSymbol(t)
}
