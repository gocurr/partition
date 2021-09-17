# Partition - Go slice partition tool

To download, run:

```bash
go get -u github.com/gocurr/partition
```

Import it in your program as:

```go
import "github.com/gocurr/partition"
```

It requires Go 1.11 or later due to usage of Go Modules.

- Usage:

```go
s := []interface{}{1, 2, 3, 4, 5, 6, 7}
partSize := 2

ranges := partition.Ranges(len(s), partSize)
for _, r := range ranges {
    process(s[r.From:r.To])
}

func process(nums []interface{}) {
    // 
}
```
