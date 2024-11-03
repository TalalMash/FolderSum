Combined files checksum tool that works on Termux, iSH, Windows, MacOS, and Linux (AMD64,ARM64,ARMv7)

Usage
```
./FolderSum <folder-path>

or 

Drag and drop a folder or file at the executable.
```
Example
```
...
Processed file: FolderSum/go.mod, checksum: 78e8e762699f68deecef407da4bead3514306285de5b224a8557e9e2bfc8f034
Processed file: FolderSum/main.go, checksum: df556f2a24bd4dac0f86e209aa837841856d12044d38427c8d31f08feb9c8092
Final checksum of combined file checksums: 10758305062ea88ff1db0d1ebdea61f1e4c4932dbed9ef778ffb1d49c5eb200f
```

Motive: Had corrupted photos during phone to desktop transfer, no phone app calculates a folder checksum.

Termux and iSH scripting was limiting and dependency heavy.

Note that unlike 7zip, the order of files (by alphabatical) or name does not affect the result.