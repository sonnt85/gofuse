package gofuse

import (
	"syscall"
)

const (
	ENODATA = Errno(syscall.ENODATA)
)

const (
	errNoXattr = ENODATA
)

func init() {
	errnoNames[errNoXattr] = "ENODATA"
}
