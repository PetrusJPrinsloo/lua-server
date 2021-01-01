package server

import (
	"fmt"
	"github.com/PetrusJPrinsloo/lua-server/route"
	"github.com/PetrusJPrinsloo/lua-server/services"
	lua "github.com/yuin/gopher-lua"
	"net/http"
)

func Start(port string, routes []route.Route) {
	for _, route := range routes {
		route := route
		http.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
			L := lua.NewState()
			defer L.Close()
			L.PreloadModule("services", services.Loader)
			L.DoFile("app/" + route.File)
		})
	}
	fmt.Println("Listening on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
