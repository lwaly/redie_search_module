package module

type Ctx uintptr
type LogLevel int
type String uintptr

type CmdFunc func(args CmdContext) int

type CmdContext struct {
	Ctx  Ctx
	Args []String
}
