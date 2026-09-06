package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	. "github.com/frida/TypeScript/tsc/pkg/fourslash/tests/util"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestGetOccurrencesStringLiterals(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `var x = "[|string|]";
function f(a = "[|initial value|]") { }`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, ToAny(f.Ranges())...)
}
