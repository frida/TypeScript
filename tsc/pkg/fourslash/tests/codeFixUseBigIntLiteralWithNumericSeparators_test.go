package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestCodeFixUseBigIntLiteralWithNumericSeparators(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `6_402_373_705_728_000;  // 18! < 2 ** 53
0x16_BE_EC_CA_73_00_00; // 18! < 2 ** 53`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCodeFixNotAvailable(t, "useBigintLiteral")
}
