package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: FolderSum <folder-path>")
		return
	}

	folderPath := os.Args[1]
	combinedHash := sha256.New()

	var totalFiles int
	var processedFiles int

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			totalFiles++
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	err = filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			fileHash := sha256.Sum256(data)
			combinedHash.Write(fileHash[:])
			processedFiles++
			fmt.Printf("Processed %d of %d files: %s\n", processedFiles, totalFiles, path)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	finalHash := combinedHash.Sum(nil)
	fmt.Printf("Checksum of all file checksums: %s\n", hex.EncodeToString(finalHash))
}
