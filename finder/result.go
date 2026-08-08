package finder

// Result represents the complete duplicate detection result.
// It maps a file content hash to the corresponding collection of files
// sharing that hash.
type Result struct {
	// key is hash
	FilesPerHash map[string]*FilesPerHash
}

// new initializes the internal hash map used to store grouped files.
// It must be called before adding any entries.
func (r *Result) new() {
	r.FilesPerHash = map[string]*FilesPerHash{}
}

// Add adds a file to the result under the specified content hash.
// If the hash already exists, the file is appended to the existing group.
// Otherwise, a new group is created for the hash.
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
