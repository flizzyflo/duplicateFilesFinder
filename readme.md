# Duplicate File Finder

This program detects duplicate files within your file system. It first scans the filesystem starting from provided root path on downwards for files. Second, it creates hashes for each individual file contents. If it finds more than one file per content-hash, it found a duplicate. This will be the result.

# Example Result raw
>  \t\t [[/user/user/someDir/someDir//someDir/some.txt, /user/user/someDir/someDir/some.txt] ... ]

# Example Result Prettified
> -----------------------------------------------------------
>[RESULT] -> 	Files listed below have the same content\n
>\t\t		/Users/florianluebke/Desktop/stuff//lexer-cpp/flex-lexer/cmake-build-debug/CMakeFiles/clion-environment.txt\n
>\t\t		/Users/florianluebke/Desktop/stuff//Lexer cpp/cmake-build-debug/CMakeFiles/clion-environment.txt\n
>\n
>-----------------------------------------------------------\n
>[RESULT] -> 	Files listed below have the same content\n
>\t\t		/Users/florianluebke/Desktop/stuff//lexer-cpp/flex-lexer/cmake-build-debug/CMakeFiles/cmake.check_cache\n
>\t\t		/Users/florianluebke/Desktop/stuff//Lexer cpp/cmake-build-debug/CMakeFiles/cmake.check_cache\n
>\t\t		...


## Dependencies
-Just uses go build-in features
*go version 1.26 recommended
+plus
