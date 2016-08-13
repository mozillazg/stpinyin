package stpinyin_test

import (
	"fmt"

	"github.com/mozillazg/stpinyin"
)

func ExampleConvert() {
	s := "you1"
	fmt.Println(stpinyin.Convert(s))
	// Output: yōu
}
