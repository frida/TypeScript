package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestDocumentHighlights_33722(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /y.ts
class Foo {
  private foo() {}
}

const f = () => new Foo();
export default f;
// @Filename: /x.ts
import y from "./y";

y().[|foo|]();`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineDocumentHighlightsWithOptions(t, nil /*preferences*/, []string{"/x.ts"}, f.Ranges()[0])
}
