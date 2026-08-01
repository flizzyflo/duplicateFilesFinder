package finder

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

const EMPTY_FILE_HASH string = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"

type FileFinder struct {
	startFolder          string
	filesToBeHashed      []string
	folderStack          Stack
	duplicateFilesResult Result
}

// TOD0: Add docstrings
// TOD0: Add errorhandling
// TOD0: Add logging
func (ff *FileFinder) Initialize(startfolder string) {

	log.Printf("[LOG]\tInitializing Finder.")
	ff.startFolder = startfolder
	content, _ := os.ReadDir(ff.startFolder)

	f := Folder{ff.startFolder, "", content, false}
	ff.filesToBeHashed = []string{}
	ff.folderStack = Stack{}
	ff.duplicateFilesResult = Result{}
	ff.duplicateFilesResult.init()
	ff.folderStack.Push(f)
	log.Printf("[LOG]\tInitialization of Finder successful.")

}

// TODO: Add proper error handling
func (ff *FileFinder) collectAllFiles() {

	log.Printf("[LOG]\tStart collecting relevant files.")

	for !ff.folderStack.IsEmpty() {
		currentFolder, err := ff.folderStack.Pop()

		if currentFolder.alreadyVisited {
			continue
		}

		if err != nil {
			fmt.Println(err)
		}

		currentFolder.extractFilesAndFolders(&ff.folderStack, &ff.filesToBeHashed)
	}
	log.Printf("[LOG]\tFinished collecting relevant files.")

}

func createResult() {

}

func (ff *FileFinder) FindDuplicates() [][]string {

	ff.collectAllFiles()
	log.Printf("[LOG]\tStart looking for duplicates.")

	// TODO: Add Result struct for easier analysis
	for _, relevantFile := range ff.filesToBeHashed {
		fileContent, _ := os.ReadFile(relevantFile)
		fileHash := fmt.Sprintf("%x", sha256.Sum256(fileContent))

		if fileHash == EMPTY_FILE_HASH {
			continue
		}

		ff.duplicateFilesResult.Add(fileHash, relevantFile)

	}

	log.Printf("[LOG]\tSuccessfully collected duplicates.")

	return ff.duplicateFilesResult.GetDuplicates()
}

func (ff *FileFinder) PrettifyResult() {
	if len(ff.filesToBeHashed) == 0 {
		fmt.Printf("[WARNING]: No files are hashed, thus no result.")
		return
	}

	ff.duplicateFilesResult.PrettyPrintResult()
	ff.information()
}

func (ff *FileFinder) information() {
	fmt.Printf("-----------------------------------------------------------\n")
	fmt.Printf("[INFORMATION] Files equality is considered based on hashed file content.\n")
	fmt.Printf("[INFORMATION] Path of files with same hash ( thus same content ) is returned as single result.\n")
	fmt.Printf("-----------------------------------------------------------\n")
}
