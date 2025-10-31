package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CopyConfigurations copies everything from {version}/configurations/ into {version}/{launcher}/
func CopyConfigurations(version string, launcher string) error {
	srcDir := filepath.Join(version, "configurations")
	dstDir := filepath.Join(version, launcher)

	// Check if the source directory exists
	info, err := os.Stat(srcDir)
	if os.IsNotExist(err) || !info.IsDir() {
		fmt.Printf("No configurations directory found at %s, skipping copy.\n", srcDir)
		return nil
	}

	fmt.Printf("Copying configurations from %s to %s...\n", srcDir, dstDir)

	// Recursive traversal of the source folder
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Calculate the relative path
		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(dstDir, relPath)

		// If it's a directory --> create it
		if info.IsDir() {
			if _, err := os.Stat(targetPath); os.IsNotExist(err) {
				if err := os.MkdirAll(targetPath, info.Mode()); err != nil {
					return err
				}
			}
			return nil
		}

		// If it's a file --> copy it
		return copyFile(path, targetPath)
	})
}

// copyFile copies the contents and permissions of a single file
func copyFile(srcFile, dstFile string) error {
	src, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer src.Close()

	// Create the target directory if needed
	if err := os.MkdirAll(filepath.Dir(dstFile), 0755); err != nil {
		return err
	}

	dst, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy the content
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	// Copy the permissions
	if info, err := os.Stat(srcFile); err == nil {
		os.Chmod(dstFile, info.Mode())
	}

	fmt.Printf("Copied %s --> %s\n", srcFile, dstFile)
	return nil
}