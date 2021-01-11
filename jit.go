package jit

import (
	"github.com/damilolarandolph/go-jit/internal/ccall"
)

var (
	TypeInt = &Type{ccall.TypeGoInt}
)
