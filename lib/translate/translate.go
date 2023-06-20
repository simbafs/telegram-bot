package translate

import (
	"errors"
	"strings"
	"sync"

	"github.com/avast/retry-go"
	"github.com/bregydoc/gtranslate"
)

// TranslateChunk translates text without split it, which may cause reach the limit of 5000 characters
func TranslateChunk(text string) (string, error) {
	return gtranslate.TranslateWithParams(text, gtranslate.TranslationParams{
		To: "zh-TW",
	})
}

// splitToChunks splits a long text into chunks that each of them is under maxLength characters, and a complete sentance at the same time
func splitToChunks(text string, maxLength int) []string {
	var chunks []string

	delimiter := "\n"
	last := ""

	for _, chunk := range strings.Split(text, delimiter) {
		if len(last)+len(chunk) < maxLength {
			last += delimiter + chunk
		} else {
			chunks = append(chunks, strings.TrimSpace(last))
			last = chunk
		}
	}
	chunks = append(chunks, strings.TrimSpace(last))

	return chunks
}

// Translate translates text with unlimit length, it will cut text to fix the limit of google translate and compose them
func Translate(text string) (string, error) {
	chunks := splitToChunks(text, 5000)
	result := make([]string, len(chunks))

	var wg sync.WaitGroup
	var err error

	for i, chunk := range chunks {
		wg.Add(1)
		go func(i int, chunk string) {
			e := retry.Do(func() error {
				translated, e := TranslateChunk(chunk)
				if e != nil {
					result[i] = chunk
				} else {
					result[i] = translated
				}
				return e
			})
			if e != nil {
				err = errors.Join(err, e)
			}
			wg.Done()
		}(i, chunk)
	}

	wg.Wait()
	return strings.Join(result, "\n"), err
}
