package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestGetEditsForFileRename_js_simple(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @Filename: /a.js
import b from "./b.js";
// @Filename: /b.js
module.exports = 1;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyWillRenameFilesEdits(t, "/b.js", "/c.js", map[string]string{
		"/a.js": `import b from "./c.js";`,
	}, nil /*preferences*/)
}
