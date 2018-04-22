# Size

The size is a simple package to handle memory or disk size calculation or format

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