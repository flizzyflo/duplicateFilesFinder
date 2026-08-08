package finder

// FilesPerHash stores all files that share the same content hash.
// It keeps track of the associated file paths, the hash value,
// the number of files, and whether duplicates exist.
type FilesPerHash struct {
	filePaths                []string
	fileContentHash          string
	fileCount                int
	fileContentHasDuplicates bool
}

// Append adds a file path to the hash group and updates the
// corresponding metadata, including the file count and duplicate flag.
func (hf *FilesPerHash) Append(hash string, filepath string) {
	hf.filePaths = append(hf.filePaths, filepath)
	hf.fileCount++
	hf.fileContentHasDuplicates = len(hf.filePaths) > NO_DUPLICATES
	hf.fileContentHash = hash

}
