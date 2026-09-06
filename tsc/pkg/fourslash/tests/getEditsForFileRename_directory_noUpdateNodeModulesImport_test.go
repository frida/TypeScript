package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestGetEditsForFileRename_directory_noUpdateNodeModulesImport(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a/b/file1.ts
import { foo } from "foo";
// @Filename: /a/b/node_modules/foo/index.d.ts
export const foo = 0;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyWillRenameFilesEdits(t, "/a/b", "/a/d", map[string]string{}, nil /*preferences*/)
}
