package hyphae

import (
	"path/filepath"
	"sync"

	"github.com/bouncepaw/mycorrhiza/internal/files"
)

// HTMLHypha is a hypha whose content is raw HTML, served as-is without mycomarkup processing.
type HTMLHypha struct {
	sync.RWMutex

	canonicalName string
	htmlFilePath  string
}

func (h *HTMLHypha) CanonicalName() string {
	return h.canonicalName
}

// HasTextFile reports false: an HTML hypha has no mycomarkup text file.
// This keeps it out of mycomarkup-driven iterations (backlinks, migrations).
func (h *HTMLHypha) HasTextFile() bool {
	return false
}

// TextFilePath returns the hypothetical path of a mycomarkup file. There is no
// such file on disk; the method exists to satisfy the ExistingHypha interface
// and to give git fallback paths something to chew on.
func (h *HTMLHypha) TextFilePath() string {
	return filepath.Join(files.HyphaeDir(), h.canonicalName+".myco")
}

// HTMLFilePath returns the on-disk path of the .html file backing this hypha.
func (h *HTMLHypha) HTMLFilePath() string {
	return h.htmlFilePath
}
