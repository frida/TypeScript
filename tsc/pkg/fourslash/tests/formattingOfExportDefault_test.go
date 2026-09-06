package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestFormattingOfExportDefault(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `namespace Foo {
/*1*/    export        default        class        Test { }
}
/*2*/export        default        function        bar() { }`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.FormatDocument(t, "")
	f.GoToMarker(t, "1")
	f.VerifyCurrentLineContent(t, `    export default class Test { }`)
	f.GoToMarker(t, "2")
	f.VerifyCurrentLineContent(t, `export default function bar() { }`)
}
