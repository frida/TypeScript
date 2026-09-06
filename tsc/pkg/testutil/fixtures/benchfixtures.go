package fixtures

import (
	"path/filepath"

	"github.com/frida/TypeScript/tsc/pkg/repo"
	"github.com/frida/TypeScript/tsc/pkg/testutil/filefixture"
)

var BenchFixtures = []filefixture.Fixture{
	filefixture.FromString("empty.ts", "empty.ts", ""),
	filefixture.FromFile("checker.ts", filepath.Join(repo.TestDataPath(), "fixtures/compiler/checker.ts")),
	filefixture.FromFile("dom.generated.d.ts", filepath.Join(repo.TestDataPath(), "fixtures/lib/dom.generated.d.ts")),
	filefixture.FromFile("Herebyfile.mjs", filepath.Join(repo.TestDataPath(), "fixtures/typescript/Herebyfile.mjs")),
	filefixture.FromFile("jsxComplexSignatureHasApplicabilityError.tsx", filepath.Join(repo.TestDataPath(), "fixtures/testcases/compiler/jsxComplexSignatureHasApplicabilityError.tsx")),
}
