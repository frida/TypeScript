package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestImportNameCodeFixNewImportIndex(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a/index.ts
export const foo = 0;
// @Filename: /b.ts
[|/**/foo;|]`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToFile(t, "/a/index.ts")
	f.GoToFile(t, "/b.ts")
	f.VerifyImportFixAtPosition(t, []string{
		`import { foo } from "./a";

foo;`,
	}, nil /*preferences*/)
}
