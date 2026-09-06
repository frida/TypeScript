package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/ls/lsutil"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestAutoImportIndexReExportPrefix(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @module: nodenext
// @Filename: /package.json
{ "type": "module" }
// @Filename: /utils/sum/index.ts
export { sum } from "./sum.js";
// @Filename: /utils/sum/sum.ts
export const sum = 0;
// @Filename: /utils/sumAB.ts
sum/**/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyImportFixModuleSpecifiers(t, "", []string{"./sum/index.js", "./sum/sum.js"}, &lsutil.UserPreferences{ImportModuleSpecifierEnding: "js"})
}
