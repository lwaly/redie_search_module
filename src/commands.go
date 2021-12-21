package src

import (
	"fmt"

	"github.com/lwaly/redie_search_module/src/module"
)

var commands map[string]*module.Command

func init() {
	command := CreateCommand_HGETSET()
	commands = make(map[string]*module.Command)
	commands[command.Name] = &command
}

func CreateModule() *module.Module {
	mod := module.NewMod()
	mod.Name = "rxhash"
	mod.Version = 1
	mod.Commands = make(map[string]*module.Command)
	for k, v := range commands {
		mod.Commands[k] = v
	}

	return mod
}

func CreateCommand_HGETSET() module.Command {
	return module.Command{
		Usage: "HGETSET key field value",
		Desc: `Sets the 'field' in Hash 'key' to 'value' and returns the previous value, if any.
Reply: String, the previous value or NULL if 'field' didn't exist. `,
		Name:     "hgetset",
		Flags:    "write fast deny-oom",
		FirstKey: 1, LastKey: 1, KeyStep: 1,
		Action: func(cmd module.CmdContext) int {
			fmt.Println("hgetset")
			cmd.Ctx.ReplyWithOK()
			return module.OK
		},
	}
}
