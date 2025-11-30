package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CopyConfigurations copies everything from the root configurations directory
// and then from {version}/configurations/ into {version}/{launcher}/
func CopyConfigurations(version, launcher string, quiet bool) error {
	rootDir := "configurations"
	dstDir := filepath.Join(version, launcher)

	// Check if the root configurations directory exists
	info, err := os.Stat(rootDir)
	if err == nil && info.IsDir() {
		if !quiet {
			fmt.Printf("Copying base configurations from %s to %s...\n", rootDir, dstDir)
		}
		if err := copyAndOverride(rootDir, dstDir, quiet); err != nil {
			return err
		}
	} else {
		fmt.Printf("No base configurations directory found at %s, skipping base copy.\n", rootDir)
	}

	// Now copy the version-specific configurations, overriding base files if needed
	srcDir := filepath.Join(version, "configurations")
	info, err = os.Stat(srcDir)
	if os.IsNotExist(err) || !info.IsDir() {
		fmt.Printf("No version-specific configurations directory found at %s, skipping version-specific copy.\n", srcDir)
		return nil
	}

	if !quiet {
		fmt.Printf("Copying version-specific configurations from %s to %s...\n", srcDir, dstDir)
	}

	// Recursive traversal of the version-specific source folder
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

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

		// If it's a file --> copy it (overriding base)
		if err := copyFile(path, targetPath); err != nil {
			return err
		}
		if !quiet {
			fmt.Printf("Overridden by version-specific: %s --> %s\n", path, targetPath)
		}
		return nil
	})
}

// copyAndOverride copies files from srcDir to dstDir only if they don’t exist in dstDir or differ by content
func copyAndOverride(srcDir, dstDir string, quiet bool) error {
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(dstDir, relPath)

		if info.IsDir() {
			if _, err := os.Stat(targetPath); os.IsNotExist(err) {
				if err := os.MkdirAll(targetPath, info.Mode()); err != nil {
					return err
				}
			}
			return nil
		}

		// Check if file exists and is identical
		dstInfo, err := os.Stat(targetPath)
		if err == nil && !dstInfo.IsDir() {
			same, err := filesAreEqual(path, targetPath)
			if err != nil {
				return err
			}
			if same {
				// Skip copying identical file
				return nil
			}
		}

		if err := copyFile(path, targetPath); err != nil {
			return err
		}
		if !quiet {
			fmt.Printf("Copied base configuration: %s --> %s\n", path, targetPath)
		}
		return nil
	})
}

// filesAreEqual compares two files by size and MD5 checksum
func filesAreEqual(file1, file2 string) (bool, error) {
	info1, err := os.Stat(file1)
	if err != nil {
		return false, err
	}
	info2, err := os.Stat(file2)
	if err != nil {
		return false, err
	}

	if info1.Size() != info2.Size() {
		return false, nil
	}

	hash1, err := fileMD5(file1)
	if err != nil {
		return false, err
	}
	hash2, err := fileMD5(file2)
	if err != nil {
		return false, err
	}

	return hash1 == hash2, nil
}

// fileMD5 computes the MD5 checksum of a file
func fileMD5(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hasher := md5.New()
	if _, err := io.Copy(hasher, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
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

	return nil
}