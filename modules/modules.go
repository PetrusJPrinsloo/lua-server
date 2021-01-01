package modules

import (
	"github.com/PetrusJPrinsloo/lua-server/route"
	"github.com/PetrusJPrinsloo/lua-server/server"
	"github.com/yuin/gopher-lua"
)

func Loader(L *lua.LState) int {
	// register functions to the table
	mod := L.SetFuncs(L.NewTable(), exports)

	// register other stuff
	L.SetGlobal("routes", L.NewTable())

	// returns the module
	L.Push(mod)

	return 1
}

//All the function to be exposed in Lua
var exports = map[string]lua.LGFunction{
	"startServer": startServer,
}

var Routes []route.Route

func startServer(L *lua.LState) int {

	//fmt.Printf("%#v\n", L.GetField("routes"))
	global := L.GetGlobal("routes").(*lua.LTable)
	global.ForEach(func(key lua.LValue, value lua.LValue) {
		Routes = append(Routes, route.Route{Path: value.String(), File: key.String() + ".lua"})
	})

	server.Start(L.ToString(1), Routes)
	return 0
}
