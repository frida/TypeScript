package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/lsp/lsproto"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestSignatureHelpInAdjacentBlockBody(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `declare function foo(...args);

foo(() => {/*1*/}/*2*/)`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "1")
	f.VerifySignatureHelpPresent(t, &lsproto.SignatureHelpContext{TriggerKind: lsproto.SignatureHelpTriggerKindInvoked})
	f.GoToMarker(t, "2")
	f.VerifySignatureHelpPresent(t, &lsproto.SignatureHelpContext{TriggerKind: lsproto.SignatureHelpTriggerKindInvoked})
}
