package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	backupFolder   string = "backed"
	dirFlag        string = "dir"
	filesCountFlag string = "remain"
)

func main() {

	defaultDir := filepath.Join(os.Getenv("HOME"), "Downloads")
	targetDir := flag.String(dirFlag, defaultDir, "Directory to be cleared")
	filesToKeep := flag.Int(filesCountFlag, 5, "Number of file to keep in the directory")

	flag.Parse()

	dir := readDir(*targetDir)
	sortDir(dir)
	staleFiles := staleFiles(dir, *filesToKeep)
	files := filterStaleFiles(staleFiles)
	for _, f := range files {
		fmt.Println(f.Name())
	}
	backupOld(files, *targetDir)
}

func readDir(targetDir string) []os.FileInfo {
	files, err := ioutil.ReadDir(targetDir)

	if err != nil {
		panic(err)
	}

	return files
}

func sortDir(files []os.FileInfo) {
	sort.SliceStable(files, func(i, j int) bool {
		return files[i].ModTime().After(files[j].ModTime())
	})
}

func staleFiles(sortedFiles []os.FileInfo, filesToKeep int) []os.FileInfo {
	if len(sortedFiles) <= filesToKeep {
		return make([]os.FileInfo, 0)
	}

	return sortedFiles[filesToKeep:]
}

func backupOld(files []os.FileInfo, targetDir string) {
	if len(files) == 0 {
		return
	}
	backupPath := filepath.Join(targetDir, backupFolder)
	e := os.MkdirAll(backupPath, os.ModePerm)
	if e != nil {
		panic(e)
	}
	for _, file := range files {
		oldPath := filepath.Join(targetDir, file.Name())
		newPath := filepath.Join(backupPath, file.Name())
		e = os.Rename(oldPath, newPath)
		if e != nil {
			log.Fatal(e)
		}
	}
}

func filterStaleFiles(staleFiles []os.FileInfo) []os.FileInfo {
	result := make([]os.FileInfo, 0, len(staleFiles))
	for _, file := range staleFiles {
		if !ignoreFile(file) {
			result = append(result, file)
		}
	}
	return result
}

func ignoreFile(file os.FileInfo) bool {
	name := file.Name()
	return strings.HasSuffix(name, ".part") || strings.HasPrefix(name, ".") || name == backupFolder
}
