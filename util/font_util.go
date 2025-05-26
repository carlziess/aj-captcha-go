package util

import (
	"os"
	"sync"
	"unicode"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

// FontCache stores cached font objects.
type FontCache struct {
	fonts map[string]*truetype.Font
	mu    sync.RWMutex
}

// NewFontCache initializes an empty font cache.
func NewFontCache() *FontCache {
	return &FontCache{
		fonts: make(map[string]*truetype.Font),
	}
}

// GetFont retrieves a font from the cache or loads it from the given path.
func (fc *FontCache) GetFont(fontPath string) (*truetype.Font, error) {
	fc.mu.RLock()
	font, found := fc.fonts[fontPath]
	fc.mu.RUnlock()

	if found {
		return font, nil
	}

	fc.mu.Lock()
	defer fc.mu.Unlock()

	// Double check if another goroutine loaded the font while we were waiting for the lock
	font, found = fc.fonts[fontPath]
	if found {
		return font, nil
	}

	fontBytes, err := os.ReadFile(fontPath)
	if err != nil {
		return nil, err
	}

	parsedFont, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, err
	}

	fc.fonts[fontPath] = parsedFont
	return parsedFont, nil
}

// GlobalFontCache is a global instance of FontCache.
var GlobalFontCache = NewFontCache()

// GetEnOrChLength calculates the display length of a string containing English and/or Chinese characters.
// This function was originally in the old font_util.go and seems generally useful for text placement,
// so I'm keeping it here.
func GetEnOrChLength(text string) int {
	enCount, zhCount := 0, 0

	for _, t := range text {
		if unicode.Is(unicode.Han, t) {
			zhCount++
		} else {
			enCount++
		}
	}

	// These offsets might need adjustment depending on the specific font and rendering.
	// They were part of the original logic.
	chOffset := (25/2)*zhCount + 5 
	enOffset := enCount * 8

	return chOffset + enOffset
}
