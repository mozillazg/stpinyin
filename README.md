stpinyin
==============

[![Build Status](https://travis-ci.org/mozillazg/stpinyin.svg?branch=master)](https://travis-ci.org/mozillazg/stpinyin)
[![Coverage Status](https://coveralls.io/repos/mozillazg/stpinyin/badge.svg?branch=master)](https://coveralls.io/r/mozillazg/stpinyin?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/mozillazg/stpinyin)](https://goreportcard.com/report/github.com/mozillazg/stpinyin)
[![GoDoc](https://godoc.org/github.com/mozillazg/stpinyin?status.svg)](https://godoc.org/github.com/mozillazg/stpinyin)

Convert pinyin like this: `you1 -> yōu`

Installation
------------

```
go get -u github.com/mozillazg/stpinyin
```

Install CLI tool:

```
go get -u github.com/mozillazg/stpinyin/stpinyin

$ stpinyin you1 hang2
yōu háng
```


Documentation
--------------

API documentation can be found here:
https://godoc.org/github.com/mozillazg/stpinyin


Usage
------

```go
package main

import (
	"fmt"

	"github.com/mozillazg/stpinyin"
)

func main() {
	s := "you1"
	fmt.Println(stpinyin.Convert(s))
	// Output: yōu
}
```

License
---------

Under the MIT License.
