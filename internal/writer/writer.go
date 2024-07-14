package writer

import (
	"fmt"

	"async-search-engine/internal/model"
)

type writer struct {
}

func NewWriter() *writer {
	return &writer{}
}

func (w *writer) WriteResult(results <-chan model.Result) {
	for result := range results {
		fmt.Println(result.FileName, ": ", w.highlightText(result.Line))
	}
}

func (w *writer) highlightText(word string) string {
	const (
		magentaColor = "\033[35m"
		resetColor   = "\033[0m"
	)

	return fmt.Sprintf(magentaColor + word + resetColor)
}
