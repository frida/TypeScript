package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/core"
	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestFormatIfTryCatchBlocks(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `try {
}
catch {
}

try {
}
catch (e) {
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	opts187 := f.GetOptions()
	opts187.FormatCodeSettings.PlaceOpenBraceOnNewLineForControlBlocks = core.TSTrue
	f.Configure(t, opts187)
	f.FormatDocument(t, "")
	f.VerifyCurrentFileContent(t, `try
{
}
catch
{
}

try
{
}
catch (e)
{
}`)
}
