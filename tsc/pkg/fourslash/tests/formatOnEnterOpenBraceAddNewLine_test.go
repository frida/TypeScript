package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/core"
	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestFormatOnEnterOpenBraceAddNewLine(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `if(true) {/*0*/}
if(false)/*1*/{
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	opts148 := f.GetOptions()
	opts148.FormatCodeSettings.PlaceOpenBraceOnNewLineForControlBlocks = core.TSTrue
	f.Configure(t, opts148)
	f.GoToMarker(t, "0")
	f.InsertLine(t, "")
	f.VerifyCurrentFileContent(t, `if (true)
{
}
if(false){
}`)
	f.GoToMarker(t, "1")
	f.InsertLine(t, "")
	f.VerifyCurrentFileContent(t, `if (true)
{
}
if (false)
{
}`)
}
