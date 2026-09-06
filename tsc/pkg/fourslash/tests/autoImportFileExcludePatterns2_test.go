package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	. "github.com/frida/TypeScript/tsc/pkg/fourslash/tests/util"
	"github.com/frida/TypeScript/tsc/pkg/ls"
	"github.com/frida/TypeScript/tsc/pkg/ls/lsutil"
	"github.com/frida/TypeScript/tsc/pkg/lsp/lsproto"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestAutoImportFileExcludePatterns2(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @lib: es5
// @Filename: /lib/components/button/Button.ts
export function Button() {}
// @Filename: /lib/components/button/index.ts
export * from "./Button";
// @Filename: /lib/components/index.ts
export * from "./button";
// @Filename: /lib/main.ts
export { Button } from "./components";
// @Filename: /lib/index.ts
export * from "./main";
// @Filename: /i-hate-index-files.ts
Button/**/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: CompletionGlobalsPlus(
				[]fourslash.CompletionsExpectedItem{
					&lsproto.CompletionItem{
						Label: "Button",
						Data: &lsproto.CompletionItemData{
							AutoImport: &lsproto.AutoImportFix{
								ModuleSpecifier: "./lib/main",
							},
						},
						AdditionalTextEdits: fourslash.AnyTextEdits,
						SortText:            new(string(ls.SortTextAutoImportSuggestions)),
					},
				}, false,
			),
		},
		UserPreferences: &lsutil.UserPreferences{AutoImportFileExcludePatterns: []string{"/**/index.*"}},
	})
	f.VerifyImportFixModuleSpecifiers(t, "", []string{"./lib/main", "./lib/components/button/Button"}, &lsutil.UserPreferences{AutoImportFileExcludePatterns: []string{"/**/index.*"}})
}
