package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestCodeFixAddMissingAwait_topLevel(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `declare function getPromise(): Promise<string>;
const p = getPromise();
while (true) {
  p/*0*/.toLowerCase();
  getPromise()/*1*/.toLowerCase();
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCodeFixNotAvailable(t, "addMissingAwait")
	f.VerifyCodeFixNotAvailable(t, "addMissingAwaitToInitializer")
}
