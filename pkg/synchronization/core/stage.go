package core

import (
	"bytes"
)

// stagingPathFinder recursively identifies paths/entries that need be staged in
// order to perform transitioning.
type stagingPathFinder struct {
	// paths is the list of paths for encountered file entries.
	paths []string
	// digests is the list of digests for encountered file entries, with length
	// and contents corresponding to paths.
	digests [][]byte
}

// find recursively searches for file entries that need staging.
func (f *stagingPathFinder) find(path string, entry *Entry) {
	if entry == nil {
		return
	} else if entry.Kind == EntryKind_Directory {
		// Compute the prefix to add to content names to compute their paths.
		var contentPathPrefix string
		if len(entry.Contents) > 0 {
			contentPathPrefix = pathJoinable(path)
		}

		// Process contents.
		for name, entry := range entry.Contents {
			f.find(contentPathPrefix+name, entry)
		}
	} else if entry.Kind == EntryKind_File {
		f.paths = append(f.paths, path)
		f.digests = append(f.digests, entry.Digest)
	}
}

// TransitionDependencies analyzes a list of transitions and determines the file
// paths (and their corresponding digests) that will need to be provided in
// order to apply the transitions using Transition. It will return these paths
// in depth-first traversal order.
func TransitionDependencies(transitions []*Change) ([]string, [][]byte) {
	// Create a path finder.
	finder := &stagingPathFinder{}

	// Have it find paths for all the transitions.
	for _, t := range transitions {
		// If this is a file-to-file transition and only the executability bit
		// is changing, then we don't need to stage, because Transition will
		// just modify the target on disk.
		fileToFileSameContents := t.Old != nil && t.New != nil &&
			t.Old.Kind == EntryKind_File && t.New.Kind == EntryKind_File &&
			bytes.Equal(t.Old.Digest, t.New.Digest)
		if fileToFileSameContents {
			continue
		}

		// Otherwise we need to perform a full scan.
		finder.find(t.Path, t.New)
	}

	// Success.
	return finder.paths, finder.digests
}
