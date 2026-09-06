package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	. "github.com/frida/TypeScript/tsc/pkg/fourslash/tests/util"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestCompletionsInEmptyTupleType(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `type UserTuple = [["name", string], ["age", number], ["address", string]];
type AdminTuple = [/*1*/];
type OtherTuple = [string, /*2*/];
type QueryTuple = [typeof /*3*/];

const User: UserTuple = [["name", "2333"], ["age", 2333], ["address", "2333"]];`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, []string{"1", "2"}, &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				"UserTuple",
			},
			Excludes: []string{
				"User",
			},
		},
	})

	// After `typeof` in a tuple type we are back in a value location, so type-only symbols shouldn't be offered.
	f.VerifyCompletions(t, "3", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				"User",
			},
			Excludes: []string{
				"UserTuple",
			},
		},
	})
}
