package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestImportNameCodeFixReExport(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.ts
export const x = 0";
// @Filename: /b.ts
[|export { x } from "./a";
x;|]`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToFile(t, "/b.ts")
	f.VerifyRangeAfterCodeFix(t, `import { x } from "./a";

export { x } from "./a";
x;`, true, 0, 0)
}
