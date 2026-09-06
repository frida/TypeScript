package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/ls/lsutil"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestCompletionsImport_quoteStyle(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @module: esnext
// @Filename: /a.ts
export const foo = 0;
// @Filename: /b.ts
fo/**/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "")
	f.VerifyApplyCodeActionFromCompletion(t, new(""), &fourslash.ApplyCodeActionFromCompletionOptions{
		Name:            "foo",
		Source:          "./a",
		Description:     "Add import from \"./a\"",
		UserPreferences: &lsutil.UserPreferences{QuotePreference: lsutil.QuotePreference("single")},
		NewFileContent: new(`import { foo } from './a';

fo`),
	})
}
