package fourslash_test

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/fourslash"
	"github.com/frida/TypeScript/tsc/pkg/testutil"
)

func TestIncrementalJsDocAdjustsLengthsRight(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @noLib: true

/**
 * Pad ` + "`" + `str` + "`" + ` to ` + "`" + `width` + "`" + `.
 *
 * @param {String} str
 * @param {Number} wid/*1*/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "1")
	f.Insert(t, "th\n@")
}
