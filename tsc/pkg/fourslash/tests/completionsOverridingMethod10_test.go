package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/core"
	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	. "github.com/frida/TypeScript/tsc/pkg/fourslash/tests/util"
	"github.com/frida/TypeScript/tsc/pkg/ls"
	"github.com/frida/TypeScript/tsc/pkg/ls/lsutil"
	"github.com/frida/TypeScript/tsc/pkg/lsp/lsproto"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestCompletionsOverridingMethod10(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: a.ts
// @newline: LF
interface Base {
    a: string;
    b(a: string): void;
    c(a: string): string;
    c(a: number): number;
}
class Sub implements Base {
   /*a*/
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "a", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:      "a",
					InsertText: new("a: string;"),
					FilterText: new("a"),
					SortText:   new(string(ls.SortTextLocationPriority)),
				},
				&lsproto.CompletionItem{
					Label:      "b",
					InsertText: new("b(a: string): void {\n}"),
					FilterText: new("b"),
					SortText:   new(string(ls.SortTextLocationPriority)),
				},
				&lsproto.CompletionItem{
					Label:      "c",
					InsertText: new("c(a: string): string;\nc(a: number): number;\nc(a: unknown): string | number {\n}"),
					FilterText: new("c"),
					SortText:   new(string(ls.SortTextLocationPriority)),
				},
			},
		},
		UserPreferences: &lsutil.UserPreferences{IncludeCompletionsWithClassMemberSnippets: core.TSTrue},
	})
}
