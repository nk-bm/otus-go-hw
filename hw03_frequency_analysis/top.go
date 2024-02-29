package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

const taskWithAsterisk = true

type word struct {
	word  string
	count int
}

type wordsList struct {
	words  []*word
	markup map[string]*word
}

var re = regexp.MustCompile(`[[:punct:]]+$`)

func (wl *wordsList) addWord(str string) {
	if taskWithAsterisk {
		cleanStr := re.ReplaceAllString(str, "")
		if cleanStr == "" {
			if len(str) == 1 {
				return
			}
		} else {
			str = cleanStr
		}
		str = strings.ToLower(str)
	}

	if w, exists := wl.markup[str]; exists {
		w.count++
	} else {
		word := word{word: str, count: 1}
		wl.words = append(wl.words, &word)
		wl.markup[str] = &word
	}
}

func (wl *wordsList) getTop(n int) (res []string) {
	sort.Slice(wl.words, func(i, j int) bool {
		if wl.words[i].count == wl.words[j].count {
			return wl.words[i].word < wl.words[j].word
		}
		return wl.words[i].count > wl.words[j].count
	})

	for i := 0; i < n && i < len(wl.words); i++ {
		res = append(res, wl.words[i].word)
	}
	return
}

func newWordsList() *wordsList {
	return &wordsList{
		words:  []*word{},
		markup: make(map[string]*word),
	}
}

func Top10(str string) []string {
	wl := newWordsList()
	for _, word := range strings.Fields(str) {
		wl.addWord(word)
	}
	return wl.getTop(10)
}
