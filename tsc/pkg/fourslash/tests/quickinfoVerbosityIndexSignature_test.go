package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestQuickinfoVerbosityIndexSignature(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `type Key = string | number;
interface Apple {
    banana: number;
}
interface Foo {
    [a/*a*/: Key]: Apple;
}
const f/*f*/: Foo = {};`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineHoverWithVerbosity(t, map[string][]int{"a": {0, 1}, "f": {0, 1, 2}})
}
