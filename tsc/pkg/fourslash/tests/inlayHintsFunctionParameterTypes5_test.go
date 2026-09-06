package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/core"
	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/ls/lsutil"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestInlayHintsFunctionParameterTypes5(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `declare const STATE_SIGNAL: unique symbol;

declare function test(
  cb: (state: { [STATE_SIGNAL]: unknown }) => void,
): unknown;

test((state) => {});`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineInlayHints(t, nil /*span*/, &lsutil.UserPreferences{InlayHints: lsutil.InlayHintsPreferences{IncludeInlayFunctionParameterTypeHints: core.TSTrue}})
}
