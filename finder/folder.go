package finder

import (
	"fmt"
	"os"
)

const HIDDEN byte = '.'

// Folder represents a directory in the file system.
// It stores the directory's absolute path, its name,
// its content, and whether it has already been processed.
type Folder struct {
	absolutePath   string
	folderName     string
	folderContent  []os.DirEntry
	alreadyVisited bool
}

// CollectFiles traverses the folder's content and separates
// directories from regular files.
//
// Hidden entries, whose names start with '.', are ignored. Subdirectories
// are added to the provided stack for later traversal, while regular files
// are appended to the files slice as absolute paths.
func (fd *Folder) CollectFiles(folders *Stack, files *[]string, perfTracker *PerformanceTracker) {
	for _, f := range fd.folderContent {

		if f.Name()[0] == HIDDEN {
			perfTracker.AddHiddenFile()
			continue
		}

		absolutePath := fmt.Sprintf("%v/%v", fd.absolutePath, f.Name())

		if f.IsDir() {
			fd.alreadyVisited = true
			c, err := os.ReadDir(absolutePath)
			perfTracker.AddFolder()

			if err != nil {
				fmt.Printf("[ERROR]: Error reading path '%s'. Following error occured: %s", absolutePath, err)
				continue
			}

			folders.Push(Folder{absolutePath, f.Name(), c, false})

		} else {
			*files = append(*files, absolutePath)
			perfTracker.AddFile()
		}
	}
}
