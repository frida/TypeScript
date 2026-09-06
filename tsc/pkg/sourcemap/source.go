package sourcemap

import "github.com/frida/TypeScript/tsc/pkg/core"

type Source interface {
	Text() string
	FileName() string
	ECMALineMap() []core.TextPos
}
