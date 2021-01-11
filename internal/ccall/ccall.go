package ccall

/*
#include "./include/jit/jit.h"

#cgo LDFLAGS: -L./libs/libjit.a -ljit
*/
import "C"

var (
	JIT_OPTLEVEL_NONE   = C.JIT_OPTLEVEL_NONE
	JIT_OPTLEVEL_NORMAL = C.JIT_OPTLEVEL_NORMAL
	JIT_NO_OFFSET       = C.JIT_NO_OFFSET
)
