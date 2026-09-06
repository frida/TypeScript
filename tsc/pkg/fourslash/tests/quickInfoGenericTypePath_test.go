package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestQuickInfoGenericTypePath(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `
function f<T>(x: T) {
  class C {
    value = x
  }
  return new C()
}

class Box<T> {
  public value: T;
  constructor(value: T) {
    this.value = value;
  }
}

const instance = f/*callF*/("hello");
const b1/*b1*/ = new Box/*newBox*/(instance);
declare const b2/*b2*/: Box<typeof instance>;
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineHover(t)
}
