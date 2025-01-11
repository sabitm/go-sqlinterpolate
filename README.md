# SQL Query Interpolator for Go

This interpolator library is scraped from [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) and my primary usage of this is for debugging purpose.

## Install

```
go get github.com/sabitm/go-sqlinterpolate
```

## Usage

```
package main

import (
	"fmt"

	"github.com/sabitm/go-sqlinterpolate"
)

func main() {
	mysql := sqlinterpolate.MySQL

	res, err := mysql.Interpolate("SELECT * FROM a WHERE name = ? AND state IN (?, ?, ?, ?, ?)", []interface{}{"I'm fine", 42, int8(8), int16(-16), int32(32), int64(64)})
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
```

It should output:

```
SELECT * FROM a WHERE name = 'I\'m fine' AND state IN (42, 8, -16, 32, 64)
```
