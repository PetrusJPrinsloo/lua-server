package server

import (
	"fmt"
	"github.com/PetrusJPrinsloo/lua-server/route"
	"github.com/PetrusJPrinsloo/lua-server/services"
	mysql "github.com/tengattack/gluasql/mysql"
	lua "github.com/yuin/gopher-lua"
	"log"
	"net/http"
)

func Start(port string, routes []route.Route) {

	for _, route := range routes {
		route := route

		http.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
			// admin
			w.Header().Set("Content-Type", "text/html")
			log.Printf("%s: %s", r.Method, r.URL)

			//set up lua state
			L := lua.NewState()
			defer L.Close()
			L.PreloadModule("services", services.Loader)
			L.PreloadModule("mysql", mysql.Loader)

			// Process endpoint
			switch r.Method {

			case http.MethodGet:
				doGet(w, r, L, route)
				break

			case http.MethodPost:
				doPost(w, r, L, route)
				break

			case http.MethodPut:
				doPut(w, r, L, route)
				break

			case http.MethodDelete:
				doDelete(w, r, L, route)
				break

			default:
				log.Printf("Unsupported method, currently only supported methods are GET, POST, PUT and DELETE")
			}
		})
	}
	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func doFile(w http.ResponseWriter, L *lua.LState, route route.Route) {
	err := L.DoFile("app/" + route.File)
	if err != nil {
		log.Printf(err.Error())
	}
	global := L.GetGlobal("response").(*lua.LTable)
	fmt.Fprintf(w, "%s", global.RawGetString("body").String())
}

func doPost(w http.ResponseWriter, r *http.Request, L *lua.LState, route route.Route) {

	table := L.NewTable()

	if err := r.ParseForm(); err != nil {
		log.Printf("ParseForm() err: %v", err)
		return
	}

	for key, value := range r.Form {
		value := value
		//fmt.Fprintf(w, "key = %s, value = %s\n", key, value)
		table.RawSetString(key, lua.LString(value[0]))
	}
	L.SetGlobal("POST_DATA", table)

	doFile(w, L, route)
}

func doPut(w http.ResponseWriter, r *http.Request, L *lua.LState, route route.Route) {
	doPost(w, r, L, route)
}

func doGet(w http.ResponseWriter, r *http.Request, L *lua.LState, route route.Route) {
	doFile(w, L, route)
}

func doDelete(w http.ResponseWriter, r *http.Request, L *lua.LState, route route.Route) {
	doGet(w, r, L, route)
}
