# Duplicate File Finder

This program detects duplicate files within your file system. It first scans the filesystem starting from provided root path on downwards for files. Second, it creates hashes for each individual file contents. If it finds more than one file per content-hash, it found a duplicate. This will be the result.

# Example Result
inline "test"

## Dependencies
-Just uses go build-in features
-go version 1.26 recommended



