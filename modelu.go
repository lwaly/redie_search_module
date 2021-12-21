package main

import (
	"github.com/lwaly/redie_search_module/src"
	"github.com/lwaly/redie_search_module/src/module"
)

func main() {
}

func init() {
	module.Mod = src.CreateModule()
}
