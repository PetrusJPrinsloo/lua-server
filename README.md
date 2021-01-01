# lua-server

![Build](https://github.com/PetrusJPrinsloo/lua-server/workflows/Build/badge.svg)

Small experiment of mine over the holiday to make a webserver with lua. This is by no means meant to be a production system so please don't use it as such.

`app/routes.lua` file starts a server like this:
```lua
local m = require("modules")

routes = {
    root = "/",
    user = "/user"
}

m.startServer('8080')
```

`root` always points to `app/index.lua`

Then for the `/user` route create a file `app/user.lua` and that file will be executed when that endpoint is hit:
```lua
print("Hello World")
-- do whatever...
```