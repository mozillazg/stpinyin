package stpinyin

import (
	"log"
	"testing"
)

func check(t *testing.T, ret, expect string) {
	if ret != expect {
		t.Errorf("Expected %v, got %v", expect, ret)
	}
}

func TestVersion(t *testing.T) {
	ret := Version()
	check(t, ret, "0.1.0")
}

func TestConvert(t *testing.T) {
	for input, expect := range map[string]string{
		"yi1":   "yī",
		"wu1":   "wū",
		"yu1":   "yū",
		"a1":    "ā",
		"ya1":   "yā",
		"wa1":   "wā",
		"o1":    "ō",
		"wo1":   "wō",
		"e2":    "é",
		"ye2":   "yé",
		"yue1":  "yuē",
		"ai1":   "āi",
		"wai1":  "wāi",
		"ei2":   "éi",
		"wei1":  "wēi",
		"ao2":   "áo",
		"yao1":  "yāo",
		"ou1":   "ōu",
		"you1":  "yōu",
		"an1":   "ān",
		"yan1":  "yān",
		"wan1":  "wān",
		"yuan1": "yuān",
		"en1":   "ēn",
		"yin1":  "yīn",
		"wen1":  "wēn",
		"yun4":  "yùn",
		"ang2":  "áng",
		"yang1": "yāng",
		"wang1": "wāng",
		"heng1": "hēng",
		"ying1": "yīng",
		"weng1": "wēng",
		"hong1": "hōng",
		"yong1": "yōng",
		"n4":    "ǹ",
		"n2g":   "ńg",
		"m2":    "ḿ",
		"mou2":  "móu",
		"móu":   "móu",
	} {
		log.Printf("input %v, expect %v", input, expect)
		ret := Convert(input)
		check(t, ret, expect)
	}
}

func Test_fixNumberPosition(t *testing.T) {
	for input, expect := range map[string]string{
		"yi1":   "yi1",
		"wu1":   "wu1",
		"yu1":   "yu1",
		"a1":    "a1",
		"ya1":   "ya1",
		"wa1":   "wa1",
		"o1":    "o1",
		"wo1":   "wo1",
		"e2":    "e2",
		"ye1":   "ye1",
		"yue1":  "yue1",
		"ai1":   "a1i",
		"wai1":  "wa1i",
		"ei2":   "e2i",
		"wei1":  "we1i",
		"ao2":   "a2o",
		"yao1":  "ya1o",
		"ou1":   "o1u",
		"you1":  "yo1u",
		"an1":   "a1n",
		"yan1":  "ya1n",
		"wan1":  "wa1n",
		"yuan1": "yua1n",
		"en1":   "e1n",
		"yin1":  "yi1n",
		"wen1":  "we1n",
		"yun4":  "yu4n",
		"ang2":  "a2ng",
		"yang1": "ya1ng",
		"wang1": "wa1ng",
		"heng1": "he1ng",
		"ying1": "yi1ng",
		"weng1": "we1ng",
		"hong1": "ho1ng",
		"yong1": "yo1ng",
		"n2":    "n2",
		"n2g":   "n2g",
		"m2":    "m2",
		"m":     "m",
		"g3":    "g3",
		"nv3":   "nv3",
		"er3":   "e3r",
	} {
		log.Printf("input %v, expect %v", input, expect)
		ret := fixNumberPosition(input)
		check(t, ret, expect)
	}
}
