package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestGetEditsForFileRename_symlink(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /foo.ts
// @Symlink: /node_modules/foo/index.ts
export const x = 0;
// @Filename: /user.ts
import { x } from 'foo';`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyNoErrors(t)
	f.VerifyWillRenameFilesEdits(t, "/user.ts", "/luser.ts", map[string]string{}, nil /*preferences*/)
}
