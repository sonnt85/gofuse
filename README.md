github.com/sonnt85/gofuse -- Filesystems in Go
===================================

`github.com/sonnt85/gofuse` is a Go library for writing FUSE userspace
filesystems.

It is a from-scratch implementation of the kernel-userspace
communication protocol, and does not use the C library from the
project called FUSE. `github.com/sonnt85/gofuse` embraces Go fully for safety and
ease of programming.

Here’s how to get going:

    go get github.com/sonnt85/gofuse

Website: http://github.com/sonnt85/gofuse/

Github repository: https://github.com/bazil/fuse

API docs: http://godoc.org/github.com/sonnt85/gofuse

Our thanks to Russ Cox for his fuse library, which this project is
based on.
