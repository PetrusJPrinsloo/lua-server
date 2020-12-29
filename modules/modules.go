package modules

import (
	"fmt"
	"github.com/yuin/gopher-lua"
	"net/http"
)

func Loader(L *lua.LState) int {
	// register functions to the table
	mod := L.SetFuncs(L.NewTable(), exports)
	// register other stuff
	//L.SetField(mod, "name", lua.LString("value"))

	http.HandleFunc("/", index)

	// returns the module
	L.Push(mod)

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
	return 1
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Req: %s %s\n", r.Host, r.URL.Path)
}

var exports = map[string]lua.LGFunction{
	"myfunc": myfunc,
}

func myfunc(L *lua.LState) int {
	fmt.Println("Hello World!")
	return 0
}
