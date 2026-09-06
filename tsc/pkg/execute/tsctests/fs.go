package tsctests

import (
	"fmt"

	"github.com/frida/TypeScript/tsc/pkg/collections"
	"github.com/frida/TypeScript/tsc/pkg/core"
	"github.com/frida/TypeScript/tsc/pkg/execute/incremental"
	"github.com/frida/TypeScript/tsc/pkg/json"
	"github.com/frida/TypeScript/tsc/pkg/testutil/fsbaselineutil"
	"github.com/frida/TypeScript/tsc/pkg/testutil/harnessutil"
	"github.com/frida/TypeScript/tsc/pkg/tspath"
	"github.com/frida/TypeScript/tsc/pkg/vfs"
)

type testFs struct {
	vfs.FS
	defaultLibs  *collections.SyncSet[string]
	writtenFiles collections.SyncSet[string]
}

func (f *testFs) removeIgnoreLibPath(path string) {
	if f.defaultLibs != nil && f.defaultLibs.Has(path) {
		f.defaultLibs.Delete(path)
	}
}

// ReadFile reads the file specified by path and returns the content.
// If the file fails to be read, ok will be false.
func (f *testFs) ReadFile(path string) (contents string, ok bool) {
	f.removeIgnoreLibPath(path)
	return f.readFileHandlingBuildInfo(path)
}

func (f *testFs) readFileHandlingBuildInfo(path string) (contents string, ok bool) {
	contents, ok = f.FS.ReadFile(path)
	if ok && tspath.FileExtensionIs(path, tspath.ExtensionTsBuildInfo) {
		// read buildinfo and modify version
		var buildInfo incremental.BuildInfo
		err := json.Unmarshal([]byte(contents), &buildInfo)
		if err == nil && buildInfo.Version == harnessutil.FakeTSVersion {
			buildInfo.Version = core.Version()
			newContents, err := json.Marshal(&buildInfo)
			if err != nil {
				panic("testFs.ReadFile: failed to marshal build info after fixing version: " + err.Error())
			}
			contents = string(newContents)
		}
	}
	return contents, ok
}

func (f *testFs) WriteFile(path string, data string) error {
	f.removeIgnoreLibPath(path)
	f.writtenFiles.Add(path)
	return f.writeFileHandlingBuildInfo(path, data)
}

func (f *testFs) writeFileHandlingBuildInfo(path string, data string) error {
	if tspath.FileExtensionIs(path, tspath.ExtensionTsBuildInfo) {
		var buildInfo incremental.BuildInfo
		if err := json.Unmarshal([]byte(data), &buildInfo); err == nil {
			if buildInfo.Version == core.Version() {
				// Change it to harnessutil.FakeTSVersion
				buildInfo.Version = harnessutil.FakeTSVersion
				newData, err := json.Marshal(&buildInfo)
				if err != nil {
					return fmt.Errorf("testFs.WriteFile: failed to marshal build info after fixing version: %w", err)
				}
				data = string(newData)
			}
			// Write readable build info version
			if err := f.WriteFile(
				path+".readable.baseline.txt",
				toReadableBuildInfo(&buildInfo, fsbaselineutil.SanitizeInternalSymbolName(data)),
			); err != nil {
				return fmt.Errorf("testFs.WriteFile: failed to write readable build info: %w", err)
			}
		} else {
			panic("testFs.WriteFile: failed to unmarshal build info: - use underlying FS's write method if this is intended use for testcase" + err.Error())
		}
	}
	return f.FS.WriteFile(path, data)
}

// Removes `path` and all its contents. Will return the first error it encounters.
func (f *testFs) Remove(path string) error {
	f.removeIgnoreLibPath(path)
	return f.FS.Remove(path)
}
