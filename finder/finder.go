package finder

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

const EMPTY_FILE_HASH string = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"

type FileFinder struct {
	duplicateFilesResult Result
	folderStack          Stack
	performanceTracker   PerformanceTracker
	duplicates           [][]string
	filesToBeHashed      []string
	startFolder          string
	enableLogs           bool
	pretty               bool
	info                 bool
}

func (ff *FileFinder) Run() {
	ff.findDuplicates()
	ff.printResults()
}

// TODO: Add docstrings
// TODO: Add errorhandling
// TODO: Add performance measurement
func (ff *FileFinder) Initialize(startfolder string, enableLogs bool, enablePerfMeasurement bool, pretty bool, info bool) error {

	ff.enableLogs = enableLogs
	ff.pretty = pretty
	ff.info = info
	ff.startFolder = startfolder

	if ff.enableLogs {
		log.Printf("[LOG]\tInitializing 'Finder'. Startfolder is: '%v'.", ff.startFolder)
	}
	ff.performanceTracker = PerformanceTracker{}
	content, err := os.ReadDir(ff.startFolder)
	ff.performanceTracker.Start()

	if err != nil {
		if ff.enableLogs {
			log.Printf("[LOG]\tCan not open startfolder '%v'.", ff.startFolder)
		}
		return fmt.Errorf("[ERROR] Can not open startfolder. Please provide other folder. Err: %v", err)
	}

	f := Folder{ff.startFolder, "", content, false}
	ff.filesToBeHashed = []string{}
	ff.folderStack = Stack{}
	ff.performanceTracker.AddFolder()

	ff.duplicateFilesResult = Result{}
	ff.duplicateFilesResult.new()
	ff.folderStack.Push(f)

	if ff.enableLogs {
		log.Printf("[LOG]\tInitialization of 'Finder' successful.")
	}

	return nil

}

// TODO: Add proper error handling
func (ff *FileFinder) collectAllFiles() {
	if ff.enableLogs {
		log.Printf("[LOG]\tStart collecting relevant files...")
	}

	for !ff.folderStack.IsEmpty() {
		currentFolder, err := ff.folderStack.Pop()

		if currentFolder.alreadyVisited {
			continue
		}

		if err != nil {
			fmt.Println(err)
		}
		currentFolder.CollectFiles(&ff.folderStack, &ff.filesToBeHashed, &ff.performanceTracker)
	}

	if ff.enableLogs {
		log.Printf("[LOG]\tFinished collecting relevant files.")
	}

}

func (ff *FileFinder) findDuplicates() {

	ff.collectAllFiles()
	if ff.enableLogs {
		log.Printf("[LOG]\tStart looking for duplicates.")
	}

	for _, relevantFile := range ff.filesToBeHashed {
		fileContent, _ := os.ReadFile(relevantFile)
		fileHash := fmt.Sprintf("%x", sha256.Sum256(fileContent))

		if fileHash == EMPTY_FILE_HASH {
			continue
		}

		ff.duplicateFilesResult.Add(fileHash, relevantFile)
	}

	ff.collectDuplicates()

	if ff.enableLogs {
		log.Printf("[LOG]\tSuccessfully collected duplicates.")
	}
}

func (ff *FileFinder) collectDuplicates() {

	var duplicateContent [][]string

	for _, file := range ff.duplicateFilesResult.FilesPerHash {

		if file.fileContentHasDuplicates {
			duplicateContent = append(duplicateContent, file.filePaths)
		}
	}
	ff.duplicates = duplicateContent
}

func (ff *FileFinder) printResults() {
	if len(ff.filesToBeHashed) == 0 {
		fmt.Printf("[WARNING]: No files are hashed, thus no result.")
		return
	}

	if ff.pretty {
		ff.prettifyResults()
	} else {
		fmt.Println(ff.duplicates)
	}
	ff.printInformation()
}

func (ff *FileFinder) prettifyResults() {

	for _, dups := range ff.duplicates {

		var resultString string

		for _, filePath := range dups {
			resultString = resultString + fmt.Sprintf("\t\t%v\n", filePath)
		}
		fmt.Printf("-----------------------------------------------------------\n")
		fmt.Printf("[RESULT] -> \tFiles listed below contain the same information\n%v\n", resultString)
	}
}

func (ff *FileFinder) printInformation() {
	if !ff.info {
		return
	}
	fmt.Printf("-----------------------------------------------------------\n")
	fmt.Printf("[INFORMATION] File equality is considered based on hashed file content.\n")
	fmt.Printf("[INFORMATION] Path of files with same hash ( thus same content ) is returned as a grouped result block, containing single results.\n")
	fmt.Printf("-----------------------------------------------------------\n")
}
