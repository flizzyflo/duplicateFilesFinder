package finder

import "fmt"

// TODO: Rename Strcuts and Vars
type FilesPerHash struct {
	filePaths                []string
	fileContentHash          string
	fileCount                int
	fileContentHasDuplicates bool
}

func (hf *FilesPerHash) Append(hash string, filepath string) {
	hf.filePaths = append(hf.filePaths, filepath)
	hf.fileCount++
	hf.fileContentHasDuplicates = len(hf.filePaths) > 1
	hf.fileContentHash = hash

}

type Result struct {
	// key is hash
	FilesPerHash map[string]*FilesPerHash
}

func (r *Result) init() {
	r.FilesPerHash = map[string]*FilesPerHash{}
}

func (r *Result) Add(hash string, filepath string) {

	files, ok := r.FilesPerHash[hash]

	if ok {
		files.Append(hash, filepath)
		r.FilesPerHash[hash] = files
	} else {
		files := FilesPerHash{}
		files.Append(hash, filepath)
		r.FilesPerHash[hash] = &files
	}
}

func (r *Result) GetDuplicates() [][]string {

	var duplicateContent [][]string

	for _, file := range r.FilesPerHash {

		if file.fileContentHasDuplicates {
			duplicateContent = append(duplicateContent, file.filePaths)
		}
	}
	return duplicateContent
}

func (r *Result) PrettyPrintResult() {

	f := r.GetDuplicates()

	for _, dups := range f {

		var resultString string

		for _, filePath := range dups {
			resultString = resultString + fmt.Sprintf("\t\t%v\n", filePath)
		}
		fmt.Printf("-----------------------------------------------------------\n")
		fmt.Printf("[RESULT] -> \tFiles listed below have the same content\n%v\n", resultString)

	}

}
