package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestLinkedEditingJsxTag11(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /customElements.tsx
const jsx = <fbt:enum knownProp="accepted"
    unknownProp="rejected">
</fbt:enum>;

const customElement = <custom-element></custom-element>;

const standardElement = 
   <Link href="/hello" passHref>
       <Button component="a">
           Next
       </Button>
   </Link>;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineLinkedEditing(t)
}
