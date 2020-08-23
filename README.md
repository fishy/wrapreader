[![PkgGoDev](https://pkg.go.dev/badge/github.com/fishy/wrapreader)](https://pkg.go.dev/github.com/fishy/wrapreader)
[![Go Report Card](https://goreportcard.com/badge/github.com/fishy/wrapreader)](https://goreportcard.com/report/github.com/fishy/wrapreader)

# WrapReader

WrapReader is a [Go](https://golang.org)
[library](https://pkg.go.dev/github.com/fishy/wrapreader).
It provides an [`io.ReadCloser`](https://pkg.go.dev/io?tab=doc#ReadCloser)
that wraps an [`io.Reader`](https://pkg.go.dev/io?tab=doc#Reader) and
an [`io.Closer`](https://pkg.go.dev/io?tab=doc#Closer) together.

It's useful when dealing with `io.Reader` implementations that wraps another
`io.ReadCloser`, but will not close the underlying reader, such as
[`gzip.Reader`](https://pkg.go.dev/compress/gzip?tab=doc#Reader).

([Example code](https://pkg.go.dev/github.com/fishy/wrapreader?tab=doc#example-Wrap))
