package services

import (
	"fmt"
	"github.com/yuin/gopher-lua"
)

func Loader(L *lua.LState) int {
	// register functions to the table
	mod := L.SetFuncs(L.NewTable(), exports)

	// returns the module
	L.Push(mod)

	return 1
}

//All the function to be exposed in Lua
var exports = map[string]lua.LGFunction{
	"res": res,
}

func res(L *lua.LState) int {
	fmt.Print(L.ToString(1))
	return 0
}
