package searcher

import (
	"sync"

	"async-search-engine/internal/model"
)

type reader interface {
	ReadFiles(root, fileType string, recursive bool) []string
	ContentContains(file string, keyword string) []string
}

type searcher struct {
	reader reader
}

func NewSearcher(
	reader reader,
) *searcher {
	return &searcher{
		reader,
	}
}

func (s *searcher) Search(directory, keyword, fileType string, poolSize int, recursive bool) <-chan model.Result {
	var wg sync.WaitGroup
	resultsChan := make(chan model.Result)
	files := s.reader.ReadFiles(directory, fileType, recursive)

	sem := make(chan struct{}, poolSize)
	for _, file := range files {
		wg.Add(1)
		sem <- struct{}{}
		go func(file string) {
			defer wg.Done()
			if matches := s.reader.ContentContains(file, keyword); len(matches) > 0 {
				for _, match := range matches {
					resultsChan <- model.Result{FileName: file, Line: match}
				}
			}
			<-sem
		}(file)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	return resultsChan
}
