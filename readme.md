# Duplicate File Finder

This program detects duplicate files within your file system. It first scans the filesystem starting from provided root path on downwards for files. Second, it creates hashes for each individual file contents. If it finds more than one file per content-hash, it found a duplicate. This will be the result.

# Example Result
> -----------------------------------------------------------
>[RESULT] -> 	Files listed below have the same content
>		/Users/florianluebke/Desktop/stuff//lexer-cpp/flex-lexer/cmake-build-debug/CMakeFiles/clion-environment.txt
>		/Users/florianluebke/Desktop/stuff//Lexer cpp/cmake-build-debug/CMakeFiles/clion-environment.txt
>
>-----------------------------------------------------------
>[RESULT] -> 	Files listed below have the same content
>		/Users/florianluebke/Desktop/stuff//lexer-cpp/flex-lexer/cmake-build-debug/CMakeFiles/cmake.check_cache
>		/Users/florianluebke/Desktop/stuff//Lexer cpp/cmake-build-debug/CMakeFiles/cmake.check_cache
>		...


[RESULT] -> 	Files listed below have the same content
		/Users/florianluebke/Desktop/stuff//lexer-cpp/flex-lexer/cmake-build-debug/CMakeFiles/3.29.6/CompilerIdC/CMakeCCompilerId.c
		/Users/florianluebke/Desktop/stuff//Lexer cpp/cmake-build-debug/CMakeFiles/3.29.6/CompilerIdC/CMakeCCompilerId.c

-----------------------------------------------------------
[RESULT] -> 	Files listed below have the same content
		/Users/florianluebke/Desktop/stuff//lexer-cpp/flex-lexer/cmake-build-debug/CMakeFiles/3.29.6/CompilerIdCXX/CMakeCXXCompilerId.cpp
		/Users/florianluebke/Desktop/stuff//Lexer cpp/cmake-build-debug/CMakeFiles/3.29.6/CompilerIdCXX/CMakeCXXCompilerId.cpp

-----------------------------------------------------------
[RESULT] -> 	Files listed below have the same content
		/Users/florianluebke/Desktop/stuff//sortVisuals/sorting/heapsort.go
		/Users/florianluebke/Desktop/stuff//sortVisuals/sorting/quicksort.go´inline "test"´

## Dependencies
-Just uses go build-in features
*go version 1.26 recommended
+plus


