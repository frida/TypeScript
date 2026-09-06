package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/core"
	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestFormatNestedClassWithOpenBraceOnNewLines(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `module A
{
    class B {
        /*1*/
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	opts168 := f.GetOptions()
	opts168.FormatCodeSettings.PlaceOpenBraceOnNewLineForControlBlocks = core.TSTrue
	f.Configure(t, opts168)
	opts232 := f.GetOptions()
	opts232.FormatCodeSettings.PlaceOpenBraceOnNewLineForFunctions = core.TSTrue
	f.Configure(t, opts232)
	f.GoToMarker(t, "1")
	f.Insert(t, "}")
	f.VerifyCurrentFileContent(t, `module A
{
    class B
    {
    }
}`)
}
