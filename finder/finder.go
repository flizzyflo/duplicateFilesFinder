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
	enableLogs           bool
	pretty               bool
}

// TODO: Add docstrings
// TODO: Add errorhandling
// TODO: Add logging
// TODO: Add performance measurement
func (ff *FileFinder) Initialize(startfolder string, enableLogs bool, enablePerfMeasurement bool, pretty bool) {

	ff.enableLogs = enableLogs
	ff.pretty = pretty
	if ff.enableLogs {
		log.Printf("[LOG]\tInitializing Finder.")
	}
	ff.startFolder = startfolder
	content, _ := os.ReadDir(ff.startFolder)

	f := Folder{ff.startFolder, "", content, false}
	ff.filesToBeHashed = []string{}
	ff.folderStack = Stack{}
	ff.duplicateFilesResult = Result{}
	ff.duplicateFilesResult.init()
	ff.folderStack.Push(f)

	if ff.enableLogs {
		log.Printf("[LOG]\tInitialization of Finder successful.")
	}

}

// TODO: Add proper error handling
func (ff *FileFinder) collectAllFiles() {
	if ff.enableLogs {
		log.Printf("[LOG]\tStart collecting relevant files.")
	}

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

	if ff.enableLogs {
		log.Printf("[LOG]\tFinished collecting relevant files.")
	}

}

func createResult() {

}

func (ff *FileFinder) FindDuplicates() {

	ff.collectAllFiles()
	if ff.enableLogs {
		log.Printf("[LOG]\tStart looking for duplicates.")
	}

	// TODO: Add Result struct for easier analysis
	for _, relevantFile := range ff.filesToBeHashed {
		fileContent, _ := os.ReadFile(relevantFile)
		fileHash := fmt.Sprintf("%x", sha256.Sum256(fileContent))

		if fileHash == EMPTY_FILE_HASH {
			continue
		}

		ff.duplicateFilesResult.Add(fileHash, relevantFile)

	}
	if ff.enableLogs {
		log.Printf("[LOG]\tSuccessfully collected duplicates.")
	}
}

func (ff *FileFinder) PrintResults() {
	if ff.pretty {

		ff.PrettyPrintResult()

	} else {
		fmt.Println(ff.GetDuplicates())

	}
}

func (ff *FileFinder) PrettifyResult() {
	if len(ff.filesToBeHashed) == 0 {
		fmt.Printf("[WARNING]: No files are hashed, thus no result.")
		return
	}

	ff.PrettyPrintResult()
	ff.information()
}

func (ff *FileFinder) GetDuplicates() [][]string {

	var duplicateContent [][]string

	for _, file := range ff.duplicateFilesResult.FilesPerHash {

		if file.fileContentHasDuplicates {
			duplicateContent = append(duplicateContent, file.filePaths)
		}
	}
	return duplicateContent
}

func (ff *FileFinder) PrettyPrintResult() {

	f := ff.GetDuplicates()

	for _, dups := range f {

		var resultString string

		for _, filePath := range dups {
			resultString = resultString + fmt.Sprintf("\t\t%v\n", filePath)
		}
		fmt.Printf("-----------------------------------------------------------\n")
		fmt.Printf("[RESULT] -> \tFiles listed below have the same content\n%v\n", resultString)
	}
}

func (ff *FileFinder) information() {
	fmt.Printf("-----------------------------------------------------------\n")
	fmt.Printf("[INFORMATION] Files equality is considered based on hashed file content.\n")
	fmt.Printf("[INFORMATION] Path of files with same hash ( thus same content ) is returned as single result.\n")
	fmt.Printf("-----------------------------------------------------------\n")
}
