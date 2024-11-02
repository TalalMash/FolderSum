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
	// Check if a folder path is provided as an argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: FolderSum(.exe) <folder-path>")
		return
	}

	folderPath := os.Args[1] // Get the folder path from the command line argument
	combinedHash := sha256.New()

	var totalFiles int
	var processedFiles int

	// First pass: count files
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

	// Second pass: calculate hashes
	err = filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Process only files (skip directories)
		if !info.IsDir() {
			// Read file content
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			// Compute SHA-256 hash for the file
			fileHash := sha256.Sum256(data)
			// Write the file hash to the combined hash
			combinedHash.Write(fileHash[:])

			// Update progress
			processedFiles++
			fmt.Printf("Processed %d of %d files: %s\n", processedFiles, totalFiles, path)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	// Final combined checksum
	finalHash := combinedHash.Sum(nil)
	fmt.Printf("Combined SHA-256 checksum: %s\n", hex.EncodeToString(finalHash))
}
