package main

import (
	"github.com/PetrusJPrinsloo/lua-server/modules"
	"github.com/yuin/gopher-lua"
)

func main() {
	L := lua.NewState()
	defer L.Close()
	L.PreloadModule("modules", modules.Loader)

	if err := L.DoFile("main.lua"); err != nil {
		panic(err)
	}
}
