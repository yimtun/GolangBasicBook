#### bytes.NewReader 可以将 []byte 转为 io.Reader 接口


```
go doc bytes.NewReader
```


```
package bytes // import "bytes"

func NewReader(b []byte) *Reader
    NewReader returns a new Reader reading from b.
```


```
go doc io.Reader
```


```
package io // import "io"

type Reader interface {
	Read(p []byte) (n int, err error)
}
```






