Combined files checksum tool that works on Termux, iSH, Windows, MacOS, and Linux (AMD64,ARM64,ARMv7)

Usage
```
./FolderSum <folder-path>

or 

Drag and drop a folder or file to the executable.
```
Example
```
...
Processed 98 of 100 files: FolderSum/README.md
Processed 99 of 100 files: FolderSum/go.mod
Processed 100 of 100 files: FolderSum/main.go
Checksum of combined file data:: c3b02d6fb72d73bd69943aaad86254bf0d4bf60cd752f867277bd6d0f31b4cee
```

Motive: Had corrupted photos during phone to desktop transfer, no phone app calculates a folder checksum.

Termux and iSH scripting was limiting and dependency heavy.

Note that unlike 7zip, the order of files (by alphabatical) or name does not affect the result.