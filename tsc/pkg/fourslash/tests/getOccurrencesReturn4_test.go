package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	. "github.com/frida/TypeScript/tsc/pkg/fourslash/tests/util"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestGetOccurrencesReturn4(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function f(a: number) {
    if (a > 0) {
        return (function () {
            return/*1*/;
            return/*2*/;
            return/*3*/;

            if (false) {
                return/*4*/ true;
            }
        })() || true;
    }

    var unusued = [1, 2, 3, 4].map(x => { return/*5*/ 4 })

    return/*6*/;
    return/*7*/ true;
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, ToAny(f.Markers())...)
}
