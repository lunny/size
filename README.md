Size
=======================

[![CircleCI](https://circleci.com/gh/lunny/size.svg?style=shield)](https://circleci.com/gh/lunny/size)  [![codecov](https://codecov.io/gh/lunny/size/branch/master/graph/badge.svg)](https://codecov.io/gh/lunny/size)
[![](https://goreportcard.com/badge/github.com/lunny/size)](https://goreportcard.com/report/github.com/lunny/size)

Package size is a simple package to handle memory or disk size calculation or format

# Installation

```
go get github.com/lunny/size
```

# Usage

```Go
import (
    "fmt"

    . "github.com/lunny/size"
)

func main() {
    fmt.Println(10*M)

    size, _ := ParseSize("1.2K")
    fmt.Println(size)
}
```