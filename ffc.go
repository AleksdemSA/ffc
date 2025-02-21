package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

type DirInfo struct {
	Path  string
	Count int
}

func countFiles(root string) map[string]int {
	fileCounts := make(map[string]int)
	dirs, err := os.ReadDir(root)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		os.Exit(1)
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			dirPath := filepath.Join(root, dir.Name())
			var count int
			filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() {
					count++
				}
				return nil
			})
			fileCounts[dirPath] = count
		}
	}

	return fileCounts
}

func main() {
	root := "." // Текущая директория по умолчанию
	if len(os.Args) > 1 {
		root = os.Args[1]
	}

	fileCounts := countFiles(root)
	dirs := make([]DirInfo, 0, len(fileCounts))
	for path, count := range fileCounts {
		dirs = append(dirs, DirInfo{Path: path, Count: count})
	}

	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].Count > dirs[j].Count
	})

	for _, dir := range dirs {
		fmt.Printf("%s: %d files\n", dir.Path, dir.Count)
	}
}
