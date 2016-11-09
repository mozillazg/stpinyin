// Package stpinyin implements convert pinyin like this: you1 -> yōu
package stpinyin

import (
	"regexp"
	"strings"
)

// Version export version info
func Version() string {
	return "0.1.0"
}

var toneMap = map[string]string{
	"a1": "ā",
	"a2": "á",
	"a3": "ǎ",
	"a4": "à",
	"e1": "ē",
	"e2": "é",
	"e3": "ě",
	"e4": "è",
	"o1": "ō",
	"o2": "ó",
	"o3": "ǒ",
	"o4": "ò",
	"i1": "ī",
	"i2": "í",
	"i3": "ǐ",
	"i4": "ì",
	"u1": "ū",
	"u2": "ú",
	"u3": "ǔ",
	"u4": "ù",
	"v0": "ü",
	"v":  "ü",
	"v2": "ǘ",
	"v3": "ǚ",
	"v4": "ǜ",
	"n2": "ń",
	"n3": "ň",
	"n4": "ǹ",
	"m2": "ḿ",
}

var tonePositions = [][2]string{
	{"ang", "a"},
	{"eng", "e"},
	{"ong", "o"},
	{"ing", "i"},
	{"en", "e"},
	{"an", "a"},
	{"ou", "o"},
	{"ai", "a"},
	{"ei", "e"},
	{"ao", "a"},
	{"in", "i"},
	{"un", "u"},
	{"a", "a"},
	{"o", "o"},
	{"e", "e"},
	{"i", "i"},
	{"u", "u"},
	{"v", "v"},
	{"m", "m"},
	{"n", "n"},
}

var reToneNumber = regexp.MustCompile("^([^0-4]+)([0-4])$")

// Convert implements you1 -> yōu
func Convert(src string) (s string) {
	s = fixNumberPosition(src)
	for key, value := range toneMap {
		if strings.Contains(s, key) {
			s = strings.Replace(s, key, value, 1)
			break
		}
	}
	return s
}

const matchNumber = 3

// hang2 -> ha2ng
func fixNumberPosition(src string) (s string) {
	matchs := reToneNumber.FindStringSubmatch(src)
	if len(matchs) != matchNumber {
		return src
	}

	letter := matchs[1]
	number := matchs[2]
	// hang2 -> ha2ng
	for _, item := range tonePositions {
		suffix := item[0]
		position := item[1]
		if strings.HasSuffix(letter, suffix) {
			prefix := letter[:len(letter)-len(suffix)]
			// ang2 -> a2ng
			suffix = strings.Replace(suffix, position, position+number, 1)
			return prefix + suffix
		}
	}
	return src
}
