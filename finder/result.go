package finder

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
