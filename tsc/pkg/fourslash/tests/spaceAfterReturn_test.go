package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestSpaceAfterReturn(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function f( ) {
return       1;/*1*/
return[1];/*2*/
return    ;/*3*/
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.FormatDocument(t, "")
	f.GoToMarker(t, "1")
	f.VerifyCurrentLineContent(t, `    return 1;`)
	f.GoToMarker(t, "2")
	f.VerifyCurrentLineContent(t, `    return [1];`)
	f.GoToMarker(t, "3")
	f.VerifyCurrentLineContent(t, `    return;`)
}
