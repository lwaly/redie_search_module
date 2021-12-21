package module

// Redis module definition
type Module struct {
	Name     string
	Version  int
	Commands map[string]*Command
}

// Is debug enabled
func IsDebugEnabled() bool {
	// TODO Check redis log level
	return true
}

type Command struct {
	Usage  string
	Desc   string
	Name   string
	Action CmdFunc `json:"-"`
	// Use BuildCommandFLags to generate this flags
	Flags    string
	FirstKey int
	LastKey  int
	KeyStep  int
}

func NewMod() *Module {
	return &Module{}
}

// This module will be loaded
var Mod *Module

func init() {
}
