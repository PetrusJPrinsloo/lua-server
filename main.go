package main

import (
	"github.com/PetrusJPrinsloo/lua-server/modules"
	"github.com/yuin/gopher-lua"
	"log"
	"os"
)

func main() {
	// Set up Logging to output to file instead of console
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	L := lua.NewState()
	defer L.Close()
	L.PreloadModule("modules", modules.Loader)

	if err := L.DoFile("app/routes.lua"); err != nil {
		log.Printf("Could not find a routes file, please create one at app/routes.lua")
		panic(err)
	}
}
