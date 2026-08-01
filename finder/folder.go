package finder

import (
	"fmt"
	"os"
)

const HIDDEN byte = '.'

// TOD0: Add docstrings
// TOD0: Add errorhandling
// TOD0: Add logging
type Folder struct {
	absolutePath   string
	folderName     string
	folderContent  []os.DirEntry
	alreadyVisited bool
}

func (fd *Folder) extractFilesAndFolders(folders *Stack, files *[]string) {
	for _, f := range fd.folderContent {

		if f.Name()[0] == HIDDEN {
			continue
		}

		absolutePath := fmt.Sprintf("%v/%v", fd.absolutePath, f.Name())

		// TODO: Error handling
		if f.IsDir() {
			fd.alreadyVisited = true
			c, _ := os.ReadDir(absolutePath)
			folders.Push(Folder{absolutePath, f.Name(), c, false})

		} else {
			*files = append(*files, absolutePath)
		}
	}
}
