package server

import (
	"bufio"
	"fmt"
	"github.com/PetrusJPrinsloo/lua-server/route"
	"github.com/PetrusJPrinsloo/lua-server/services"
	json "github.com/layeh/gopher-json"
	mysql "github.com/tengattack/gluasql/mysql"
	lua "github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/parse"
	"log"
	"net/http"
	"os"
)

func Start(port string, routes []route.Route) {

	for _, route := range routes {
		route := route
		codeToShare, err := CompileLua("app/" + route.File)

		if err != nil {
			log.Printf(err.Error())
		}

		http.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
			// admin
			w.Header().Set("Content-Type", "text/html")
			log.Printf("%s: %s", r.Method, r.URL)

			//set up lua state
			L := lua.NewState()
			defer L.Close()
			L.PreloadModule("services", services.Loader)
			L.PreloadModule("mysql", mysql.Loader)
			L.PreloadModule("json", json.Loader)

			// Process endpoint
			switch r.Method {

			case http.MethodGet:
				doGet(w, r, L, codeToShare)
				break

			case http.MethodPost:
				doPost(w, r, L, codeToShare)
				break

			case http.MethodPut:
				doPut(w, r, L, codeToShare)
				break

			case http.MethodDelete:
				doDelete(w, r, L, codeToShare)
				break

			default:
				log.Printf("Unsupported method, currently only supported methods are GET, POST, PUT and DELETE")
			}
		})
	}
	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// CompileLua reads the passed lua file from disk and compiles it.
func CompileLua(filePath string) (*lua.FunctionProto, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	chunk, err := parse.Parse(reader, filePath)
	if err != nil {
		return nil, err
	}
	proto, err := lua.Compile(chunk, filePath)
	if err != nil {
		return nil, err
	}
	return proto, nil
}

// DoCompiledFile takes a FunctionProto, as returned by CompileLua, and runs it in the LState. It is equivalent
// to calling DoFile on the LState with the original source file.
func DoCompiledFile(L *lua.LState, proto *lua.FunctionProto) error {
	lfunc := L.NewFunctionFromProto(proto)
	L.Push(lfunc)
	return L.PCall(0, lua.MultRet, nil)
}

func doFile(w http.ResponseWriter, L *lua.LState, proto *lua.FunctionProto) {
	DoCompiledFile(L, proto)
	global := L.GetGlobal("response").(*lua.LTable)
	fmt.Fprintf(w, "%s", global.RawGetString("body").String())
}

func doPost(w http.ResponseWriter, r *http.Request, L *lua.LState, proto *lua.FunctionProto) {

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

	doFile(w, L, proto)
}

func doPut(w http.ResponseWriter, r *http.Request, L *lua.LState, proto *lua.FunctionProto) {
	doPost(w, r, L, proto)
}

func doGet(w http.ResponseWriter, r *http.Request, L *lua.LState, proto *lua.FunctionProto) {
	doFile(w, L, proto)
}

func doDelete(w http.ResponseWriter, r *http.Request, L *lua.LState, proto *lua.FunctionProto) {
	doGet(w, r, L, proto)
}
