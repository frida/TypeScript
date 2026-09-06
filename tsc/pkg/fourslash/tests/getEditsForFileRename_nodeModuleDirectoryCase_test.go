package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestGetEditsForFileRename_nodeModuleDirectoryCase(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a/b/file1.ts
import { foo } from "foo";
// @Filename: /a/node_modules/foo/index.d.ts
export const foo = 0;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyWillRenameFilesEdits(t, "/a/b", "/a/B", map[string]string{}, nil /*preferences*/)
}
