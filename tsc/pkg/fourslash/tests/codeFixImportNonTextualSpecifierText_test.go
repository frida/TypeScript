package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

// https://github.com/frida/TypeScript/tsc/issues/3944
func TestCodeFixImportNonTextualSpecifierText(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = "// @Filename: /a.ts\n" +
		"import type { A } from `./${myFolder}/${myFile}`;\n" +
		"\n" +
		"new A/**/()"
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "")
	f.VerifyImportFixAtPosition(t, []string{
		"import { A } from `./${myFolder}/${myFile}`;\n" +
			"\n" +
			"new A()",
	}, nil /*preferences*/)
}
