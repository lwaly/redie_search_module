package module

/*
#include <stdlib.h>
#include <stdint.h>
*/
import "C"
import (
	"bytes"
	"fmt"
	"unsafe"
)

//export RedisModule_OnLoad
func RedisModule_OnLoad(ctx uintptr, argv uintptr, argc int) C.int {
	return C.int(Ctx(ctx).Load(Mod, toStringSlice(argv, argc)))
}

func toStringSlice(argv uintptr, argc int) []String {
	args := make([]String, argc)
	size := int(unsafe.Sizeof(C.uintptr_t(0)))
	for i := 0; i < argc; i++ {
		ptr := unsafe.Pointer(argv + uintptr(size*i))
		args[i] = String(uintptr(*(*C.uintptr_t)(ptr)))
	}
	return args
}

//export cmd_func_call
func cmd_func_call(ctx uintptr, argv uintptr, argc int) C.int {
	args := toStringSlice(argv, argc)
	c := Ctx(ctx)
	// tempS := *(*string)(unsafe.Pointer(&args[0]))
	cmd, ok := Mod.Commands[args[0].String()]
	if ok {
		c.LogDebug(cmd.Name)
	}
	if IsDebugEnabled() {
		buf := bytes.NewBufferString(fmt.Sprintf("CmdFuncCall: %v", cmd.Name))
		for i := 0; i < argc; i++ {
			buf.WriteString(" ")
			buf.WriteString(args[i].String())
		}
		c.LogDebug(buf.String())
	}
	return C.int(cmd.Action(CmdContext{Ctx: c, Args: args}))
}

func init() {
	_ = C.int(0)
}
