package src

import (
	"fmt"

	"github.com/lwaly/redie_search_module/src/module"
)

var commands map[string]*module.Command

// var dataTypes []module.DataType

func init() {
	command := CreateCommand_HGETSET()
	commands = make(map[string]*module.Command)
	commands[command.Name] = &command
}

func CreateModule() *module.Module {
	mod := module.NewMod()
	mod.Name = "rxhash"
	mod.Version = 1
	mod.SemVer = "1.0.2-BETA"
	mod.Author = "wenerme"
	mod.Website = "http://github.com/wenerme/go-rm"
	mod.Desc = `This module will extends redis hash function`
	mod.Commands = make(map[string]*module.Command)
	for k, v := range commands {
		mod.Commands[k] = v
	}
	// mod.DataTypes = dataTypes
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
			// ctx, args := cmd.Ctx, cmd.Args
			// if len(cmd.Args) != 4 {
			// 	return ctx.WrongArity()
			// }
			// ctx.AutoMemory()
			// key, ok := openHashKey(ctx, args[1])
			// if !ok {
			// 	return module.ERR
			// }
			// // get the current value of the hash element
			// var val module.String
			// key.HashGet(module.HASH_NONE, cmd.Args[2], (*uintptr)(&val))
			// // set the element to the new value
			// key.HashSet(module.HASH_NONE, cmd.Args[2], cmd.Args[3])
			// if val.IsNull() {
			cmd.Ctx.ReplyWithOK()
			// } else {
			// 	ctx.ReplyWithString(val)
			// }
			return module.OK
		},
	}
}

// func openHashKey(ctx module.Ctx, k module.String) (module.Key, bool) {
// 	// key := ctx.OpenKey(k, module.READ|module.WRITE)
// 	// if key.KeyType() != module.KEYTYPE_EMPTY && key.KeyType() != module.KEYTYPE_HASH {
// 	// 	ctx.ReplyWithError(module.ERRORMSG_WRONGTYPE)
// 	// 	return module.Key(0), false
// 	// }
// 	// return key, true
// }
