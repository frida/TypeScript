package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	. "github.com/frida/TypeScript/tsc/pkg/fourslash/tests/util"
	"github.com/frida/TypeScript/tsc/pkg/ls"
	"github.com/frida/TypeScript/tsc/pkg/lsp/lsproto"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestCompletionsImport_filteredByPackageJson_typesOnly(t *testing.T) {
	t.Skip("Known failing fourslash test")
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `//@noEmit: true
//@Filename: /package.json
{
  "devDependencies": {
    "@types/react": "*"
  }
}
//@Filename: /node_modules/@types/react/index.d.ts
export declare var React: any;
//@Filename: /node_modules/@types/react/package.json
{
  "name": "@types/react"
}
//@Filename: /node_modules/@types/fake-react/index.d.ts
export declare var ReactFake: any;
//@Filename: /node_modules/@types/fake-react/package.json
{
  "name": "@types/fake-react"
}
//@Filename: /src/index.ts
const x = Re/**/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:               "React",
					AdditionalTextEdits: fourslash.AnyTextEdits,
					Data: &lsproto.CompletionItemData{
						AutoImport: &lsproto.AutoImportFix{
							ModuleSpecifier: "react",
						},
					},
					SortText: new(string(ls.SortTextAutoImportSuggestions)),
				},
			},
			Excludes: []string{
				"ReactFake",
			},
		},
	})
}
