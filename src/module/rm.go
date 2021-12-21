package module

//#include "./rm.h"
//#include "./wrapper.h"
import "C"
import (
	"fmt"
	"syscall"
	"unsafe"
)

/* ---------------- Defines common between core and modules --------------- */

/* Error status return values. */
const OK = C.REDISMODULE_OK
const ERR = C.REDISMODULE_ERR

/* API versions. */
const APIVER_1 = C.REDISMODULE_APIVER_1

/* API flags and constants */
const READ = C.REDISMODULE_READ
const WRITE = C.REDISMODULE_WRITE

const LIST_HEAD = C.REDISMODULE_LIST_HEAD
const LIST_TAIL = C.REDISMODULE_LIST_TAIL

/* Key types. */
const (
	// Return the type of the key. If the key pointer is NULL then `REDISMODULE_KEYTYPE_EMPTY` is returned.
	KEYTYPE_EMPTY = iota
	KEYTYPE_STRING
	KEYTYPE_LIST
	KEYTYPE_HASH
	KEYTYPE_SET
	KEYTYPE_ZSET
	KEYTYPE_MODULE
)

/* Reply types. */
const (
	REPLY_UNKNOWN = iota - 1
	REPLY_STRING
	REPLY_ERROR
	REPLY_INTEGER
	REPLY_ARRAY
	REPLY_NULL
)

/* Postponed array length. */
const POSTPONED_ARRAY_LEN = C.REDISMODULE_POSTPONED_ARRAY_LEN

/* Expire */
const NO_EXPIRE = C.REDISMODULE_NO_EXPIRE

/* Sorted set API flags. */
const (
	ZADD_XX = 1 << iota
	ZADD_NX
	ZADD_ADDED
	ZADD_UPDATED
	ZADD_NOP
)

/* Hash API flags. */
const (
	HASH_NONE = 0
	// Set if non-exists
	HASH_NX = 1 << iota
	// Set if exists
	HASH_XX
	// Use *char as args, ! do not use this flag
	HASH_CFIELDS
	// Check field exists
	HASH_EXISTS
)

/* Error messages. */
//const ERRORMSG_WRONGTYPE = C.REDISMODULE_ERRORMSG_WRONGTYPE
const ERRORMSG_WRONGTYPE = "WRONGTYPE Operation against a key holding the wrong kind of value"

const (
	LOG_DEBUG LogLevel = iota
	LOG_VERBOSE
	LOG_NOTICE
	LOG_WARNING
)

func getErrno() syscall.Errno {
	return syscall.Errno(C.get_errno())
}

// Produces a log message to the standard Redis log, the format accepts
// printf-alike specifiers, while level is a string describing the log
// level to use when emitting the log, and must be one of the following:
//
// * "debug"
// * "verbose"
// * "notice"
// * "warning"
//
// If the specified log level is invalid, verbose is used by default.
// There is a fixed limit to the length of the log line this function is able
// to emit, this limti is not specified but is guaranteed to be more than
// a few lines of text.
// void RM_Log(RedisModuleCtx *ctx, const char *levelstr, const char *fmt, ...);
func (ctx Ctx) Log(l LogLevel, format string, args ...interface{}) {
	c := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(c))
	C.CtxLog((*C.struct_RedisModuleCtx)(ctx.ptr()), C.int(l), c)
}
func (ctx Ctx) LogDebug(format string, args ...interface{}) {
	ctx.Log(LOG_DEBUG, format, args...)
}
func (ctx Ctx) LogVerbose(format string, args ...interface{}) {
	ctx.Log(LOG_VERBOSE, format, args...)
}
func (ctx Ctx) LogNotice(format string, args ...interface{}) {
	ctx.Log(LOG_NOTICE, format, args...)
}
func (ctx Ctx) LogWarn(format string, args ...interface{}) {
	ctx.Log(LOG_WARNING, format, args...)
}

func (ctx Ctx) Init(name string, version int, apiVersion int) int {
	c := C.CString(name)
	defer C.free(unsafe.Pointer(c))
	return (int)(C.RedisModule_Init((*C.struct_RedisModuleCtx)(ctx.ptr()), c, (C.int)(version), (C.int)(apiVersion)))
}
func (c Ctx) Load(mod *Module, args []String) int {
	if mod == nil {
		c.LogWarn("Load Mod must not nil")
		return ERR
	}
	if c.Init(mod.Name, mod.Version, APIVER_1) == ERR {
		c.LogWarn("Init mod %s failed", mod.Name)
		return ERR
	}
	c.LogDebug("Load module %s %v", mod.Name, args)

	for _, cmd := range mod.Commands {
		if c.CreateCommand(cmd) == ERR {
			return ERR
		}
	}

	return OK
}

func (c Ctx) CreateCommand(cmd *Command) int {
	name := C.CString(cmd.Name)
	defer C.free(unsafe.Pointer(name))
	flags := C.CString(cmd.Flags)
	defer C.free(unsafe.Pointer(flags))
	c.LogVerbose("CreateCommand#%v %s", cmd.Name, cmd.Usage)
	return (int)(C.CreateCommandCall((*C.struct_RedisModuleCtx)(c.ptr()), name, flags, C.int(cmd.FirstKey), C.int(cmd.LastKey), C.int(cmd.KeyStep)))
	// return (int)(C.CreateCommandCallID((*C.struct_RedisModuleCtx)(c.ptr()), C.int(id), name, flags, C.int(cmd.FirstKey), C.int(cmd.LastKey), C.int(cmd.KeyStep)))
}

func (ctx Ctx) ReplyWithOK() int {
	return int(C.ReplyWithSimpleString((*C.struct_RedisModuleCtx)(ctx.ptr()), C.CString("OK")))
}

func (v Ctx) ptr() unsafe.Pointer {
	return unsafe.Pointer(v)
}

func (v String) ptr() unsafe.Pointer {
	return unsafe.Pointer(v)
}

// =============================================================================
// ========================== String functions
// =============================================================================

// Given a string module object, this function returns the string pointer
// and length of the string. The returned pointer and length should only
// be used for read only accesses and never modified.
// const char *RM_StringPtrLen(RedisModuleString *str, size_t *len);
func (str String) String() string {
	l := uint64(0)
	ptr := C.StringPtrLen((*C.struct_RedisModuleString)(str.ptr()), (*C.size_t)(&l))
	return C.GoStringN(ptr, C.int(l))
}
