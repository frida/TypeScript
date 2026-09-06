package testrunner

import (
	"testing"

	"github.com/frida/TypeScript/tsc/pkg/bundled"
	"github.com/frida/TypeScript/tsc/pkg/collections"
	"github.com/frida/TypeScript/tsc/pkg/tspath"
	"gotest.tools/v3/assert"
)

func TestLocal(t *testing.T) { runCompilerTests(t) } //nolint:paralleltest

func runCompilerTests(t *testing.T) {
	t.Parallel()
	if !bundled.Embedded {
		t.Skip("bundled files are not embedded")
	}
	runners := []*CompilerBaselineRunner{
		NewCompilerBaselineRunner(TestTypeRegression),
		NewCompilerBaselineRunner(TestTypeConformance),
	}
	var seenTests collections.Set[string]
	for _, runner := range runners {
		for _, test := range runner.EnumerateTestFiles() {
			test = tspath.GetBaseFileName(test)
			assert.Assert(t, !seenTests.Has(test), "Duplicate test file: %s", test)
			seenTests.Add(test)
		}
	}
	for _, runner := range runners {
		runner.RunTests(t)
	}
}
