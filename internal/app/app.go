package app

import (
	"flag"
	"log"
	"unicode/utf8"

	"async-search-engine/internal/reader"
	"async-search-engine/internal/searcher"
	"async-search-engine/internal/writer"
	"async-search-engine/pkg/config"
)

var (
	directory, keyword, fileType string
	showHelp, recursive          bool
)

func init() {
	flag.StringVar(&directory, "directory", ".", "Directory to search")
	flag.StringVar(&keyword, "keyword", "", "Keyword to search for")
	flag.StringVar(&fileType, "file-type", "", "Type of files to search (e.g., .txt)")
	flag.BoolVar(&showHelp, "h", false, "Show help and exit")
	flag.Parse()
}

func Start() {
	if showHelp {
		flag.PrintDefaults()
		return
	}

	validateFlags()

	cfg := config.LoadConfig()
	reader := reader.NewReader()
	searcher := searcher.NewSearcher(
		reader,
	)
	writer := writer.NewWriter()

	resultsChan := searcher.Search(directory, keyword, fileType, cfg.MaxGoroutines, recursive)

	writer.WriteResult(resultsChan)
}

func validateFlags() {
	if utf8.RuneCountInString(keyword) < 3 {
		log.Fatal("required at least 3 symbol in a keyword")
	}
}
