package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mattn/go-isatty"
	"github.com/mozillazg/stpinyin"
)

func main() {
	version := flag.Bool("V", false, "Output version info")
	flag.Parse()
	if *version {
		v := stpinyin.Version()
		fmt.Printf("stpinyin %s\n", v)
		os.Exit(0)
	}

	textSlice := flag.Args()
	stdin := []byte{}
	if !isatty.IsTerminal(os.Stdin.Fd()) {
		stdin, _ = ioutil.ReadAll(os.Stdin)
	}
	if len(stdin) > 0 {
		textSlice = append(textSlice, string(stdin))
	}

	if len(textSlice) == 0 {
		fmt.Println("Usage: stpinyin STRING")
		os.Exit(1)
	}

	ret := []string{}
	for _, s := range textSlice {
		ret = append(ret, stpinyin.Convert(s))
	}
	fmt.Println(strings.Join(ret, " "))
}
