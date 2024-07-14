package reader

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type reader struct {
}

func NewReader() *reader {
	return &reader{}
}

func (r *reader) ReadFiles(root, fileType string, recursive bool) []string {
	var fileList []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), fileType) {
			fileList = append(fileList, path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return fileList
}

func (r *reader) ContentContains(file, keyword string) []string {
	f, err := os.Open(file)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer f.Close()

	var matches []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			matches = append(matches, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return matches
}
