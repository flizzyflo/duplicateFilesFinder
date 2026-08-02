# Duplicate File Finder

A lightweight command-line tool written in Go that detects duplicate files by comparing their **content**, not their filenames.

The program does not use recursion, it scans a directory tree stack based, computes a hash for every discovered file, and groups files that share the same hash. Files with identical hashes are reported as duplicates.

---

## How it works

The program performs the following steps:

1. **stack-based file system scan**
   - Traverses the directory tree starting from the provided root directory.
   - Collects all regular files.

2. **Content hashing**
   - Computes a hash for every discovered file **content**. Uses SHA256.
   - Uses the **content-hash** as a unique identifier for the file contents.

3. **Duplicate detection**
   - Groups files by their content hash.
   - Reports every group containing more than one file.

---

## Example Output

### Raw Output

```text
[
    [
        /user/user/someDir/file1.txt,
        /user/user/anotherDir/file1_copy.txt
    ],
    ...
]
```

### Prettified Output

```
-----------------------------------------------------------
[RESULT] Files listed below have identical content

    /Users/florianluebke/Desktop/stuff/lexer-cpp/flex-lexer/cmake-build-debug/CMakeFiles/clion-environment.txt
    /Users/florianluebke/Desktop/stuff/Lexer cpp/cmake-build-debug/CMakeFiles/clion-environment.txt

-----------------------------------------------------------
[RESULT] Files listed below have identical content

    /Users/florianluebke/Desktop/stuff/lexer-cpp/flex-lexer/cmake-build-debug/CMakeFiles/cmake.check_cache
    /Users/florianluebke/Desktop/stuff/Lexer cpp/cmake-build-debug/CMakeFiles/cmake.check_cache
```

---

## Features

- Stack-based directory traversal
- Content-based duplicate detection
- Hash-based comparison
- Lightweight implementation
- No external dependencies
- Cross-platform (Linux, macOS, Windows)

---

## Requirements

- Go **1.26** (recommended)

---

## Dependencies

This project uses **only the Go standard library**.

No third-party packages are required.

---

## Build

```bash
go build
```

or

```bash
go build -o duplicate-file-finder
```

---

## Run

```bash
./duplicate-file-finder <root-directory>
```

Example:

```bash
./duplicate-file-finder ~/Documents
```

---

```bash
Example usage:\tmain.go [args, optional] [filepath to start from]")
'-help' for help overview;                          main.go -help")
'-log' for enabling detailed logging;               main.go -log")
'-perf' for enabling performance measurement;       main.go -perf")
'-pretty' for enabling pretty output of results;    main.go -pretty")
'-info' for enabling information about collection;  main.go -info")
```
---

## License

This project is released under the MIT License.
