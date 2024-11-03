package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: FolderSum <folder-path>")
		return
	}

	folderPath := os.Args[1]

	var fileChecksums []string

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		fileChecksum, err := checksumFile(path)
		if err != nil {
			return err
		}

		fileChecksums = append(fileChecksums, fileChecksum)
		fmt.Printf("Processed file: %s, checksum: %s\n", path, fileChecksum)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	sort.Strings(fileChecksums)

	combinedChecksum := sha256.New()
	for _, chk := range fileChecksums {
		combinedChecksum.Write([]byte(chk))
	}

	finalHash := combinedChecksum.Sum(nil)
	fmt.Printf("Final checksum of combined file checksums: %s\n", hex.EncodeToString(finalHash))

	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}

func checksumFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
